from rest_framework.status import *
from datetime import date
from dateutil.relativedelta import relativedelta
from django.db import transaction
import csv

from account.models import FoundationChannel
from warranty_manager.dal.foundation_business_dal import FoundationBusinessDAL
from warranty_manager.dal.foundation_category_dal import FoundationCategoryDAL
from warranty_manager.dal.foundation_channel_dal import FoundationChannelDAL
from warranty_manager.dal.newbasewarranty_dal import NewBaseWarrantyDAL
from warranty_manager.utils.exceptions import CustomError
from homeassure.models import NewBaseWarranty as BaseWarranty
from zopperassure.tasks import send_partner_warranty_verify_email_sms_to_customer

class RelianceResQPlanActivation:
    channel_name = "Reliance ResQ"

    def __init__(self):
        self.initiate_plan_activation()

    def activate_warranty(self, warranty_obj, duration):
        warranty_duration = int(duration)
        brand_warranty_duration = warranty_obj.brand_warranty_duration
        item_purchase_date = warranty_obj.financepartner_warranty.financepartner_item.purchasedate

        start_date = None
        end_date = None

        if warranty_obj.warranty_type == 13:
            start_date = item_purchase_date + relativedelta(months=brand_warranty_duration)
            end_date = start_date + relativedelta(months=warranty_duration) - relativedelta(days=1)

        elif warranty_obj.warranty_type == 17:
            start_date = item_purchase_date
            end_date = start_date + relativedelta(months=warranty_duration) - relativedelta(days=1)

        elif warranty_obj.warranty_type == 31:
            start_date = item_purchase_date
            end_date = start_date + relativedelta(months=warranty_duration) - relativedelta(days=1)

        with transaction.atomic():
            warranty_obj.status = BaseWarranty.ACTIVE
            warranty_obj.warranty_start_date = start_date
            warranty_obj.warranty_end_date = end_date
            warranty_obj.is_verified = BaseWarranty.VERIFIED
            warranty_obj.duration = warranty_duration
            warranty_obj.save()

            warranty_obj.financepartner_warranty.duration = warranty_duration
            warranty_obj.financepartner_warranty.save()
    
    def send_customer_communication(self, warranty_obj):
        try:
            new_base_warranty_obj = warranty_obj
            zopperassure_warranty_obj = warranty_obj.financepartner_warranty
            zopperassure_item_obj = warranty_obj.financepartner_warranty.financepartner_item
            warranty_detail_obj = warranty_obj.warrantydetail_set.all().first()

            category_dal = FoundationCategoryDAL()
            category_obj = category_dal.get_category_from_id(id=zopperassure_item_obj.category_id)

            channel_dal = FoundationChannelDAL()
            channel_obj = channel_dal.get_channel_from_name(name=self.channel_name)

            email_arg = {
                        "warranty_id": new_base_warranty_obj.warranty_id if new_base_warranty_obj.warranty_id else None,
                        "duration": new_base_warranty_obj.duration if new_base_warranty_obj.duration else None,
                        "brand_warranty_duration": new_base_warranty_obj.brand_warranty_duration if new_base_warranty_obj.brand_warranty_duration else None,
                        "warranty_start_date": new_base_warranty_obj.warranty_start_date.strftime('%d %b %Y') if new_base_warranty_obj.warranty_start_date else None,
                        "warranty_end_date": new_base_warranty_obj.warranty_end_date.strftime('%d %b %Y') if new_base_warranty_obj.warranty_end_date else None,
                        "warranty_type": new_base_warranty_obj.warranty_type if new_base_warranty_obj.warranty_type else None,
                        "purchased_on": new_base_warranty_obj.purchased_on.strftime('%d %b %Y') if new_base_warranty_obj.purchased_on else None,
                        "invoice_no": None,
                        "EW_start_date": new_base_warranty_obj.extrainfo.get('ew_start_date', None) if new_base_warranty_obj.extrainfo else None,
                        "EW_end_date": new_base_warranty_obj.extrainfo.get('ew_end_date', None) if new_base_warranty_obj.extrainfo else None,
                        "ADLD_start_date": new_base_warranty_obj.extrainfo.get('adld_start_date', None) if new_base_warranty_obj.extrainfo else None,
                        "ADLD_end_date": new_base_warranty_obj.extrainfo.get('adld_end_date', None) if new_base_warranty_obj.extrainfo else None,
                        "user_name": zopperassure_warranty_obj.customer.name.replace("\u00a0","").strip() if zopperassure_warranty_obj.customer.name else None,
                        "user_email": zopperassure_warranty_obj.customer.email if zopperassure_warranty_obj.customer.email else None,
                        "user_contact_no": zopperassure_warranty_obj.customer.mobile if zopperassure_warranty_obj.customer.mobile else None,
                        "user_address": zopperassure_warranty_obj.customer.address.replace("\u00a0","").strip() if zopperassure_warranty_obj.customer.address else None,
                        "amount": zopperassure_warranty_obj.kitprice.price if zopperassure_warranty_obj.kitprice.price else None,
                        "customer_amount": zopperassure_warranty_obj.kitprice.soldprice if zopperassure_warranty_obj.kitprice.soldprice else None,
                        "category_name": category_obj.name if category_obj else None,
                        "product_name": zopperassure_item_obj.model if zopperassure_item_obj.model else None,
                        "serialnumber": zopperassure_item_obj.serialnumber if zopperassure_item_obj.serialnumber else None,
                        "item_serialnumber": zopperassure_item_obj.imeinumber,
                        "purchase_date": zopperassure_item_obj.purchasedate.strftime('%d %b %Y') if zopperassure_item_obj.purchasedate else None,
                        "brand_name": zopperassure_item_obj.brand if zopperassure_item_obj.brand else None,
                        "product_purchased_on": zopperassure_item_obj.purchasedate.strftime('%d %b %Y') if zopperassure_item_obj.purchasedate else None,
                        "product_invoiceno": zopperassure_item_obj.invoicenumber,
                        "category_id": category_obj.id if category_obj else None,
                        "token": "Token 4ffa27086eb8aa70d651f09d8e3db63d2d983f16",
                        "sms_text": None,
                        "channel": channel_obj.name if channel_obj else None,
                        "display_name": new_base_warranty_obj.display_plan_name,
                        "ew_start_date": new_base_warranty_obj.warranty_start_date.strftime('%d %b %Y') if new_base_warranty_obj.warranty_start_date else None,
                        "ew_end_date": new_base_warranty_obj.warranty_end_date.strftime('%d %b %Y') if new_base_warranty_obj.warranty_end_date else None,
                        "ew_duration": new_base_warranty_obj.duration - new_base_warranty_obj.brand_warranty_duration if new_base_warranty_obj.duration and new_base_warranty_obj.brand_warranty_duration else None,
                        "adld_start_date": warranty_detail_obj.warranty_start_date.strftime('%d %b %Y') if warranty_detail_obj and warranty_detail_obj.warranty_start_date else None,
                        "adld_end_date": warranty_detail_obj.warranty_end_date.strftime('%d %b %Y') if warranty_detail_obj and warranty_detail_obj.warranty_end_date else None,
                        "adld_duration": warranty_detail_obj.duration if warranty_detail_obj and warranty_detail_obj.duration else None,
                        "product_price": zopperassure_item_obj.price if zopperassure_item_obj.price else None,
                    }
            
            send_partner_warranty_verify_email_sms_to_customer.apply_async(args=(email_arg,))
        except CustomError as error:
            print('Custom error in outer try-except: {}'.format(error))
        except Exception as error:
            print('Exception in outer try-except: {}'.format(error))

    
    def initiate_plan_activation(self):
        try:
            with open("assurance_3568_data.csv") as file_obj:
                csv_reader = csv.DictReader(file_obj)

                for row in csv_reader:
                    filters=dict()
                    filters['id'] = row['id']

                    nbw_warranty_dal = NewBaseWarrantyDAL()
                    select_related_list = ['financepartner_warranty', 'financepartner_warranty__customer', 'financepartner_warranty__kitprice', 'financepartner_warranty__financepartner_item', 'financepartner_warranty__kitprice__payment']

                    api_warranties = nbw_warranty_dal.filter_warranties(filters, select_related_list=select_related_list)
                    print("api_warranties : ", api_warranties)

                    all_warranties = list(api_warranties)

                    self.activate_warranty(all_warranties[0], row['duration'])
                    self.send_customer_communication(all_warranties[0])
        except CustomError as error:
            print('Custom error in outer try-except: {}'.format(error))
        except Exception as error:
            print('Exception in outer try-except: {}'.format(error))

RelianceResQPlanActivation()