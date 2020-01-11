package client

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type UserGroup struct {
	Id int32 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ResultInfo struct {
	Curpage int32 `xml:"curpage,omitempty"`

	Perpage int32 `xml:"perpage,omitempty"`

	Numpages int32 `xml:"numpages,omitempty"`

	Numresults int32 `xml:"numresults,omitempty"`

	Maxresults int32 `xml:"maxresults,omitempty"`
}

type UserGroupArray struct {
	Item []*UserGroup `xml:"item,omitempty"`
}

type UserGroupPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *UserGroupArray `xml:"results,omitempty"`
}

type User struct {
	Id int32 `xml:"id,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	Email string `xml:"email,omitempty"`

	Companyname string `xml:"companyname,omitempty"`

	Address string `xml:"address,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Credits int32 `xml:"credits,omitempty"`

	Clientcode string `xml:"clientcode,omitempty"`

	Comments string `xml:"comments,omitempty"`

	Creationtime string `xml:"creationtime,omitempty"`

	Lastmodified string `xml:"lastmodified,omitempty"`

	Active bool `xml:"active,omitempty"`

	Account_id int32 `xml:"account_id,omitempty"`

	Account_name string `xml:"account_name,omitempty"`

	Usergroups *UserGroupArray `xml:"usergroups,omitempty"`

	Creditthreshold int32 `xml:"creditthreshold,omitempty"`

	Notificationrecipients string `xml:"notificationrecipients,omitempty"`
}

type UserCredentials struct {
	Id int32 `xml:"id,omitempty"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type UserV2 struct {
	Id int32 `xml:"id,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	Email string `xml:"email,omitempty"`

	Balance float32 `xml:"balance,omitempty"`

	Balancethreshold float32 `xml:"balancethreshold,omitempty"`

	Notificationrecipients string `xml:"notificationrecipients,omitempty"`

	Companyname string `xml:"companyname,omitempty"`

	Address string `xml:"address,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Clientcode string `xml:"clientcode,omitempty"`

	Comments string `xml:"comments,omitempty"`

	Creationtime string `xml:"creationtime,omitempty"`

	Lastmodified string `xml:"lastmodified,omitempty"`

	Active bool `xml:"active,omitempty"`

	Account_id int32 `xml:"account_id,omitempty"`

	Account_name string `xml:"account_name,omitempty"`

	Usergroups *UserGroupArray `xml:"usergroups,omitempty"`
}

type UserV2Array struct {
	Item []*UserV2 `xml:"item,omitempty"`
}

type UserV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *UserV2Array `xml:"results,omitempty"`
}

type Account struct {
	Id int32 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`

	Environment string `xml:"environment,omitempty"`

	Reseller bool `xml:"reseller,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Description string `xml:"description,omitempty"`

	Creationtime string `xml:"creationtime,omitempty"`

	Lastmodified string `xml:"lastmodified,omitempty"`

	Credits int32 `xml:"credits,omitempty"`

	Clientcode string `xml:"clientcode,omitempty"`

	Comments string `xml:"comments,omitempty"`

	Usernameprefix string `xml:"usernameprefix,omitempty"`

	Creditthreshold int32 `xml:"creditthreshold,omitempty"`
}

type AccountArray struct {
	Item []*Account `xml:"item,omitempty"`
}

type AccountPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *AccountArray `xml:"results,omitempty"`
}

type AccountV2 struct {
	Id int32 `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`

	Environment string `xml:"environment,omitempty"`

	Reseller bool `xml:"reseller,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Description string `xml:"description,omitempty"`

	Creationtime string `xml:"creationtime,omitempty"`

	Lastmodified string `xml:"lastmodified,omitempty"`

	Clientcode string `xml:"clientcode,omitempty"`

	Comments string `xml:"comments,omitempty"`

	Usernameprefix string `xml:"usernameprefix,omitempty"`

	Balance float32 `xml:"balance,omitempty"`

	Balancethreshold float32 `xml:"balancethreshold,omitempty"`
}

type AccountV2Array struct {
	Item []*AccountV2 `xml:"item,omitempty"`
}

type AccountV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *AccountV2Array `xml:"results,omitempty"`
}

type Session struct {
	Reactid string `xml:"reactid,omitempty"`

	User_id int32 `xml:"user_id,omitempty"`

	User_nickname string `xml:"user_nickname,omitempty"`

	Creationtime string `xml:"creationtime,omitempty"`

	Lastvisit string `xml:"lastvisit,omitempty"`

	Hits int32 `xml:"hits,omitempty"`

	Ip string `xml:"ip,omitempty"`
}

type SessionArray struct {
	Item []*Session `xml:"item,omitempty"`
}

type SessionPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *SessionArray `xml:"results,omitempty"`
}

type IntArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ intArray"`

	Item []int32 `xml:"item,omitempty"`
}

type AccountCreationToken struct {
	Token string `xml:"token,omitempty"`

	Url string `xml:"url,omitempty"`
}

type AccountOrderToken struct {
	Token string `xml:"token,omitempty"`

	Url string `xml:"url,omitempty"`
}

type PCReeks struct {
	Reeksid int32 `xml:"reeksid,omitempty"`

	Huisnr_van int32 `xml:"huisnr_van,omitempty"`

	Huisnr_tm int32 `xml:"huisnr_tm,omitempty"`

	Wijkcode string `xml:"wijkcode,omitempty"`

	Lettercombinatie string `xml:"lettercombinatie,omitempty"`

	Reeksindicatie string `xml:"reeksindicatie,omitempty"`

	Straatid int32 `xml:"straatid,omitempty"`

	Straatnaam string `xml:"straatnaam,omitempty"`

	Straatnaam_nen string `xml:"straatnaam_nen,omitempty"`

	Straatnaam_ptt string `xml:"straatnaam_ptt,omitempty"`

	Straatnaam_extract string `xml:"straatnaam_extract,omitempty"`

	Plaatsid int32 `xml:"plaatsid,omitempty"`

	Plaatsnaam string `xml:"plaatsnaam,omitempty"`

	Plaatsnaam_ptt string `xml:"plaatsnaam_ptt,omitempty"`

	Plaatsnaam_extract string `xml:"plaatsnaam_extract,omitempty"`

	Gemeenteid int32 `xml:"gemeenteid,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Gemeentecode int32 `xml:"gemeentecode,omitempty"`

	Cebucocode int32 `xml:"cebucocode,omitempty"`

	Provinciecode string `xml:"provinciecode,omitempty"`

	Provincienaam string `xml:"provincienaam,omitempty"`
}

type SearchParts struct {
	Huisnr int32 `xml:"huisnr,omitempty"`

	Toevoeging string `xml:"toevoeging,omitempty"`

	Wijkcode string `xml:"wijkcode,omitempty"`

	Lettercombinatie string `xml:"lettercombinatie,omitempty"`

	Straat string `xml:"straat,omitempty"`

	Plaats string `xml:"plaats,omitempty"`

	Gemeente string `xml:"gemeente,omitempty"`

	Provincie string `xml:"provincie,omitempty"`

	Reeksindicatie string `xml:"reeksindicatie,omitempty"`
}

type PCReeksArray struct {
	Item []*PCReeks `xml:"item,omitempty"`
}

type PCReeksSearchPartsPagedResult struct {
	Extra *SearchParts `xml:"extra,omitempty"`

	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *PCReeksArray `xml:"results,omitempty"`
}

type Perceel struct {
	Perceelid int32 `xml:"perceelid,omitempty"`

	Huisnr int32 `xml:"huisnr,omitempty"`

	Huisnr_toevoeging string `xml:"huisnr_toevoeging,omitempty"`

	Perceelnummer int32 `xml:"perceelnummer,omitempty"`

	Reeksid int32 `xml:"reeksid,omitempty"`

	Huisnr_van int32 `xml:"huisnr_van,omitempty"`

	Huisnr_tm int32 `xml:"huisnr_tm,omitempty"`

	Wijkcode string `xml:"wijkcode,omitempty"`

	Lettercombinatie string `xml:"lettercombinatie,omitempty"`

	Reeksindicatie string `xml:"reeksindicatie,omitempty"`

	Straatid int32 `xml:"straatid,omitempty"`

	Straatnaam string `xml:"straatnaam,omitempty"`

	Straatnaam_nen string `xml:"straatnaam_nen,omitempty"`

	Straatnaam_ptt string `xml:"straatnaam_ptt,omitempty"`

	Straatnaam_extract string `xml:"straatnaam_extract,omitempty"`

	Plaatsid int32 `xml:"plaatsid,omitempty"`

	Plaatsnaam string `xml:"plaatsnaam,omitempty"`

	Plaatsnaam_ptt string `xml:"plaatsnaam_ptt,omitempty"`

	Plaatsnaam_extract string `xml:"plaatsnaam_extract,omitempty"`

	Gemeenteid int32 `xml:"gemeenteid,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Gemeentecode int32 `xml:"gemeentecode,omitempty"`

	Cebucocode int32 `xml:"cebucocode,omitempty"`

	Provinciecode string `xml:"provinciecode,omitempty"`

	Provincienaam string `xml:"provincienaam,omitempty"`
}

type PerceelArray struct {
	Item []*Perceel `xml:"item,omitempty"`
}

type PerceelSearchPartsPagedResult struct {
	Extra *SearchParts `xml:"extra,omitempty"`

	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *PerceelArray `xml:"results,omitempty"`
}

type RangeAddress struct {
	Reeks *PCReeks `xml:"reeks,omitempty"`

	Huisnr int32 `xml:"huisnr,omitempty"`

	Huisnr_toevoeging string `xml:"huisnr_toevoeging,omitempty"`
}

type RangeAddressArray struct {
	Item []*RangeAddress `xml:"item,omitempty"`
}

type RangeAddressPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *RangeAddressArray `xml:"results,omitempty"`
}

type District struct {
	Gemeenteid int32 `xml:"gemeenteid,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Gemeentecode int32 `xml:"gemeentecode,omitempty"`

	Provinciecode string `xml:"provinciecode,omitempty"`
}

type DistrictArray struct {
	Item []*District `xml:"item,omitempty"`
}

type DistrictPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DistrictArray `xml:"results,omitempty"`
}

type City struct {
	Plaatsid int32 `xml:"plaatsid,omitempty"`

	Plaatsnaam string `xml:"plaatsnaam,omitempty"`

	Plaatsnaam_ptt string `xml:"plaatsnaam_ptt,omitempty"`

	Plaatsnaam_extract string `xml:"plaatsnaam_extract,omitempty"`

	Gemeenteid int32 `xml:"gemeenteid,omitempty"`
}

type CityArray struct {
	Item []*City `xml:"item,omitempty"`
}

type CityPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CityArray `xml:"results,omitempty"`
}

type CityV2 struct {
	Plaatsid int32 `xml:"plaatsid,omitempty"`

	Plaatsnaam string `xml:"plaatsnaam,omitempty"`

	Plaatsnaam_ptt string `xml:"plaatsnaam_ptt,omitempty"`

	Plaatsnaam_extract string `xml:"plaatsnaam_extract,omitempty"`

	Gemeenteid int32 `xml:"gemeenteid,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Gemeentecode int32 `xml:"gemeentecode,omitempty"`

	Cebucocode int32 `xml:"cebucocode,omitempty"`

	Provinciecode string `xml:"provinciecode,omitempty"`

	Provincienaam string `xml:"provincienaam,omitempty"`
}

type CityV2Array struct {
	Item []*CityV2 `xml:"item,omitempty"`
}

type CityV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CityV2Array `xml:"results,omitempty"`
}

type Province struct {
	Provinciecode string `xml:"provinciecode,omitempty"`

	Provincienaam string `xml:"provincienaam,omitempty"`
}

type ProvinceArray struct {
	Item []*Province `xml:"item,omitempty"`
}

type ProvincePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *ProvinceArray `xml:"results,omitempty"`
}

type Neighborhood struct {
	Nbcode string `xml:"nbcode,omitempty"`
}

type NeighborhoodArray struct {
	Item []*Neighborhood `xml:"item,omitempty"`
}

type NeighborhoodPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *NeighborhoodArray `xml:"results,omitempty"`
}

type NeighborhoodName struct {
	Nbcode string `xml:"nbcode,omitempty"`

	Nbname string `xml:"nbname,omitempty"`
}

type NeighborhoodNameArray struct {
	Item []*NeighborhoodName `xml:"item,omitempty"`
}

type NeighborhoodNamePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *NeighborhoodNameArray `xml:"results,omitempty"`
}

type AreaCode struct {
	Areacode string `xml:"areacode,omitempty"`
}

type AreaCodeArray struct {
	Item []*AreaCode `xml:"item,omitempty"`
}

type AreaCodePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *AreaCodeArray `xml:"results,omitempty"`
}

type HeaderAuthenticateType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ HeaderAuthenticate"`

	Reactid string `xml:"reactid,omitempty"`
}

type HeaderLoginType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ HeaderLogin"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type BelgianBusinessDossier struct {
	Errors *BelgianBusinessDossierErrors `xml:"errors,omitempty"`

	Company *BelgianBusinessCompany `xml:"company,omitempty"`

	Bank_accounts *BelgianBusinessBankAccountArray `xml:"bank_accounts,omitempty"`

	Annual_financial_statements *BelgianBusinessAnnualFinancialStatementArray `xml:"annual_financial_statements,omitempty"`

	Legal_representative_companies *BelgianBusinessLegalRepresentativeCompanyArray `xml:"legal_representative_companies,omitempty"`

	Legal_representatives *BelgianBusinessLegalRepresentativeArray `xml:"legal_representatives,omitempty"`

	Stakeholders *BelgianBusinessStakeHolderArray `xml:"stakeholders,omitempty"`
}

type BelgianBusinessDossierErrors struct {
	Bank_accounts_error string `xml:"bank_accounts_error,omitempty"`

	Annual_financial_statements_error string `xml:"annual_financial_statements_error,omitempty"`

	Legal_representative_companies_error string `xml:"legal_representative_companies_error,omitempty"`

	Legal_representatives_error string `xml:"legal_representatives_error,omitempty"`

	Stakeholders_error string `xml:"stakeholders_error,omitempty"`
}

type BelgianBusinessCompany struct {
	Vat_number string `xml:"vat_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Website string `xml:"website,omitempty"`

	Start_date *BelgianBusinessDate `xml:"start_date,omitempty"`

	Establishment_units int32 `xml:"establishment_units,omitempty"`

	Type_ *BelgianBusinessDatedItem `xml:"type,omitempty"`

	Activities *BelgianBusinessDatedItemArray `xml:"activities,omitempty"`

	Addresses *BelgianBusinessAddressArray `xml:"addresses,omitempty"`

	Juridical_form *BelgianBusinessDatedItem `xml:"juridical_form,omitempty"`

	Juridical_situation *BelgianBusinessDatedItem `xml:"juridical_situation,omitempty"`

	Legal_representatives_count string `xml:"legal_representatives_count,omitempty"`

	Legal_representative_companies_count string `xml:"legal_representative_companies_count,omitempty"`
}

type BelgianBusinessAnnualFinancialStatementArray struct {
	Item []*BelgianBusinessAnnualFinancialStatement `xml:"item,omitempty"`
}

type BelgianBusinessAnnualFinancialStatement struct {
	Id string `xml:"id,omitempty"`

	Year int32 `xml:"year,omitempty"`

	Number_of_employees int32 `xml:"number_of_employees,omitempty"`

	Turnover float32 `xml:"turnover,omitempty"`

	Equity float32 `xml:"equity,omitempty"`

	Revisors *BelgianBusinessRevisorArray `xml:"revisors,omitempty"`

	Account_date *BelgianBusinessDate `xml:"account_date,omitempty"`

	Current_assets float32 `xml:"current_assets,omitempty"`

	Gross_operating_margin float32 `xml:"gross_operating_margin,omitempty"`

	Tangible_fixed_assets float32 `xml:"tangible_fixed_assets,omitempty"`

	Gain_loss_period float32 `xml:"gain_loss_period,omitempty"`

	Current_ratio float32 `xml:"current_ratio,omitempty"`

	Net_cash float32 `xml:"net_cash,omitempty"`

	Self_financing_degree float32 `xml:"self_financing_degree,omitempty"`

	Return_on_equity float32 `xml:"return_on_equity,omitempty"`

	Added_value float32 `xml:"added_value,omitempty"`

	Balance_type float32 `xml:"balance_type,omitempty"`

	Ratios *BelgianBusinessRatios `xml:"ratios,omitempty"`

	Social_balance *BelgianBusinessSocialBalance `xml:"social_balance,omitempty"`

	Financial_values *BelgianBusinessFinancialValues `xml:"financial_values,omitempty"`
}

type BelgianBusinessStakeHolderArray struct {
	Item []*BelgianBusinessStakeHolder `xml:"item,omitempty"`
}

type BelgianBusinessStakeHolder struct {
	Type_ string `xml:"type,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Id string `xml:"id,omitempty"`

	Profession string `xml:"profession,omitempty"`

	Address *BelgianBusinessInternationalAddress `xml:"address,omitempty"`

	Mandate *BelgianBusinessDatedItem `xml:"mandate,omitempty"`

	Source_date *BelgianBusinessDate `xml:"source_date,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`
}

type BelgianBusinessLegalRepresentativeArray struct {
	Item []*BelgianBusinessLegalRepresentative `xml:"item,omitempty"`
}

type BelgianBusinessLegalRepresentative struct {
	Company_vat_number string `xml:"company_vat_number,omitempty"`

	Role *BelgianBusinessDatedItem `xml:"role,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`
}

type BelgianBusinessLegalRepresentativeCompanyArray struct {
	Item []*BelgianBusinessLegalRepresentativeCompany `xml:"item,omitempty"`
}

type BelgianBusinessLegalRepresentativeCompany struct {
	Vat_number string `xml:"vat_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Role *BelgianBusinessDatedItem `xml:"role,omitempty"`
}

type BelgianBusinessFinancialValues struct {
	Fixed_assets float32 `xml:"fixed_assets,omitempty"`

	Number_of_employees int32 `xml:"number_of_employees,omitempty"`

	Number_of_employees_men int32 `xml:"number_of_employees_men,omitempty"`

	Number_of_employees_women int32 `xml:"number_of_employees_women,omitempty"`

	Capital float32 `xml:"capital,omitempty"`

	Financial_fixed_assets float32 `xml:"financial_fixed_assets,omitempty"`

	Currents_assets float32 `xml:"currents_assets,omitempty"`

	Amounts_receivable_more_one_year float32 `xml:"amounts_receivable_more_one_year,omitempty"`

	Trade_debtors_within_one_year float32 `xml:"trade_debtors_within_one_year,omitempty"`

	Equity float32 `xml:"equity,omitempty"`

	Amounts_payable_within_one_year float32 `xml:"amounts_payable_within_one_year,omitempty"`

	Trade_debts_payable_within_one_year float32 `xml:"trade_debts_payable_within_one_year,omitempty"`

	Turnover float32 `xml:"turnover,omitempty"`

	Gross_operating_margin float32 `xml:"gross_operating_margin,omitempty"`

	Increase_decrease_stocks_work_contracts_progress float32 `xml:"increase_decrease_stocks_work_contracts_progress,omitempty"`

	Contributions_gifts_legacies_grants float32 `xml:"contributions_gifts_legacies_grants,omitempty"`

	Remuneration_social_security_pensions float32 `xml:"remuneration_social_security_pensions,omitempty"`

	Depreciation_other_amounts_written_down_formation_expenses_intangible_tangible_fixed_assets float32 `xml:"depreciation_other_amounts_written_down_formation_expenses_intangible_tangible_fixed_assets,omitempty"`

	Amounts_written_down_stocks_contracts_progress_trade_debtors_appropriations_write_backs float32 `xml:"amounts_written_down_stocks_contracts_progress_trade_debtors_appropriations_write_backs,omitempty"`

	Provisions_risks_charges_appropriations_write_backs float32 `xml:"provisions_risks_charges_appropriations_write_backs,omitempty"`

	Operating_charges float32 `xml:"operating_charges,omitempty"`

	Operating_profit_loss float32 `xml:"operating_profit_loss,omitempty"`

	Financial_income float32 `xml:"financial_income,omitempty"`

	Financial_charges float32 `xml:"financial_charges,omitempty"`

	Extraordinary_income float32 `xml:"extraordinary_income,omitempty"`

	Extraordinary_charges float32 `xml:"extraordinary_charges,omitempty"`

	Income_taxes float32 `xml:"income_taxes,omitempty"`

	Gain_loss_period float32 `xml:"gain_loss_period,omitempty"`

	Tangible_fixed_assets_acquisition_including_produced_fixed_assets float32 `xml:"tangible_fixed_assets_acquisition_including_produced_fixed_assets,omitempty"`

	Tangible_fixed_assets_revaluation_gains_third_parties float32 `xml:"tangible_fixed_assets_revaluation_gains_third_parties,omitempty"`

	Tangible_fixed_assets_depreciations_amount_written_down_acquired_third_party float32 `xml:"tangible_fixed_assets_depreciations_amount_written_down_acquired_third_party,omitempty"`

	Outstanding_taxes_payable_due_tax_authorities float32 `xml:"outstanding_taxes_payable_due_tax_authorities,omitempty"`

	Remuneration_social_security_recipient_national_social_security_office float32 `xml:"remuneration_social_security_recipient_national_social_security_office,omitempty"`

	Amounts_receivable_within_one_year float32 `xml:"amounts_receivable_within_one_year,omitempty"`

	Accrued_charges_deferred_income float32 `xml:"accrued_charges_deferred_income,omitempty"`

	Current_investments float32 `xml:"current_investments,omitempty"`

	Cash_bank_hand float32 `xml:"cash_bank_hand,omitempty"`

	Equity_liabilities float32 `xml:"equity_liabilities,omitempty"`

	Operating_income float32 `xml:"operating_income,omitempty"`

	Other_operating_income float32 `xml:"other_operating_income,omitempty"`

	Operating_subsidies_compensatory_amounts float32 `xml:"operating_subsidies_compensatory_amounts,omitempty"`

	Capital_subsidies_granted_public_authorities_recorded_income_period float32 `xml:"capital_subsidies_granted_public_authorities_recorded_income_period,omitempty"`

	Raw_materials_consumables float32 `xml:"raw_materials_consumables,omitempty"`

	Services_other_goods float32 `xml:"services_other_goods,omitempty"`

	Hired_temporary_staff_costs_enterprise float32 `xml:"hired_temporary_staff_costs_enterprise,omitempty"`

	Persons_placed_enterprises_disposal_costs_enterprise float32 `xml:"persons_placed_enterprises_disposal_costs_enterprise,omitempty"`

	Employees_recorded_personnel_register_average_number_employees_calculated_full_time_equivalents float32 `xml:"employees_recorded_personnel_register_average_number_employees_calculated_full_time_equivalents,omitempty"`

	Stock_raw_materials_consumables float32 `xml:"stock_raw_materials_consumables,omitempty"`

	Stock_goods_purchased_resale float32 `xml:"stock_goods_purchased_resale,omitempty"`

	Stock_immovable_property_intended_sale float32 `xml:"stock_immovable_property_intended_sale,omitempty"`

	Advance_payments_purchases_stocks float32 `xml:"advance_payments_purchases_stocks,omitempty"`

	Stocks_contracts_progress float32 `xml:"stocks_contracts_progress,omitempty"`

	Deferred_charges_accrued_income float32 `xml:"deferred_charges_accrued_income,omitempty"`

	Own_construction_capitalised float32 `xml:"own_construction_capitalised,omitempty"`

	Stock_work_progress float32 `xml:"stock_work_progress,omitempty"`

	Stock_finished_goods float32 `xml:"stock_finished_goods,omitempty"`

	Contracts_progress float32 `xml:"contracts_progress,omitempty"`

	Personal_guarantees_provided_or_irrevocably_promised_by_enterprise float32 `xml:"personal_guarantees_provided_or_irrevocably_promised_by_enterprise,omitempty"`

	Value_added_taxes_charged_by_enterprise float32 `xml:"value_added_taxes_charged_by_enterprise,omitempty"`

	Purchases_raw_materials_consumables float32 `xml:"purchases_raw_materials_consumables,omitempty"`

	Value_added_taxes_charged_to_enterprise float32 `xml:"value_added_taxes_charged_to_enterprise,omitempty"`

	Amounts_written_down_current_assets float32 `xml:"amounts_written_down_current_assets,omitempty"`

	Provisions_financial_nature_appropriations float32 `xml:"provisions_financial_nature_appropriations,omitempty"`

	Provisions_financial_nature_uses_write_backs float32 `xml:"provisions_financial_nature_uses_write_backs,omitempty"`

	Extraordinary_depreciation_extraordinary_amounts_written_down float32 `xml:"extraordinary_depreciation_extraordinary_amounts_written_down,omitempty"`

	Amounts_written_down_financial_fixed_assets float32 `xml:"amounts_written_down_financial_fixed_assets,omitempty"`

	Provisions_extraordinary_liabilities_charges_appropriations_uses float32 `xml:"provisions_extraordinary_liabilities_charges_appropriations_uses,omitempty"`

	Write_back_depreciation_amounts_written_down_intangible_tangible_fixed_assets float32 `xml:"write_back_depreciation_amounts_written_down_intangible_tangible_fixed_assets,omitempty"`

	Write_back_amounts_written_down_financial_fixed_assets float32 `xml:"write_back_amounts_written_down_financial_fixed_assets,omitempty"`

	Write_back_provisions_extraordinary_liabilities_charges float32 `xml:"write_back_provisions_extraordinary_liabilities_charges,omitempty"`

	Losses_disposal_fixed_assets float32 `xml:"losses_disposal_fixed_assets,omitempty"`

	Income_taxes_result_period float32 `xml:"income_taxes_result_period,omitempty"`

	Transfer_from_deferred_taxes float32 `xml:"transfer_from_deferred_taxes,omitempty"`

	Transfer_to_deferred_taxes float32 `xml:"transfer_to_deferred_taxes,omitempty"`

	Assets float32 `xml:"assets,omitempty"`

	Interest_subsidies_granted_public_authorities_recorded_income_period float32 `xml:"interest_subsidies_granted_public_authorities_recorded_income_period,omitempty"`

	Charges_discounting_amounts_receivable float32 `xml:"charges_discounting_amounts_receivable,omitempty"`

	Debt_charges float32 `xml:"debt_charges,omitempty"`

	Depreciation_loan_issue_expenses_reimbursement_premiums float32 `xml:"depreciation_loan_issue_expenses_reimbursement_premiums,omitempty"`

	Provisions_pensions_similar_obligations_appropriations_write_backs float32 `xml:"provisions_pensions_similar_obligations_appropriations_write_backs,omitempty"`

	Current_portion_amounts_payable_more_one_year_falling_due_within_one_year float32 `xml:"current_portion_amounts_payable_more_one_year_falling_due_within_one_year,omitempty"`

	Amounts_payable_more_one_year float32 `xml:"amounts_payable_more_one_year,omitempty"`

	Provisions_deferred_taxes float32 `xml:"provisions_deferred_taxes,omitempty"`

	Amounts_payable float32 `xml:"amounts_payable,omitempty"`

	Reserves float32 `xml:"reserves,omitempty"`

	Accumulated_profits_losses float32 `xml:"accumulated_profits_losses,omitempty"`

	Financial_debts_payable_within_one_year float32 `xml:"financial_debts_payable_within_one_year,omitempty"`

	Formation_expenses float32 `xml:"formation_expenses,omitempty"`

	Intangible_fixed_assets float32 `xml:"intangible_fixed_assets,omitempty"`

	Tangible_fixed_assets float32 `xml:"tangible_fixed_assets,omitempty"`

	Land_buildings_acquisition_including_produced_fixed_assets float32 `xml:"land_buildings_acquisition_including_produced_fixed_assets,omitempty"`

	Plant_machinery_equipment_acquisition_including_produced_fixed_assets float32 `xml:"plant_machinery_equipment_acquisition_including_produced_fixed_assets,omitempty"`

	Furniture_vehicles_acquisition_including_produced_fixed_assets float32 `xml:"furniture_vehicles_acquisition_including_produced_fixed_assets,omitempty"`

	Leasing_other_similar_rights_acquisition_including_produced_fixed_assets float32 `xml:"leasing_other_similar_rights_acquisition_including_produced_fixed_assets,omitempty"`

	Other_tangible_fixed_assets_acquisition_including_produced_fixed_assets float32 `xml:"other_tangible_fixed_assets_acquisition_including_produced_fixed_assets,omitempty"`

	Tangible_fixed_assets_under_construction float32 `xml:"tangible_fixed_assets_under_construction,omitempty"`

	Land_buildings_revaluation_gains_third_parties float32 `xml:"land_buildings_revaluation_gains_third_parties,omitempty"`

	Plant_machinery_equipment_revaluation_gains_third_parties float32 `xml:"plant_machinery_equipment_revaluation_gains_third_parties,omitempty"`

	Furniture_vehicles_revaluation_gains_third_parties float32 `xml:"furniture_vehicles_revaluation_gains_third_parties,omitempty"`

	Leasing_other_similar_rights_revaluation_gains_third_parties float32 `xml:"leasing_other_similar_rights_revaluation_gains_third_parties,omitempty"`

	Other_tangible_fixed_assets_revaluation_gains_third_parties float32 `xml:"other_tangible_fixed_assets_revaluation_gains_third_parties,omitempty"`

	Tangible_fixed_assets_under_construction_advance_payments_revaluation_gains_third_parties float32 `xml:"tangible_fixed_assets_under_construction_advance_payments_revaluation_gains_third_parties,omitempty"`

	Land_buildings_depreciations_amount_written_down_acquired_third_party float32 `xml:"land_buildings_depreciations_amount_written_down_acquired_third_party,omitempty"`

	Plant_machinery_equipment_depreciations_amount_written_down_acquired_third_party float32 `xml:"plant_machinery_equipment_depreciations_amount_written_down_acquired_third_party,omitempty"`

	Furniture_vehicles_depreciations_amount_written_down_acquired_third_party float32 `xml:"furniture_vehicles_depreciations_amount_written_down_acquired_third_party,omitempty"`

	Leasing_other_similar_rights_depreciations_amount_written_down_acquired_third_party float32 `xml:"leasing_other_similar_rights_depreciations_amount_written_down_acquired_third_party,omitempty"`

	Other_tangible_fixed_assets_depreciations_amount_written_down_acquired_third_party float32 `xml:"other_tangible_fixed_assets_depreciations_amount_written_down_acquired_third_party,omitempty"`

	Tangible_fixed_assets_under_construction_advance_payments_depreciations float32 `xml:"tangible_fixed_assets_under_construction_advance_payments_depreciations,omitempty"`

	Employees_recorded_personnel_register_total_number_closing_date float32 `xml:"employees_recorded_personnel_register_total_number_closing_date,omitempty"`

	Number_employees_personnel_register_closing_date_financial_year_men_full_time float32 `xml:"number_employees_personnel_register_closing_date_financial_year_men_full_time,omitempty"`

	Number_employees_personnel_register_closing_date_financial_year_women_full_time float32 `xml:"number_employees_personnel_register_closing_date_financial_year_women_full_time,omitempty"`

	Number_employees_personnel_register_closing_date_financial_year_men_part_time float32 `xml:"number_employees_personnel_register_closing_date_financial_year_men_part_time,omitempty"`

	Number_employees_personnel_register_closing_date_financial_year_women_part_time float32 `xml:"number_employees_personnel_register_closing_date_financial_year_women_part_time,omitempty"`

	Investments float32 `xml:"investments,omitempty"`

	Added_value float32 `xml:"added_value,omitempty"`
}

type BelgianBusinessSocialBalance struct {
	Employees_worked_hours int32 `xml:"employees_worked_hours,omitempty"`

	Advantages_addition_wages float32 `xml:"advantages_addition_wages,omitempty"`

	Professional_training_hours float32 `xml:"professional_training_hours,omitempty"`

	Professional_training_net_costs float32 `xml:"professional_training_net_costs,omitempty"`

	Average_number_of_employees_fte float32 `xml:"average_number_of_employees_fte,omitempty"`

	Employees_end_of_year_management_fte float32 `xml:"employees_end_of_year_management_fte,omitempty"`

	Employees_end_of_year float32 `xml:"employees_end_of_year,omitempty"`

	Employees_end_of_year_workers_fte float32 `xml:"employees_end_of_year_workers_fte,omitempty"`

	Employees_end_of_year_others_fte float32 `xml:"employees_end_of_year_others_fte,omitempty"`

	Employees_new_fte float32 `xml:"employees_new_fte,omitempty"`

	Employees_end_of_year_men_fte float32 `xml:"employees_end_of_year_men_fte,omitempty"`

	Employees_end_of_year_women_fte float32 `xml:"employees_end_of_year_women_fte,omitempty"`

	Employees_end_of_year_university_fte float32 `xml:"employees_end_of_year_university_fte,omitempty"`

	Employees_end_of_year_non_university_fte float32 `xml:"employees_end_of_year_non_university_fte,omitempty"`

	Employees_end_of_year_secondary_education_fte float32 `xml:"employees_end_of_year_secondary_education_fte,omitempty"`

	Employees_end_of_year_primary_education_fte float32 `xml:"employees_end_of_year_primary_education_fte,omitempty"`

	Employees_contract_indefinite_period_fte float32 `xml:"employees_contract_indefinite_period_fte,omitempty"`

	Employees_contract_definite_period_fte float32 `xml:"employees_contract_definite_period_fte,omitempty"`

	Employees_contract_termination_fte float32 `xml:"employees_contract_termination_fte,omitempty"`

	Employees_contract_termination_retirement_fte float32 `xml:"employees_contract_termination_retirement_fte,omitempty"`

	Employees_contract_termination_early_retirement_fte float32 `xml:"employees_contract_termination_early_retirement_fte,omitempty"`

	Employees_contract_termination_dismissal float32 `xml:"employees_contract_termination_dismissal,omitempty"`

	Employees_contract_termination_other_reason float32 `xml:"employees_contract_termination_other_reason,omitempty"`

	Hired_temporary_staff_fte float32 `xml:"hired_temporary_staff_fte,omitempty"`

	Hired_temporary_staff_worked_hours float32 `xml:"hired_temporary_staff_worked_hours,omitempty"`

	Hired_temporary_staff_costs float32 `xml:"hired_temporary_staff_costs,omitempty"`
}

type BelgianBusinessRatios struct {
	Net_assets float32 `xml:"net_assets,omitempty"`

	Return_on_equity float32 `xml:"return_on_equity,omitempty"`

	Gross_operating_sales_margin float32 `xml:"gross_operating_sales_margin,omitempty"`

	Net_operating_sales_margin float32 `xml:"net_operating_sales_margin,omitempty"`

	Rotation_fixed_assets float32 `xml:"rotation_fixed_assets,omitempty"`

	Total_assets_to_turnover float32 `xml:"total_assets_to_turnover,omitempty"`

	Current_ratio float32 `xml:"current_ratio,omitempty"`

	Quick_ratio float32 `xml:"quick_ratio,omitempty"`

	Immediate_liquidity float32 `xml:"immediate_liquidity,omitempty"`

	Write_downs_to_added_value float32 `xml:"write_downs_to_added_value,omitempty"`

	Net_cash float32 `xml:"net_cash,omitempty"`

	Net_cash_ratio float32 `xml:"net_cash_ratio,omitempty"`

	Net_current_assets float32 `xml:"net_current_assets,omitempty"`

	Cash_flow float32 `xml:"cash_flow,omitempty"`

	Number_days_customer_credit float32 `xml:"number_days_customer_credit,omitempty"`

	Number_days_supplier_credit float32 `xml:"number_days_supplier_credit,omitempty"`

	Investments float32 `xml:"investments,omitempty"`

	Dfl float32 `xml:"dfl,omitempty"`

	Own_assets_to_capital float32 `xml:"own_assets_to_capital,omitempty"`

	Cash_flow_to_debt float32 `xml:"cash_flow_to_debt,omitempty"`

	Cash_flow_to_long_term_debt float32 `xml:"cash_flow_to_long_term_debt,omitempty"`

	Long_term_debt_to_short_term_debt float32 `xml:"long_term_debt_to_short_term_debt,omitempty"`

	Debt_repayment float32 `xml:"debt_repayment,omitempty"`

	Self_financing_degree float32 `xml:"self_financing_degree,omitempty"`

	Outstanding_tax_to_liabilities float32 `xml:"outstanding_tax_to_liabilities,omitempty"`

	Added_value float32 `xml:"added_value,omitempty"`

	Payroll_costs_to_added_value float32 `xml:"payroll_costs_to_added_value,omitempty"`

	Financial_charges_to_added_value float32 `xml:"financial_charges_to_added_value,omitempty"`

	Net_profit_to_added_value float32 `xml:"net_profit_to_added_value,omitempty"`

	Added_value_to_operating_income float32 `xml:"added_value_to_operating_income,omitempty"`

	Investment_ratio float32 `xml:"investment_ratio,omitempty"`

	Added_value_per_employee float32 `xml:"added_value_per_employee,omitempty"`
}

type BelgianBusinessRevisorArray struct {
	Item []*BelgianBusinessRevisor `xml:"item,omitempty"`
}

type BelgianBusinessRevisor struct {
	Type_ string `xml:"type,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Id string `xml:"id,omitempty"`

	Profession string `xml:"profession,omitempty"`

	Address *BelgianBusinessInternationalAddress `xml:"address,omitempty"`

	Mandate *BelgianBusinessDatedItem `xml:"mandate,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`
}

type BelgianBusinessBankAccountArray struct {
	Item []*BelgianBusinessBankAccount `xml:"item,omitempty"`
}

type BelgianBusinessBankAccount struct {
	Iban string `xml:"iban,omitempty"`

	Bic string `xml:"bic,omitempty"`

	Iban_valid bool `xml:"iban_valid,omitempty"`

	Start_date *BelgianBusinessDate `xml:"start_date,omitempty"`
}

type BelgianBusinessAddressArray struct {
	Item []*BelgianBusinessAddress `xml:"item,omitempty"`
}

type BelgianBusinessAddress struct {
	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	Box string `xml:"box,omitempty"`

	Country_code string `xml:"country_code,omitempty"`

	Municipality string `xml:"municipality,omitempty"`

	Post_code string `xml:"post_code,omitempty"`
}

type BelgianBusinessInternationalAddressArray struct {
	Item []*BelgianBusinessInternationalAddress `xml:"item,omitempty"`
}

type BelgianBusinessInternationalAddress struct {
	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	Box string `xml:"box,omitempty"`

	Country_code string `xml:"country_code,omitempty"`

	Municipality string `xml:"municipality,omitempty"`

	Post_code string `xml:"post_code,omitempty"`
}

type BelgianBusinessDatedItemArray struct {
	Item []*BelgianBusinessDatedItem `xml:"item,omitempty"`
}

type BelgianBusinessDatedItem struct {
	Code string `xml:"code,omitempty"`

	Label string `xml:"label,omitempty"`

	Start_date *BelgianBusinessDate `xml:"start_date,omitempty"`

	End_date *BelgianBusinessDate `xml:"end_date,omitempty"`
}

type BelgianBusinessDate struct {
	Day int32 `xml:"day,omitempty"`

	Month int32 `xml:"month,omitempty"`

	Year int32 `xml:"year,omitempty"`
}

type BovagMember struct {
	Bovag_id string `xml:"bovag_id,omitempty"`

	Name string `xml:"name,omitempty"`

	Certificate_url string `xml:"certificate_url,omitempty"`

	Departments *BovagDepartmentArray `xml:"departments,omitempty"`

	Dutch_business_identifiers *DutchBusinessIdentifiers `xml:"dutch_business_identifiers,omitempty"`
}

type BovagDepartment struct {
	Code string `xml:"code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type BovagDepartmentArray struct {
	Item []*BovagDepartment `xml:"item,omitempty"`
}

type DutchBusinessIdentifiers struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Rsi_number string `xml:"rsi_number,omitempty"`
}

type BusinessReference struct {
	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	Tradename string `xml:"Tradename,omitempty"`

	Streetname string `xml:"Streetname,omitempty"`

	City string `xml:"City,omitempty"`
}

type BusinessReferenceArray struct {
	Item []*BusinessReference `xml:"item,omitempty"`
}

type BusinessReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessReferenceArray `xml:"results,omitempty"`
}

type BusinessReferenceV2 struct {
	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	Tradename string `xml:"Tradename,omitempty"`

	EstablishmentStreetname string `xml:"EstablishmentStreetname,omitempty"`

	EstablishmentCity string `xml:"EstablishmentCity,omitempty"`

	CorrespondenceStreetname string `xml:"CorrespondenceStreetname,omitempty"`

	CorrespondenceCity string `xml:"CorrespondenceCity,omitempty"`
}

type BusinessReferenceV2Array struct {
	Item []*BusinessReferenceV2 `xml:"item,omitempty"`
}

type BusinessReferenceV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessReferenceV2Array `xml:"results,omitempty"`
}

type BusinessReferenceV3 struct {
	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	Tradename string `xml:"Tradename,omitempty"`

	EstablishmentStreetname string `xml:"EstablishmentStreetname,omitempty"`

	EstablishmentCity string `xml:"EstablishmentCity,omitempty"`

	CorrespondenceStreetname string `xml:"CorrespondenceStreetname,omitempty"`

	CorrespondenceCity string `xml:"CorrespondenceCity,omitempty"`

	IndicationEconomicallyActive bool `xml:"IndicationEconomicallyActive,omitempty"`
}

type BusinessReferenceV3Array struct {
	Item []*BusinessReferenceV3 `xml:"item,omitempty"`
}

type BusinessReferenceV3PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessReferenceV3Array `xml:"results,omitempty"`
}

type BusinessDossier struct {
	RegisterLetter string `xml:"RegisterLetter,omitempty"`

	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	ChamberNo string `xml:"ChamberNo,omitempty"`

	Legalformcode string `xml:"Legalformcode,omitempty"`

	LegalformcodeText string `xml:"LegalformcodeText,omitempty"`

	PreviousDossierNo string `xml:"PreviousDossierNo,omitempty"`

	PreviousSubDossierNo string `xml:"PreviousSubDossierNo,omitempty"`

	Tradename30 string `xml:"Tradename30,omitempty"`

	Tradename1x30 string `xml:"Tradename1x30,omitempty"`

	Tradename2x30 string `xml:"Tradename2x30,omitempty"`

	Tradename45 string `xml:"Tradename45,omitempty"`

	TradenameFull string `xml:"TradenameFull,omitempty"`

	EstablishmentAddress string `xml:"EstablishmentAddress,omitempty"`

	EstablishmentPostcodeCity string `xml:"EstablishmentPostcodeCity,omitempty"`

	EstablishmentStreetname string `xml:"EstablishmentStreetname,omitempty"`

	EstablishmentHouseNo int32 `xml:"EstablishmentHouseNo,omitempty"`

	EstablishmentHouseNoAddition string `xml:"EstablishmentHouseNoAddition,omitempty"`

	EstablishmentPostcode string `xml:"EstablishmentPostcode,omitempty"`

	CorrespondenceAddress string `xml:"CorrespondenceAddress,omitempty"`

	CorrespondencePostcodeCity string `xml:"CorrespondencePostcodeCity,omitempty"`

	CorrespondenceStreetname string `xml:"CorrespondenceStreetname,omitempty"`

	CorrespondenceHouseNo int32 `xml:"CorrespondenceHouseNo,omitempty"`

	CorrespondenceHouseNoAddition string `xml:"CorrespondenceHouseNoAddition,omitempty"`

	CorrespondencePostcode string `xml:"CorrespondencePostcode,omitempty"`

	TelephoneNo string `xml:"TelephoneNo,omitempty"`

	Domainname string `xml:"Domainname,omitempty"`

	PrimaryActivitycode string `xml:"PrimaryActivitycode,omitempty"`

	SecondaryActivitycode1 string `xml:"SecondaryActivitycode1,omitempty"`

	SecondaryActivitycode2 string `xml:"SecondaryActivitycode2,omitempty"`

	PrimaryActivitycodeText string `xml:"PrimaryActivitycodeText,omitempty"`

	SecondaryActivitycode1Text string `xml:"SecondaryActivitycode1Text,omitempty"`

	SecondaryActivitycode2Text string `xml:"SecondaryActivitycode2Text,omitempty"`

	PersonnelFulltime string `xml:"PersonnelFulltime,omitempty"`

	ClassPersonnelFulltime string `xml:"ClassPersonnelFulltime,omitempty"`

	IndicationOrganisationcode string `xml:"IndicationOrganisationcode,omitempty"`

	IndicationEconomicallyActive string `xml:"IndicationEconomicallyActive,omitempty"`

	IndicationNonMailing string `xml:"IndicationNonMailing,omitempty"`

	IndicationBankruptcy string `xml:"IndicationBankruptcy,omitempty"`

	IndicationDIP string `xml:"IndicationDIP,omitempty"`
}

type BusinessDossierArray struct {
	Item []*BusinessDossier `xml:"item,omitempty"`
}

type BusinessDossierPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessDossierArray `xml:"results,omitempty"`
}

type BusinessDossierV2 struct {
	RegisterLetter string `xml:"RegisterLetter,omitempty"`

	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	ChamberNo string `xml:"ChamberNo,omitempty"`

	Legalformcode string `xml:"Legalformcode,omitempty"`

	LegalformcodeText string `xml:"LegalformcodeText,omitempty"`

	PreviousDossierNo string `xml:"PreviousDossierNo,omitempty"`

	PreviousSubDossierNo string `xml:"PreviousSubDossierNo,omitempty"`

	Tradename30 string `xml:"Tradename30,omitempty"`

	Tradename1x30 string `xml:"Tradename1x30,omitempty"`

	Tradename2x30 string `xml:"Tradename2x30,omitempty"`

	Tradename45 string `xml:"Tradename45,omitempty"`

	TradenameFull string `xml:"TradenameFull,omitempty"`

	EstablishmentPostcode string `xml:"EstablishmentPostcode,omitempty"`

	EstablishmentCity string `xml:"EstablishmentCity,omitempty"`

	EstablishmentStreetname string `xml:"EstablishmentStreetname,omitempty"`

	EstablishmentHouseNo int32 `xml:"EstablishmentHouseNo,omitempty"`

	EstablishmentHouseNoAddition string `xml:"EstablishmentHouseNoAddition,omitempty"`

	CorrespondencePostcode string `xml:"CorrespondencePostcode,omitempty"`

	CorrespondenceCity string `xml:"CorrespondenceCity,omitempty"`

	CorrespondenceStreetname string `xml:"CorrespondenceStreetname,omitempty"`

	CorrespondenceHouseNo int32 `xml:"CorrespondenceHouseNo,omitempty"`

	CorrespondenceHouseNoAddition string `xml:"CorrespondenceHouseNoAddition,omitempty"`

	TelephoneNo string `xml:"TelephoneNo,omitempty"`

	Domainname string `xml:"Domainname,omitempty"`

	PrimaryActivitycode string `xml:"PrimaryActivitycode,omitempty"`

	SecondaryActivitycode1 string `xml:"SecondaryActivitycode1,omitempty"`

	SecondaryActivitycode2 string `xml:"SecondaryActivitycode2,omitempty"`

	PrimaryActivitycodeText string `xml:"PrimaryActivitycodeText,omitempty"`

	SecondaryActivitycode1Text string `xml:"SecondaryActivitycode1Text,omitempty"`

	SecondaryActivitycode2Text string `xml:"SecondaryActivitycode2Text,omitempty"`

	PersonnelFulltime int32 `xml:"PersonnelFulltime,omitempty"`

	ClassPersonnelFulltime int32 `xml:"ClassPersonnelFulltime,omitempty"`

	IndicationOrganisationcode string `xml:"IndicationOrganisationcode,omitempty"`

	IndicationEconomicallyActive bool `xml:"IndicationEconomicallyActive,omitempty"`

	IndicationNonMailing bool `xml:"IndicationNonMailing,omitempty"`

	IndicationBankruptcy bool `xml:"IndicationBankruptcy,omitempty"`

	IndicationDIP bool `xml:"IndicationDIP,omitempty"`
}

type BusinessDossierV2Array struct {
	Item []*BusinessDossierV2 `xml:"item,omitempty"`
}

type BusinessDossierV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessDossierV2Array `xml:"results,omitempty"`
}

type BusinessDossierV3 struct {
	RegisterLetter string `xml:"RegisterLetter,omitempty"`

	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	ChamberNo string `xml:"ChamberNo,omitempty"`

	Legalformcode string `xml:"Legalformcode,omitempty"`

	LegalformcodeText string `xml:"LegalformcodeText,omitempty"`

	PreviousDossierNo string `xml:"PreviousDossierNo,omitempty"`

	PreviousSubDossierNo string `xml:"PreviousSubDossierNo,omitempty"`

	Tradename45 string `xml:"Tradename45,omitempty"`

	TradenameFull string `xml:"TradenameFull,omitempty"`

	EstablishmentPostcode string `xml:"EstablishmentPostcode,omitempty"`

	EstablishmentCity string `xml:"EstablishmentCity,omitempty"`

	EstablishmentStreetname string `xml:"EstablishmentStreetname,omitempty"`

	EstablishmentHouseNo int32 `xml:"EstablishmentHouseNo,omitempty"`

	EstablishmentHouseNoAddition string `xml:"EstablishmentHouseNoAddition,omitempty"`

	CorrespondencePostcode string `xml:"CorrespondencePostcode,omitempty"`

	CorrespondenceCity string `xml:"CorrespondenceCity,omitempty"`

	CorrespondenceStreetname string `xml:"CorrespondenceStreetname,omitempty"`

	CorrespondenceHouseNo int32 `xml:"CorrespondenceHouseNo,omitempty"`

	CorrespondenceHouseNoAddition string `xml:"CorrespondenceHouseNoAddition,omitempty"`

	TelephoneNo string `xml:"TelephoneNo,omitempty"`

	Domainname string `xml:"Domainname,omitempty"`

	PrimarySBICode string `xml:"PrimarySBICode,omitempty"`

	SecondarySBICode1 string `xml:"SecondarySBICode1,omitempty"`

	SecondarySBICode2 string `xml:"SecondarySBICode2,omitempty"`

	PrimarySBICodeText string `xml:"PrimarySBICodeText,omitempty"`

	SecondarySBICode1Text string `xml:"SecondarySBICode1Text,omitempty"`

	SecondarySBICode2Text string `xml:"SecondarySBICode2Text,omitempty"`

	Personnel int32 `xml:"Personnel,omitempty"`

	ClassPersonnel int32 `xml:"ClassPersonnel,omitempty"`

	IndicationOrganisationcode string `xml:"IndicationOrganisationcode,omitempty"`

	IndicationEconomicallyActive bool `xml:"IndicationEconomicallyActive,omitempty"`

	IndicationNonMailing bool `xml:"IndicationNonMailing,omitempty"`

	IndicationBankruptcy bool `xml:"IndicationBankruptcy,omitempty"`

	IndicationDIP bool `xml:"IndicationDIP,omitempty"`
}

type BusinessDossierV3Array struct {
	Item []*BusinessDossierV3 `xml:"item,omitempty"`
}

type BusinessDossierV3PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessDossierV3Array `xml:"results,omitempty"`
}

type BusinessDossierExtended struct {
	BasicDossierPart *BusinessDossierV3 `xml:"BasicDossierPart,omitempty"`

	ExtendedDossierPart *BusinessDossierExtendedPart `xml:"ExtendedDossierPart,omitempty"`
}

type BusinessDossierExtendedPart struct {
	MainBranchIndication bool `xml:"MainBranchIndication,omitempty"`

	MainBranchDossierNo string `xml:"MainBranchDossierNo,omitempty"`

	MainBranchSubDossierNo string `xml:"MainBranchSubDossierNo,omitempty"`

	ImportIndication bool `xml:"ImportIndication,omitempty"`

	ExportIndication bool `xml:"ExportIndication,omitempty"`

	LegalName string `xml:"LegalName,omitempty"`

	ContactTitle1 string `xml:"ContactTitle1,omitempty"`

	ContactTitle2 string `xml:"ContactTitle2,omitempty"`

	ContactInitials string `xml:"ContactInitials,omitempty"`

	ContactPrefix string `xml:"ContactPrefix,omitempty"`

	ContactSurname string `xml:"ContactSurname,omitempty"`

	ContactGender string `xml:"ContactGender,omitempty"`

	AuthorizedShareCapital int64 `xml:"AuthorizedShareCapital,omitempty"`

	AuthorizedShareCapitalCurrency string `xml:"AuthorizedShareCapitalCurrency,omitempty"`

	PaidUpShareCapital int64 `xml:"PaidUpShareCapital,omitempty"`

	PaidUpShareCapitalCurrency string `xml:"PaidUpShareCapitalCurrency,omitempty"`

	IssuedShareCapital int64 `xml:"IssuedShareCapital,omitempty"`

	IssuedShareCapitalCurrency string `xml:"IssuedShareCapitalCurrency,omitempty"`

	FoundingDate *BusinessDate `xml:"FoundingDate,omitempty"`

	ContinuationDate *BusinessDate `xml:"ContinuationDate,omitempty"`

	EstablishmentDate *BusinessDate `xml:"EstablishmentDate,omitempty"`
}

type BusinessDate struct {
	Year *GYear `xml:"Year,omitempty"`

	Month int32 `xml:"Month,omitempty"`

	Day int32 `xml:"Day,omitempty"`
}

type BusinessDossierSBI struct {
	PrimarySBICode string `xml:"PrimarySBICode,omitempty"`

	SecondarySBICode1 string `xml:"SecondarySBICode1,omitempty"`

	SecondarySBICode2 string `xml:"SecondarySBICode2,omitempty"`

	PrimarySBICodeText string `xml:"PrimarySBICodeText,omitempty"`

	SecondarySBICode1Text string `xml:"SecondarySBICode1Text,omitempty"`

	SecondarySBICode2Text string `xml:"SecondarySBICode2Text,omitempty"`
}

type BusinessBIKSection struct {
	Sectioncode string `xml:"sectioncode,omitempty"`

	Description string `xml:"description,omitempty"`
}

type BusinessBIKSectionArray struct {
	Item []*BusinessBIKSection `xml:"item,omitempty"`
}

type BusinessSBISection struct {
	Sectioncode string `xml:"sectioncode,omitempty"`

	Description string `xml:"description,omitempty"`
}

type BusinessSBISectionArray struct {
	Item []*BusinessSBISection `xml:"item,omitempty"`
}

type BusinessBIKCode struct {
	Bikcode string `xml:"bikcode,omitempty"`

	Description string `xml:"description,omitempty"`
}

type BusinessBIKCodeArray struct {
	Item []*BusinessBIKCode `xml:"item,omitempty"`
}

type BusinessSBICode struct {
	Sbicode string `xml:"sbicode,omitempty"`

	Description string `xml:"description,omitempty"`
}

type BusinessSBICodeArray struct {
	Item []*BusinessSBICode `xml:"item,omitempty"`
}

type BusinessBIKCodeInfo struct {
	Sections *BusinessBIKSectionArray `xml:"sections,omitempty"`

	Bikcodes *BusinessBIKCodeArray `xml:"bikcodes,omitempty"`
}

type BusinessSBICodeInfo struct {
	Section *BusinessSBISection `xml:"section,omitempty"`

	Sbicodes *BusinessSBICodeArray `xml:"sbicodes,omitempty"`
}

type BusinessUpdateReference struct {
	DossierNo string `xml:"DossierNo,omitempty"`

	SubDossierNo string `xml:"SubDossierNo,omitempty"`

	UpdateTypes *StringArray `xml:"UpdateTypes,omitempty"`

	DateLastUpdate time.Time `xml:"DateLastUpdate,omitempty"`
}

type BusinessUpdateReferenceArray struct {
	Item []*BusinessUpdateReference `xml:"item,omitempty"`
}

type BusinessUpdateReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *BusinessUpdateReferenceArray `xml:"results,omitempty"`
}

type StringArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ stringArray"`

	Item []string `xml:"item,omitempty"`
}

type Car struct {
	License_plate string `xml:"license_plate,omitempty"`

	Category string `xml:"category,omitempty"`

	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Colors string `xml:"colors,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Cylinders int32 `xml:"cylinders,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Standing_room int32 `xml:"standing_room,omitempty"`

	Unladen_mass int32 `xml:"unladen_mass,omitempty"`

	Gross_vehicle_mass int32 `xml:"gross_vehicle_mass,omitempty"`

	Mass_ready int32 `xml:"mass_ready,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`

	Maximum_mass_unbraked int32 `xml:"maximum_mass_unbraked,omitempty"`

	Maximum_mass_braked int32 `xml:"maximum_mass_braked,omitempty"`

	Maximum_mass_trailer_braked int32 `xml:"maximum_mass_trailer_braked,omitempty"`

	Maximum_mass_self_braked int32 `xml:"maximum_mass_self_braked,omitempty"`

	Maximum_mass_axle_braked int32 `xml:"maximum_mass_axle_braked,omitempty"`

	Date_first_issuing string `xml:"date_first_issuing,omitempty"`

	Date_first_admission string `xml:"date_first_admission,omitempty"`

	Date_latest_name_registration string `xml:"date_latest_name_registration,omitempty"`

	Apk_due_date string `xml:"apk_due_date,omitempty"`
}

type CarV3 struct {
	License_plate string `xml:"license_plate,omitempty"`

	Category string `xml:"category,omitempty"`

	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Colors string `xml:"colors,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Cylinders int32 `xml:"cylinders,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Standing_room int32 `xml:"standing_room,omitempty"`

	Unladen_mass int32 `xml:"unladen_mass,omitempty"`

	Gross_vehicle_mass int32 `xml:"gross_vehicle_mass,omitempty"`

	Mass_ready int32 `xml:"mass_ready,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`

	Maximum_mass_unbraked int32 `xml:"maximum_mass_unbraked,omitempty"`

	Maximum_mass_braked int32 `xml:"maximum_mass_braked,omitempty"`

	Maximum_mass_trailer_braked int32 `xml:"maximum_mass_trailer_braked,omitempty"`

	Maximum_mass_self_braked int32 `xml:"maximum_mass_self_braked,omitempty"`

	Maximum_mass_axle_braked int32 `xml:"maximum_mass_axle_braked,omitempty"`

	Date_first_issuing string `xml:"date_first_issuing,omitempty"`

	Date_first_admission string `xml:"date_first_admission,omitempty"`

	Date_latest_name_registration string `xml:"date_latest_name_registration,omitempty"`

	Apk_due_date string `xml:"apk_due_date,omitempty"`

	Co2_emission int32 `xml:"co2_emission,omitempty"`
}

type CarV2 struct {
	Car_data *Car `xml:"car_data,omitempty"`

	Types *CarTypeArray `xml:"types,omitempty"`
}

type CarDataV3Result struct {
	Car_data *CarV3 `xml:"car_data,omitempty"`

	Types *CarTypeArray `xml:"types,omitempty"`
}

type CarBP struct {
	License_plate string `xml:"license_plate,omitempty"`

	Category string `xml:"category,omitempty"`

	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Colors string `xml:"colors,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Cylinders int32 `xml:"cylinders,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Standing_room int32 `xml:"standing_room,omitempty"`

	Unladen_mass int32 `xml:"unladen_mass,omitempty"`

	Gross_vehicle_mass int32 `xml:"gross_vehicle_mass,omitempty"`

	Mass_ready int32 `xml:"mass_ready,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`

	Maximum_mass_unbraked int32 `xml:"maximum_mass_unbraked,omitempty"`

	Maximum_mass_braked int32 `xml:"maximum_mass_braked,omitempty"`

	Maximum_mass_trailer_braked int32 `xml:"maximum_mass_trailer_braked,omitempty"`

	Maximum_mass_self_braked int32 `xml:"maximum_mass_self_braked,omitempty"`

	Maximum_mass_axle_braked int32 `xml:"maximum_mass_axle_braked,omitempty"`

	Date_first_issuing string `xml:"date_first_issuing,omitempty"`

	Date_first_admission string `xml:"date_first_admission,omitempty"`

	Date_latest_name_registration string `xml:"date_latest_name_registration,omitempty"`

	Apk_due_date string `xml:"apk_due_date,omitempty"`

	Bpm int32 `xml:"bpm,omitempty"`

	Power int32 `xml:"power,omitempty"`
}

type CarBPV2 struct {
	Car_data *CarBP `xml:"car_data,omitempty"`

	Types *CarTypeArray `xml:"types,omitempty"`
}

type CarExtended struct {
	License_plate string `xml:"license_plate,omitempty"`

	Code string `xml:"code,omitempty"`

	Check_code_status bool `xml:"check_code_status,omitempty"`

	Category string `xml:"category,omitempty"`

	Brand string `xml:"brand,omitempty"`

	Brand_code string `xml:"brand_code,omitempty"`

	Model *StringArray `xml:"model,omitempty"`

	Body_style string `xml:"body_style,omitempty"`

	Colors *StringArray `xml:"colors,omitempty"`

	Fuel_types *StringArray `xml:"fuel_types,omitempty"`

	Cylinders int32 `xml:"cylinders,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Standing_room int32 `xml:"standing_room,omitempty"`

	Unladen_mass int32 `xml:"unladen_mass,omitempty"`

	Gross_vehicle_mass int32 `xml:"gross_vehicle_mass,omitempty"`

	Mass_ready int32 `xml:"mass_ready,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`

	Maximum_mass_unbraked int32 `xml:"maximum_mass_unbraked,omitempty"`

	Maximum_mass_braked int32 `xml:"maximum_mass_braked,omitempty"`

	Maximum_mass_trailer_braked int32 `xml:"maximum_mass_trailer_braked,omitempty"`

	Maximum_mass_self_braked int32 `xml:"maximum_mass_self_braked,omitempty"`

	Maximum_mass_axle_braked int32 `xml:"maximum_mass_axle_braked,omitempty"`

	Date_first_issuing string `xml:"date_first_issuing,omitempty"`

	Date_first_admission string `xml:"date_first_admission,omitempty"`

	Date_latest_name_registration string `xml:"date_latest_name_registration,omitempty"`

	Apk_due_date string `xml:"apk_due_date,omitempty"`

	Bpm int32 `xml:"bpm,omitempty"`

	Power int32 `xml:"power,omitempty"`

	G3_indicator string `xml:"g3_indicator,omitempty"`

	Particulate_filter string `xml:"particulate_filter,omitempty"`

	Co2_emission_combined string `xml:"co2_emission_combined,omitempty"`

	Emission_code string `xml:"emission_code,omitempty"`

	Top_speed int32 `xml:"top_speed,omitempty"`

	Fuel_consumption_city string `xml:"fuel_consumption_city,omitempty"`

	Fuel_consumption_freeway string `xml:"fuel_consumption_freeway,omitempty"`

	Fuel_consumption_combined string `xml:"fuel_consumption_combined,omitempty"`

	Energy_label string `xml:"energy_label,omitempty"`

	Stolen bool `xml:"stolen,omitempty"`

	Insured bool `xml:"insured,omitempty"`

	License_plate_signal bool `xml:"license_plate_signal,omitempty"`

	Awaiting_inspection bool `xml:"awaiting_inspection,omitempty"`

	Catalog_price string `xml:"catalog_price,omitempty"`

	European_type_approval_mark string `xml:"european_type_approval_mark,omitempty"`

	Types *CarTypeArray `xml:"types,omitempty"`
}

type CarCheckCode struct {
	License_plate string `xml:"license_plate,omitempty"`

	Correct bool `xml:"correct,omitempty"`

	Active bool `xml:"active,omitempty"`
}

type CarVWEMeldcodeCheck struct {
	License_plate string `xml:"license_plate,omitempty"`

	Correct bool `xml:"correct,omitempty"`
}

type CarType struct {
	Car_id string `xml:"car_id,omitempty"`

	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Gearbox string `xml:"gearbox,omitempty"`

	First_year int32 `xml:"first_year,omitempty"`

	Last_year int32 `xml:"last_year,omitempty"`
}

type CarTypeArray struct {
	Item []*CarType `xml:"item,omitempty"`
}

type CarOptions struct {
	Car_id string `xml:"car_id,omitempty"`

	Merk string `xml:"merk,omitempty"`

	Serie string `xml:"serie,omitempty"`

	Serie_jaar_van int32 `xml:"serie_jaar_van,omitempty"`

	Serie_jaar_tot int32 `xml:"serie_jaar_tot,omitempty"`

	Model string `xml:"model,omitempty"`

	Model_maand_van int32 `xml:"model_maand_van,omitempty"`

	Model_jaar_van int32 `xml:"model_jaar_van,omitempty"`

	Model_maand_tot int32 `xml:"model_maand_tot,omitempty"`

	Model_jaar_tot int32 `xml:"model_jaar_tot,omitempty"`

	Belasting_min string `xml:"belasting_min,omitempty"`

	Belasting_max string `xml:"belasting_max,omitempty"`

	Algemene_garantie string `xml:"algemene_garantie,omitempty"`

	Algemene_garantie_km string `xml:"algemene_garantie_km,omitempty"`

	Carrosserie_garantie string `xml:"carrosserie_garantie,omitempty"`

	Deuren string `xml:"deuren,omitempty"`

	Carrosserietype string `xml:"carrosserietype,omitempty"`

	Aantal_versnellingen string `xml:"aantal_versnellingen,omitempty"`

	Soort_schakeling string `xml:"soort_schakeling,omitempty"`

	Aandrijving string `xml:"aandrijving,omitempty"`

	Brandstoftype string `xml:"brandstoftype,omitempty"`

	Aantal_cilinders string `xml:"aantal_cilinders,omitempty"`

	Bouwwijze string `xml:"bouwwijze,omitempty"`

	Kleppen_per_cilinder string `xml:"kleppen_per_cilinder,omitempty"`

	Cilinderinhoud string `xml:"cilinderinhoud,omitempty"`

	Boring string `xml:"boring,omitempty"`

	Slag string `xml:"slag,omitempty"`

	Compressieverhouding string `xml:"compressieverhouding,omitempty"`

	Vermogen_kw string `xml:"vermogen_kw,omitempty"`

	Vermogen_pk string `xml:"vermogen_pk,omitempty"`

	Vermogen_tpm string `xml:"vermogen_tpm,omitempty"`

	Koppel string `xml:"koppel,omitempty"`

	Koppel_tpm string `xml:"koppel_tpm,omitempty"`

	Katalysator string `xml:"katalysator,omitempty"`

	Brandstofsysteem string `xml:"brandstofsysteem,omitempty"`

	Klepbediening string `xml:"klepbediening,omitempty"`

	Turbo string `xml:"turbo,omitempty"`

	Wielophanging_voor string `xml:"wielophanging_voor,omitempty"`

	Wielophanging_achter string `xml:"wielophanging_achter,omitempty"`

	Vering_voor string `xml:"vering_voor,omitempty"`

	Vering_achter string `xml:"vering_achter,omitempty"`

	Stabilisator_voor string `xml:"stabilisator_voor,omitempty"`

	Stabilisator_achter string `xml:"stabilisator_achter,omitempty"`

	Remmen_voor string `xml:"remmen_voor,omitempty"`

	Remmen_voor_mm string `xml:"remmen_voor_mm,omitempty"`

	Remmen_achter string `xml:"remmen_achter,omitempty"`

	Remmen_achter_mm string `xml:"remmen_achter_mm,omitempty"`

	Cw_waarde string `xml:"cw_waarde,omitempty"`

	Lengte string `xml:"lengte,omitempty"`

	Breedte string `xml:"breedte,omitempty"`

	Hoogte string `xml:"hoogte,omitempty"`

	Wielbasis string `xml:"wielbasis,omitempty"`

	Spoorbreedte_voor string `xml:"spoorbreedte_voor,omitempty"`

	Spoorbreedte_achter string `xml:"spoorbreedte_achter,omitempty"`

	Draaicirkel string `xml:"draaicirkel,omitempty"`

	Bandenmaat_voor string `xml:"bandenmaat_voor,omitempty"`

	Bandenmaat_achter string `xml:"bandenmaat_achter,omitempty"`

	Massa string `xml:"massa,omitempty"`

	Max_toelaatbare_massa string `xml:"max_toelaatbare_massa,omitempty"`

	Max_aanhanger_geremd string `xml:"max_aanhanger_geremd,omitempty"`

	Max_aanhanger_ongeremd string `xml:"max_aanhanger_ongeremd,omitempty"`

	Max_kogeldruk string `xml:"max_kogeldruk,omitempty"`

	Max_dakbelasting string `xml:"max_dakbelasting,omitempty"`

	Koffer_min string `xml:"koffer_min,omitempty"`

	Koffer_max string `xml:"koffer_max,omitempty"`

	Tankinhoud string `xml:"tankinhoud,omitempty"`

	Topsnelheid string `xml:"topsnelheid,omitempty"`

	Acceleratie string `xml:"acceleratie,omitempty"`

	Verbruik_binnen_bebouwde_kom string `xml:"verbruik_binnen_bebouwde_kom,omitempty"`

	Verbruik_buiten_bebouwde_kom string `xml:"verbruik_buiten_bebouwde_kom,omitempty"`

	Verbruik_gecombineerd string `xml:"verbruik_gecombineerd,omitempty"`

	Co2_uitstoot string `xml:"co2_uitstoot,omitempty"`

	Verbruik_gemiddeld string `xml:"verbruik_gemiddeld,omitempty"`

	Abs *CarOption `xml:"abs,omitempty"`

	Remkrachtverdeling *CarOption `xml:"remkrachtverdeling,omitempty"`

	Brakeassist *CarOption `xml:"brakeassist,omitempty"`

	Tractiecontrole *CarOption `xml:"tractiecontrole,omitempty"`

	Sperdifferentieel *CarOption `xml:"sperdifferentieel,omitempty"`

	Stabiliteitsregeling *CarOption `xml:"stabiliteitsregeling,omitempty"`

	Regelbare_schokdemping *CarOption `xml:"regelbare_schokdemping,omitempty"`

	Niveauregeling *CarOption `xml:"niveauregeling,omitempty"`

	Airbag_bestuurder *CarOption `xml:"airbag_bestuurder,omitempty"`

	Airbag_passagier *CarOption `xml:"airbag_passagier,omitempty"`

	Airbag_opzijvoor *CarOption `xml:"airbag_opzijvoor,omitempty"`

	Airbag_opzijachter *CarOption `xml:"airbag_opzijachter,omitempty"`

	Airbag_hoofdvoor *CarOption `xml:"airbag_hoofdvoor,omitempty"`

	Airbag_hoofdachter *CarOption `xml:"airbag_hoofdachter,omitempty"`

	Gordelspanners *CarOption `xml:"gordelspanners,omitempty"`

	Verstelbare_gordelhoogte *CarOption `xml:"verstelbare_gordelhoogte,omitempty"`

	Startonderbreker *CarOption `xml:"startonderbreker,omitempty"`

	Inbraakalarm *CarOption `xml:"inbraakalarm,omitempty"`

	Centrale_deurvergrendeling *CarOption `xml:"centrale_deurvergrendeling,omitempty"`

	Afstandbediening_deurvergrendeling *CarOption `xml:"afstandbediening_deurvergrendeling,omitempty"`

	Automatische_vergrendeling *CarOption `xml:"automatische_vergrendeling,omitempty"`

	Keyless_entry *CarOption `xml:"keyless_entry,omitempty"`

	Elektrische_ramen_voor *CarOption `xml:"elektrische_ramen_voor,omitempty"`

	Elektrische_ramen_achter *CarOption `xml:"elektrische_ramen_achter,omitempty"`

	Stuurbekrachtiging *CarOption `xml:"stuurbekrachtiging,omitempty"`

	Stuurschakeling *CarOption `xml:"stuurschakeling,omitempty"`

	Cruise_control *CarOption `xml:"cruise_control,omitempty"`

	Radar_cruise_control *CarOption `xml:"radar_cruise_control,omitempty"`

	Airconditioning *CarOption `xml:"airconditioning,omitempty"`

	Klimaatregeling *CarOption `xml:"klimaatregeling,omitempty"`

	Gescheiden_temperatuurregeling *CarOption `xml:"gescheiden_temperatuurregeling,omitempty"`

	Pollenfilter *CarOption `xml:"pollenfilter,omitempty"`

	Parkeersensor *CarOption `xml:"parkeersensor,omitempty"`

	Hoofdsteunen_voor *CarOption `xml:"hoofdsteunen_voor,omitempty"`

	Hoofdsteunen_achter *CarOption `xml:"hoofdsteunen_achter,omitempty"`

	Verstelbare_hoofdsteunen_achter *CarOption `xml:"verstelbare_hoofdsteunen_achter,omitempty"`

	Derde_hoofdsteun_achter *CarOption `xml:"derde_hoofdsteun_achter,omitempty"`

	Derde_driepuntsgordel_achter *CarOption `xml:"derde_driepuntsgordel_achter,omitempty"`

	Verstelbare_lendensteun_bestuurder *CarOption `xml:"verstelbare_lendensteun_bestuurder,omitempty"`

	Verstelbare_lendensteun_passagier *CarOption `xml:"verstelbare_lendensteun_passagier,omitempty"`

	Hoogteverstelling_bestuurder *CarOption `xml:"hoogteverstelling_bestuurder,omitempty"`

	Hoogteverstelling_passagier *CarOption `xml:"hoogteverstelling_passagier,omitempty"`

	Elektrische_stoelverstelling *CarOption `xml:"elektrische_stoelverstelling,omitempty"`

	Geheugen_stoelverstelling *CarOption `xml:"geheugen_stoelverstelling,omitempty"`

	Verwarmde_zitplaatsen_voor *CarOption `xml:"verwarmde_zitplaatsen_voor,omitempty"`

	Verwarmde_zitplaatsen_achter *CarOption `xml:"verwarmde_zitplaatsen_achter,omitempty"`

	Sportstoelen *CarOption `xml:"sportstoelen,omitempty"`

	Bekledingstof_leer *CarOption `xml:"bekledingstof_leer,omitempty"`

	Lerenbekleding *CarOption `xml:"lerenbekleding,omitempty"`

	Hoogteverstelling_stuur *CarOption `xml:"hoogteverstelling_stuur,omitempty"`

	Diepteverstelling_stuur *CarOption `xml:"diepteverstelling_stuur,omitempty"`

	Elektrische_stuurverstelling *CarOption `xml:"elektrische_stuurverstelling,omitempty"`

	Leer_bekleed_stuur *CarOption `xml:"leer_bekleed_stuur,omitempty"`

	Middenarmsteun_voor *CarOption `xml:"middenarmsteun_voor,omitempty"`

	Middenarmsteun_achter *CarOption `xml:"middenarmsteun_achter,omitempty"`

	Neerklapbare_achterbank *CarOption `xml:"neerklapbare_achterbank,omitempty"`

	In_delen_neerklapbaar_achterbank *CarOption `xml:"in_delen_neerklapbaar_achterbank,omitempty"`

	Opbergvak_linkerportier *CarOption `xml:"opbergvak_linkerportier,omitempty"`

	Bekerhouder_voor *CarOption `xml:"bekerhouder_voor,omitempty"`

	Bekerhouder_achter *CarOption `xml:"bekerhouder_achter,omitempty"`

	Zonnescherm_achterruit *CarOption `xml:"zonnescherm_achterruit,omitempty"`

	Dagteller *CarOption `xml:"dagteller,omitempty"`

	Toerenteller *CarOption `xml:"toerenteller,omitempty"`

	Temperatuurmeter *CarOption `xml:"temperatuurmeter,omitempty"`

	Oliepeilmeter *CarOption `xml:"oliepeilmeter,omitempty"`

	Oliedrukmeter *CarOption `xml:"oliedrukmeter,omitempty"`

	Olietemperatuurmeter *CarOption `xml:"olietemperatuurmeter,omitempty"`

	Voltmeter *CarOption `xml:"voltmeter,omitempty"`

	Verbruiksmeter *CarOption `xml:"verbruiksmeter,omitempty"`

	Turbodrukmeter *CarOption `xml:"turbodrukmeter,omitempty"`

	Buitentemperatuurmeter *CarOption `xml:"buitentemperatuurmeter,omitempty"`

	Klokje *CarOption `xml:"klokje,omitempty"`

	Boardcomputer *CarOption `xml:"boardcomputer,omitempty"`

	Luidsprekers *CarOption `xml:"luidsprekers,omitempty"`

	Stereoinstallatie *CarOption `xml:"stereoinstallatie,omitempty"`

	Cdspeler *CarOption `xml:"cdspeler,omitempty"`

	Cdwisselaar *CarOption `xml:"cdwisselaar,omitempty"`

	Audiostuurbediening *CarOption `xml:"audiostuurbediening,omitempty"`

	Bandenspanningssensor *CarOption `xml:"bandenspanningssensor,omitempty"`

	Navigatiesysteem *CarOption `xml:"navigatiesysteem,omitempty"`

	Televisie *CarOption `xml:"televisie,omitempty"`

	Telefoonvoorbereiding *CarOption `xml:"telefoonvoorbereiding,omitempty"`

	Telefoon *CarOption `xml:"telefoon,omitempty"`

	Intervalruitenwisser *CarOption `xml:"intervalruitenwisser,omitempty"`

	Regelbare_interval *CarOption `xml:"regelbare_interval,omitempty"`

	Regensensor *CarOption `xml:"regensensor,omitempty"`

	Ruitenwisser_achter *CarOption `xml:"ruitenwisser_achter,omitempty"`

	Achterruitverwarming *CarOption `xml:"achterruitverwarming,omitempty"`

	Voorruitverwarming *CarOption `xml:"voorruitverwarming,omitempty"`

	Verwarmde_ruitensproeiers *CarOption `xml:"verwarmde_ruitensproeiers,omitempty"`

	Gelaagde_voorruit *CarOption `xml:"gelaagde_voorruit,omitempty"`

	Getintglas *CarOption `xml:"getintglas,omitempty"`

	Schuifdak *CarOption `xml:"schuifdak,omitempty"`

	Elektrisch_schuifdak *CarOption `xml:"elektrisch_schuifdak,omitempty"`

	Beschermstrips_opzij *CarOption `xml:"beschermstrips_opzij,omitempty"`

	Bumpers_meegespoten *CarOption `xml:"bumpers_meegespoten,omitempty"`

	Metallic_lak *CarOption `xml:"metallic_lak,omitempty"`

	Lichtmetalen_velgen *CarOption `xml:"lichtmetalen_velgen,omitempty"`

	Rechter_buitenspiegel *CarOption `xml:"rechter_buitenspiegel,omitempty"`

	Binnenuit_verstelbare_buitenspiegel *CarOption `xml:"binnenuit_verstelbare_buitenspiegel,omitempty"`

	Elektrische_spiegels *CarOption `xml:"elektrische_spiegels,omitempty"`

	Verwarmde_spiegels *CarOption `xml:"verwarmde_spiegels,omitempty"`

	Inklapbare_spiegels *CarOption `xml:"inklapbare_spiegels,omitempty"`

	Meegespoten_spiegels *CarOption `xml:"meegespoten_spiegels,omitempty"`

	Dodehoekspiegel_bestuurder *CarOption `xml:"dodehoekspiegel_bestuurder,omitempty"`

	Dimmende_binnenspiegel *CarOption `xml:"dimmende_binnenspiegel,omitempty"`

	Dimmende_buitenspiegel *CarOption `xml:"dimmende_buitenspiegel,omitempty"`

	Dakrails *CarOption `xml:"dakrails,omitempty"`

	Binnenuit_te_openen_tankklep *CarOption `xml:"binnenuit_te_openen_tankklep,omitempty"`

	Binnenuit_te_openen_bagageklep *CarOption `xml:"binnenuit_te_openen_bagageklep,omitempty"`

	Op_afstand_te_openen_bagageklep *CarOption `xml:"op_afstand_te_openen_bagageklep,omitempty"`

	Halogeen_koplampen *CarOption `xml:"halogeen_koplampen,omitempty"`

	Xenon_koplampen *CarOption `xml:"xenon_koplampen,omitempty"`

	Verstelbare_koplampen *CarOption `xml:"verstelbare_koplampen,omitempty"`

	Koplampsproeiers *CarOption `xml:"koplampsproeiers,omitempty"`

	Automatisch_inschakelende_koplampen *CarOption `xml:"automatisch_inschakelende_koplampen,omitempty"`

	Verstralers *CarOption `xml:"verstralers,omitempty"`

	Mistlampen_voor *CarOption `xml:"mistlampen_voor,omitempty"`

	Mistlampen_achter *CarOption `xml:"mistlampen_achter,omitempty"`

	Zoemer_vergeten_verlichting *CarOption `xml:"zoemer_vergeten_verlichting,omitempty"`

	Verlichte_bagageruimte *CarOption `xml:"verlichte_bagageruimte,omitempty"`

	Verlichte_motorruimte *CarOption `xml:"verlichte_motorruimte,omitempty"`

	Verlicht_dashboardkastje *CarOption `xml:"verlicht_dashboardkastje,omitempty"`

	Regelbare_dashboardverlichting *CarOption `xml:"regelbare_dashboardverlichting,omitempty"`

	Leeslampje_voor *CarOption `xml:"leeslampje_voor,omitempty"`

	Leeslampje_achter *CarOption `xml:"leeslampje_achter,omitempty"`

	Verlichte_makeupspiegel *CarOption `xml:"verlichte_makeupspiegel,omitempty"`

	Vertraging_interieurverlichting *CarOption `xml:"vertraging_interieurverlichting,omitempty"`

	Portierverlichting *CarOption `xml:"portierverlichting,omitempty"`
}

type CarOption struct {
	Standaard bool `xml:"standaard,omitempty"`

	Pakket int32 `xml:"pakket,omitempty"`

	Standaard_sinds_datum *GYearMonth `xml:"standaard_sinds_datum,omitempty"`

	Prijs int32 `xml:"prijs,omitempty"`
}

type CarVWEPrices struct {
	Current_price_retail int32 `xml:"current_price_retail,omitempty"`

	Current_price_exchange int32 `xml:"current_price_exchange,omitempty"`

	Current_price_trade int32 `xml:"current_price_trade,omitempty"`

	Date_price_list time.Time `xml:"date_price_list,omitempty"`

	Catalog_price int32 `xml:"catalog_price,omitempty"`

	Btw_new_price int32 `xml:"btw_new_price,omitempty"`

	Bpm_new_price int32 `xml:"bpm_new_price,omitempty"`

	Net_catalog_price int32 `xml:"net_catalog_price,omitempty"`

	Recalculated_catalog_price int32 `xml:"recalculated_catalog_price,omitempty"`

	Recalculated_btw_new_price int32 `xml:"recalculated_btw_new_price,omitempty"`

	Recalculated_net_catalog_price int32 `xml:"recalculated_net_catalog_price,omitempty"`
}

type CarVWEBasicTypeData struct {
	Category string `xml:"category,omitempty"`

	Brand_rdw string `xml:"brand_rdw,omitempty"`

	Brand_atl string `xml:"brand_atl,omitempty"`

	Model_rdw string `xml:"model_rdw,omitempty"`

	Model_atl string `xml:"model_atl,omitempty"`

	Short_type string `xml:"short_type,omitempty"`

	Colors *StringArray `xml:"colors,omitempty"`

	Fuel_types *StringArray `xml:"fuel_types,omitempty"`

	Cylinders int32 `xml:"cylinders,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Standing_room int32 `xml:"standing_room,omitempty"`

	Unladen_mass int32 `xml:"unladen_mass,omitempty"`

	Gross_vehicle_mass int32 `xml:"gross_vehicle_mass,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`

	Date_first_issuing time.Time `xml:"date_first_issuing,omitempty"`

	Date_first_admission time.Time `xml:"date_first_admission,omitempty"`

	Date_latest_name_registration time.Time `xml:"date_latest_name_registration,omitempty"`

	Apk_due_date time.Time `xml:"apk_due_date,omitempty"`

	Bpm int32 `xml:"bpm,omitempty"`

	Bpm_currency string `xml:"bpm_currency,omitempty"`

	Power int32 `xml:"power,omitempty"`

	G3_indication bool `xml:"g3_indication,omitempty"`

	Co2_emission int32 `xml:"co2_emission,omitempty"`

	Emission_particles string `xml:"emission_particles,omitempty"`

	Energy_label string `xml:"energy_label,omitempty"`

	Emission_code string `xml:"emission_code,omitempty"`

	Interior_code string `xml:"interior_code,omitempty"`

	Interior_description string `xml:"interior_description,omitempty"`

	Doors int32 `xml:"doors,omitempty"`

	Tax_write_off string `xml:"tax_write_off,omitempty"`

	Parallel_import bool `xml:"parallel_import,omitempty"`

	License_plate_signal string `xml:"license_plate_signal,omitempty"`

	Gear string `xml:"gear,omitempty"`

	Brand_approval string `xml:"brand_approval,omitempty"`

	Top_speed int32 `xml:"top_speed,omitempty"`

	Types *CarVWEVersionPriceReferenceArray `xml:"types,omitempty"`
}

type CarVWEVersionPrice struct {
	Brand_atl string `xml:"brand_atl,omitempty"`

	Model_atl string `xml:"model_atl,omitempty"`

	Short_type string `xml:"short_type,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Power int32 `xml:"power,omitempty"`

	Energy_label string `xml:"energy_label,omitempty"`

	Emission_code string `xml:"emission_code,omitempty"`

	Doors int32 `xml:"doors,omitempty"`

	Gear string `xml:"gear,omitempty"`

	Top_speed int32 `xml:"top_speed,omitempty"`

	Folder_weight int32 `xml:"folder_weight,omitempty"`

	Train_weight int32 `xml:"train_weight,omitempty"`

	Layout string `xml:"layout,omitempty"`

	Design string `xml:"design,omitempty"`

	Acceleration float64 `xml:"acceleration,omitempty"`

	Sportivity float32 `xml:"sportivity,omitempty"`

	Version *CarVWEVersionPriceReference `xml:"version,omitempty"`

	Prices *CarVWEPrices `xml:"prices,omitempty"`
}

type CarVWEVersionYearData struct {
	Atl_code int32 `xml:"atl_code,omitempty"`

	Brand_atl string `xml:"brand_atl,omitempty"`

	Model_atl string `xml:"model_atl,omitempty"`

	Version_name string `xml:"version_name,omitempty"`

	Short_type string `xml:"short_type,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Power int32 `xml:"power,omitempty"`

	Energy_label string `xml:"energy_label,omitempty"`

	Emission_code string `xml:"emission_code,omitempty"`

	Doors int32 `xml:"doors,omitempty"`

	Gear string `xml:"gear,omitempty"`

	Top_speed int32 `xml:"top_speed,omitempty"`

	Folder_weight int32 `xml:"folder_weight,omitempty"`

	Train_weight int32 `xml:"train_weight,omitempty"`

	Layout string `xml:"layout,omitempty"`

	Body_style string `xml:"body_style,omitempty"`

	Acceleration float64 `xml:"acceleration,omitempty"`

	Sportivity float32 `xml:"sportivity,omitempty"`

	Prices *CarVWEPrices `xml:"prices,omitempty"`
}

type CarVWEVersionPriceReference struct {
	Order_number int32 `xml:"order_number,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`

	Model_name string `xml:"model_name,omitempty"`
}

type CarVWEVersionPriceReferenceArray struct {
	Item []*CarVWEVersionPriceReference `xml:"item,omitempty"`
}

type CarVWEBrand struct {
	Brand_id int32 `xml:"brand_id,omitempty"`

	Brand_name string `xml:"brand_name,omitempty"`
}

type CarVWEBrandArray struct {
	Item []*CarVWEBrand `xml:"item,omitempty"`
}

type CarVWEBrandPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CarVWEBrandArray `xml:"results,omitempty"`
}

type CarVWEModel struct {
	Model_id int32 `xml:"model_id,omitempty"`

	Model_name string `xml:"model_name,omitempty"`
}

type CarVWEModelArray struct {
	Item []*CarVWEModel `xml:"item,omitempty"`
}

type CarVWEModelPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CarVWEModelArray `xml:"results,omitempty"`
}

type CarVWEVersion struct {
	Atl_code int32 `xml:"atl_code,omitempty"`

	Version_name string `xml:"version_name,omitempty"`

	Fuel_type_id int32 `xml:"fuel_type_id,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Body_style_id int32 `xml:"body_style_id,omitempty"`

	Body_style string `xml:"body_style,omitempty"`

	Doors int32 `xml:"doors,omitempty"`

	Power int32 `xml:"power,omitempty"`

	Gear_id int32 `xml:"gear_id,omitempty"`

	Gear string `xml:"gear,omitempty"`

	Valid_from time.Time `xml:"valid_from,omitempty"`

	Valid_until time.Time `xml:"valid_until,omitempty"`
}

type CarVWEVersionArray struct {
	Item []*CarVWEVersion `xml:"item,omitempty"`
}

type CarVWEVersionPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CarVWEVersionArray `xml:"results,omitempty"`
}

type CarVWEPhoto struct {
	Image []byte `xml:"image,omitempty"`

	Side string `xml:"side,omitempty"`

	Size string `xml:"size,omitempty"`
}

type CarVWEPhotoArray struct {
	Item []*CarVWEPhoto `xml:"item,omitempty"`
}

type CarVWEOptions struct {
	Options_standard *CarVWEOptionsStandard `xml:"options_standard,omitempty"`

	Options_manufacture *CarVWEOptionsManufacture `xml:"options_manufacture,omitempty"`

	Options_package *CarVWEOptionsPackage `xml:"options_package,omitempty"`
}

type CarVWEOptionsStandard struct {
	License_plate string `xml:"license_plate,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`

	Options *CarVWEOptionArray `xml:"options,omitempty"`
}

type CarVWEOptionsManufacture struct {
	License_plate string `xml:"license_plate,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`

	Options *CarVWEOptionArray `xml:"options,omitempty"`
}

type CarVWEOptionsPackage struct {
	License_plate string `xml:"license_plate,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`

	Packages *CarVWEOptionPackageArray `xml:"packages,omitempty"`
}

type CarVWEOptionPackage struct {
	Package_id int32 `xml:"package_id,omitempty"`

	Price int32 `xml:"price,omitempty"`

	Description string `xml:"description,omitempty"`

	Options *CarVWEOptionArray `xml:"options,omitempty"`
}

type CarVWEOptionPackageArray struct {
	Item []*CarVWEOptionPackage `xml:"item,omitempty"`
}

type CarVWEOption struct {
	Option_id int32 `xml:"option_id,omitempty"`

	Description_short string `xml:"description_short,omitempty"`

	Description string `xml:"description,omitempty"`

	Price int32 `xml:"price,omitempty"`
}

type CarVWEOptionArray struct {
	Item []*CarVWEOption `xml:"item,omitempty"`
}

type CarRDWCarDataPrice struct {
	Basic_car_data *CarV3 `xml:"basic_car_data,omitempty"`

	Environment_data *CarEnvironment `xml:"environment_data,omitempty"`

	Status_data *CarStatus `xml:"status_data,omitempty"`

	Power float64 `xml:"power,omitempty"`

	Fuel_types *StringArray `xml:"fuel_types,omitempty"`

	Top_speed int32 `xml:"top_speed,omitempty"`

	Catalog_price int32 `xml:"catalog_price,omitempty"`

	Bpm int32 `xml:"bpm,omitempty"`

	Bodies *CarBodyArray `xml:"bodies,omitempty"`
}

type CarStatus struct {
	Insured bool `xml:"insured,omitempty"`

	Awaiting_inspection bool `xml:"awaiting_inspection,omitempty"`

	Missing bool `xml:"missing,omitempty"`

	Stolen bool `xml:"stolen,omitempty"`
}

type CarEnvironment struct {
	Emission_code string `xml:"emission_code,omitempty"`

	Energy_label string `xml:"energy_label,omitempty"`

	G3_indication bool `xml:"g3_indication,omitempty"`

	Particulate_filter string `xml:"particulate_filter,omitempty"`

	Fuel_usage_city float64 `xml:"fuel_usage_city,omitempty"`

	Fuel_usage_highway float64 `xml:"fuel_usage_highway,omitempty"`

	Fuel_usage_combined float64 `xml:"fuel_usage_combined,omitempty"`
}

type CarBody struct {
	Id int32 `xml:"id,omitempty"`

	Code string `xml:"code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type CarBodyArray struct {
	Item []*CarBody `xml:"item,omitempty"`
}

type CarATDPrices struct {
	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Valuations *CarATDValuationArray `xml:"valuations,omitempty"`
}

type CarATDValuation struct {
	Type_id int32 `xml:"type_id,omitempty"`

	Price_retail int32 `xml:"price_retail,omitempty"`

	Price_exchange int32 `xml:"price_exchange,omitempty"`

	Price_trade int32 `xml:"price_trade,omitempty"`

	Type_names *StringArray `xml:"type_names,omitempty"`
}

type CarATDValuationArray struct {
	Item []*CarATDValuation `xml:"item,omitempty"`
}

type CompliancePersonReference struct {
	Compliance_person_id int32 `xml:"compliance_person_id,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Relation_description string `xml:"relation_description,omitempty"`
}

type CompliancePersonReferenceArray struct {
	Item []*CompliancePersonReference `xml:"item,omitempty"`
}

type CompliancePersonSearchReference struct {
	Match_score int32 `xml:"match_score,omitempty"`

	Compliance_person_id int32 `xml:"compliance_person_id,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Nationality string `xml:"nationality,omitempty"`

	Flags *StringArray `xml:"flags,omitempty"`
}

type CompliancePersonSearchReferenceArray struct {
	Item []*CompliancePersonSearchReference `xml:"item,omitempty"`
}

type CompliancePersonSearchReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CompliancePersonSearchReferenceArray `xml:"results,omitempty"`
}

type CompliancePerson struct {
	Compliance_person_id int32 `xml:"compliance_person_id,omitempty"`

	Pep_level int32 `xml:"pep_level,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Middle_name string `xml:"middle_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Gender string `xml:"gender,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Date_of_death time.Time `xml:"date_of_death,omitempty"`

	Deceased bool `xml:"deceased,omitempty"`

	Nationality string `xml:"nationality,omitempty"`

	Image_url string `xml:"image_url,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Mobile_number string `xml:"mobile_number,omitempty"`

	Email string `xml:"email,omitempty"`

	Sanctions *ComplianceSanctionReferenceArray `xml:"sanctions,omitempty"`

	Flags *StringArray `xml:"flags,omitempty"`

	Aliases *StringArray `xml:"aliases,omitempty"`

	Addresses *ComplianceAddressReferenceArray `xml:"addresses,omitempty"`

	Roles *ComplianceRoleReferenceArray `xml:"roles,omitempty"`

	Related_businesses *ComplianceBusinessReferenceArray `xml:"related_businesses,omitempty"`

	Related_persons *CompliancePersonReferenceArray `xml:"related_persons,omitempty"`

	Notes *ComplianceNoteReferenceArray `xml:"notes,omitempty"`

	Documents *ComplianceDocumentReferenceArray `xml:"documents,omitempty"`
}

type ComplianceBusiness struct {
	Compliance_business_id int32 `xml:"compliance_business_id,omitempty"`

	Business_name string `xml:"business_name,omitempty"`

	Website_url string `xml:"website_url,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Sanctions *ComplianceSanctionReferenceArray `xml:"sanctions,omitempty"`

	Flags *StringArray `xml:"flags,omitempty"`

	Aliases *StringArray `xml:"aliases,omitempty"`

	Addresses *ComplianceAddressReferenceArray `xml:"addresses,omitempty"`

	Related_businesses *ComplianceBusinessReferenceArray `xml:"related_businesses,omitempty"`

	Related_persons *CompliancePersonReferenceArray `xml:"related_persons,omitempty"`

	Notes *ComplianceNoteReferenceArray `xml:"notes,omitempty"`

	Documents *ComplianceDocumentReferenceArray `xml:"documents,omitempty"`
}

type ComplianceAddressReference struct {
	Address_lines *StringArray `xml:"address_lines,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	City string `xml:"city,omitempty"`

	County string `xml:"county,omitempty"`

	Country string `xml:"country,omitempty"`
}

type ComplianceAddressReferenceArray struct {
	Item []*ComplianceAddressReference `xml:"item,omitempty"`
}

type ComplianceRoleReference struct {
	Role string `xml:"role,omitempty"`

	Country string `xml:"country,omitempty"`

	Start_date *ComplianceDate `xml:"start_date,omitempty"`

	End_date *ComplianceDate `xml:"end_date,omitempty"`
}

type ComplianceRoleReferenceArray struct {
	Item []*ComplianceRoleReference `xml:"item,omitempty"`
}

type ComplianceDate struct {
	Year int32 `xml:"year,omitempty"`

	Month int32 `xml:"month,omitempty"`

	Day int32 `xml:"day,omitempty"`
}

type ComplianceBusinessReference struct {
	Compliance_business_id int32 `xml:"compliance_business_id,omitempty"`

	Name string `xml:"name,omitempty"`

	Relation_description string `xml:"relation_description,omitempty"`
}

type ComplianceBusinessReferenceArray struct {
	Item []*ComplianceBusinessReference `xml:"item,omitempty"`
}

type ComplianceBusinessSearchReference struct {
	Match_score int32 `xml:"match_score,omitempty"`

	Compliance_business_id int32 `xml:"compliance_business_id,omitempty"`

	Name string `xml:"name,omitempty"`

	Country string `xml:"country,omitempty"`

	Flags *StringArray `xml:"flags,omitempty"`
}

type ComplianceBusinessSearchReferenceArray struct {
	Item []*ComplianceBusinessSearchReference `xml:"item,omitempty"`
}

type ComplianceBusinessSearchReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *ComplianceBusinessSearchReferenceArray `xml:"results,omitempty"`
}

type ComplianceDocumentReference struct {
	Original_url string `xml:"original_url,omitempty"`

	Cached_url string `xml:"cached_url,omitempty"`

	Date_collected time.Time `xml:"date_collected,omitempty"`

	Categories *StringArray `xml:"categories,omitempty"`
}

type ComplianceDocumentReferenceArray struct {
	Item []*ComplianceDocumentReference `xml:"item,omitempty"`
}

type ComplianceSanctionReference struct {
	Organisation string `xml:"organisation,omitempty"`

	Active bool `xml:"active,omitempty"`
}

type ComplianceSanctionReferenceArray struct {
	Item []*ComplianceSanctionReference `xml:"item,omitempty"`
}

type ComplianceNoteReference struct {
	Source string `xml:"source,omitempty"`

	Text string `xml:"text,omitempty"`
}

type ComplianceNoteReferenceArray struct {
	Item []*ComplianceNoteReference `xml:"item,omitempty"`
}

type CreditsafeCompany struct {
	Name string `xml:"name,omitempty"`

	Registration_number string `xml:"registration_number,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Status string `xml:"status,omitempty"`

	Office_type string `xml:"office_type,omitempty"`

	Address string `xml:"address,omitempty"`

	Available_report_types *StringArray `xml:"available_report_types,omitempty"`

	Available_languages *StringArray `xml:"available_languages,omitempty"`

	Id string `xml:"id,omitempty"`

	Country string `xml:"country,omitempty"`

	Online_reports bool `xml:"online_reports,omitempty"`
}

type CreditsafeCompanyArray struct {
	Item []*CreditsafeCompany `xml:"item,omitempty"`
}

type CreditsafeCompanyPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CreditsafeCompanyArray `xml:"results,omitempty"`
}

type CreditsafeCompanyReportFull struct {
	Company_id string `xml:"company_id,omitempty"`

	Order_number uint64 `xml:"order_number,omitempty"`

	Language string `xml:"language,omitempty"`

	Company_summary *CreditsafeLtdCompanySummary `xml:"company_summary,omitempty"`

	Company_identification *CreditsafeLtdCompanyIdentification `xml:"company_identification,omitempty"`

	Credit_score *CreditsafeLtdCreditScore `xml:"credit_score,omitempty"`

	Contact_information *CreditsafeLtdContactInformation `xml:"contact_information,omitempty"`

	Share_capital_structure *CreditsafeLtdShareCapitalStructure `xml:"share_capital_structure,omitempty"`

	Directors *CreditsafeLtdDirectors `xml:"directors,omitempty"`

	Other_information *CreditsafeLtdOtherInformation `xml:"other_information,omitempty"`

	Group_structure *CreditsafeLtdGroupStructure `xml:"group_structure,omitempty"`

	Financial_statements *CreditsafeLtdFinancialStatementArray `xml:"financial_statements,omitempty"`

	Additional_information *CreditsafeAdditionalInformation `xml:"additional_information,omitempty"`

	Document []byte `xml:"document,omitempty"`
}

type CreditsafeCompanyV2 struct {
	Address *CreditsafeAddressV2 `xml:"address,omitempty"`

	Creditsafe_number string `xml:"creditsafe_number,omitempty"`

	Date_last_account time.Time `xml:"date_last_account,omitempty"`

	Date_last_change time.Time `xml:"date_last_change,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Vat_numbers *StringArray `xml:"vat_numbers,omitempty"`

	Country string `xml:"country,omitempty"`

	Registration_number string `xml:"registration_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Id string `xml:"id,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Status string `xml:"status,omitempty"`

	Office_type string `xml:"office_type,omitempty"`
}

type CreditsafeLtdCompanySummary struct {
	Business_name string `xml:"business_name,omitempty"`

	Country string `xml:"country,omitempty"`

	Number string `xml:"number,omitempty"`

	Company_registration_number string `xml:"company_registration_number,omitempty"`

	Main_activity *CreditsafeCompanyActivity `xml:"main_activity,omitempty"`

	Company_status string `xml:"company_status,omitempty"`

	Latest_turnover_figure float64 `xml:"latest_turnover_figure,omitempty"`

	Latest_shareholders_equity_figure float64 `xml:"latest_shareholders_equity_figure,omitempty"`

	Credit_rating string `xml:"credit_rating,omitempty"`

	Credit_limit string `xml:"credit_limit,omitempty"`

	Rating_description string `xml:"rating_description,omitempty"`
}

type CreditsafeLtdCompanyIdentification struct {
	Basic_information *CreditsafeLtdCompanyBasicInformation `xml:"basic_information,omitempty"`

	Activities *CreditsafeCompanyActivityArray `xml:"activities,omitempty"`

	Previous_names *CreditsafeChangedStringArray `xml:"previous_names,omitempty"`

	Previous_legal_forms *CreditsafeChangedStringArray `xml:"previous_legal_forms,omitempty"`
}

type CreditsafeChangedString struct {
	Value string `xml:"value,omitempty"`

	Date_changed time.Time `xml:"date_changed,omitempty"`
}

type CreditsafeChangedStringArray struct {
	Item []*CreditsafeChangedString `xml:"item,omitempty"`
}

type CreditsafeLtdCreditScore struct {
	Current_credit_rating string `xml:"current_credit_rating,omitempty"`

	Current_rating_description string `xml:"current_rating_description,omitempty"`

	Current_credit_limit string `xml:"current_credit_limit,omitempty"`

	Current_contract_limit float64 `xml:"current_contract_limit,omitempty"`

	Previous_credit_limit string `xml:"previous_credit_limit,omitempty"`

	Previous_credit_rating string `xml:"previous_credit_rating,omitempty"`

	Previous_rating_description string `xml:"previous_rating_description,omitempty"`

	Date_of_latest_rating_change time.Time `xml:"date_of_latest_rating_change,omitempty"`
}

type CreditsafeLtdContactInformation struct {
	Main_address *CreditsafeStreetAddressWithTelephone `xml:"main_address,omitempty"`

	Other_addresses *CreditsafeStreetAddressWithTelephoneArray `xml:"other_addresses,omitempty"`

	Previous_addresses *CreditsafeStreetAddressArray `xml:"previous_addresses,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Websites *StringArray `xml:"websites,omitempty"`
}

type CreditsafeLtdShareCapitalStructure struct {
	Nominal_share_capital float64 `xml:"nominal_share_capital,omitempty"`

	Issued_share_capital float64 `xml:"issued_share_capital,omitempty"`

	Share_holders *CreditsafeShareHolderArray `xml:"share_holders,omitempty"`
}

type CreditsafeLtdDirectors struct {
	Current_directors *CreditsafeDirectorArray `xml:"current_directors,omitempty"`

	Previous_directors *CreditsafePreviousDirectorArray `xml:"previous_directors,omitempty"`
}

type CreditsafeLtdOtherInformation struct {
	Bankers *CreditsafeBankerArray `xml:"bankers,omitempty"`

	Advisors *CreditsafeAdvisorArray `xml:"advisors,omitempty"`

	Employees_information *CreditsafeEmployeeInformationArray `xml:"employees_information,omitempty"`
}

type CreditsafeLtdGroupStructure struct {
	Ultimate_parent *CreditsafeCompanyListEntry `xml:"ultimate_parent,omitempty"`

	Immediate_parent *CreditsafeCompanyListEntry `xml:"immediate_parent,omitempty"`

	Subsidiary_companies *CreditsafeCompanyListEntryArray `xml:"subsidiary_companies,omitempty"`

	Affiliated_companies *CreditsafeCompanyListEntryArray `xml:"affiliated_companies,omitempty"`
}

type CreditsafeLtdCompanyBasicInformation struct {
	Business_name string `xml:"business_name,omitempty"`

	Registered_company_name string `xml:"registered_company_name,omitempty"`

	Company_registration_number string `xml:"company_registration_number,omitempty"`

	Country string `xml:"country,omitempty"`

	Vat_registration_number string `xml:"vat_registration_number,omitempty"`

	Vat_registration_date time.Time `xml:"vat_registration_date,omitempty"`

	Dateof_company_registration time.Time `xml:"dateof_company_registration,omitempty"`

	Dateof_starting_operations time.Time `xml:"dateof_starting_operations,omitempty"`

	Commercial_court string `xml:"commercial_court,omitempty"`

	Legal_form string `xml:"legal_form,omitempty"`

	Typeof_ownership string `xml:"typeof_ownership,omitempty"`

	Company_status string `xml:"company_status,omitempty"`

	Report_currency string `xml:"report_currency,omitempty"`

	Principal_activity *CreditsafeCompanyActivity `xml:"principal_activity,omitempty"`

	Contact_address string `xml:"contact_address,omitempty"`

	Contact_telephone_number string `xml:"contact_telephone_number,omitempty"`
}

type CreditsafeStreetAddressWithTelephone struct {
	Address string `xml:"address,omitempty"`

	Country string `xml:"country,omitempty"`

	Telephone string `xml:"telephone,omitempty"`
}

type CreditsafeStreetAddressWithTelephoneArray struct {
	Item []*CreditsafeStreetAddressWithTelephone `xml:"item,omitempty"`
}

type CreditsafeCompanyListEntry struct {
	Country string `xml:"country,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Company_id string `xml:"company_id,omitempty"`

	Number string `xml:"number,omitempty"`
}

type CreditsafeCompanyListEntryArray struct {
	Item []*CreditsafeCompanyListEntry `xml:"item,omitempty"`
}

type CreditsafeLtdFinancialStatement struct {
	Profit_and_loss *CreditsafeProfitAndLossFigures `xml:"profit_and_loss,omitempty"`

	Balance_sheet *CreditsafeBalanceSheet `xml:"balance_sheet,omitempty"`

	Other_financials *CreditsafeOtherFinancials `xml:"other_financials,omitempty"`

	Ratios *CreditsafeFinancialRatios `xml:"ratios,omitempty"`
}

type CreditsafeLtdFinancialStatementArray struct {
	Item []*CreditsafeLtdFinancialStatement `xml:"item,omitempty"`
}

type CreditsafeCompanyActivity struct {
	Activity_code string `xml:"activity_code,omitempty"`

	Activity_description string `xml:"activity_description,omitempty"`
}

type CreditsafeCompanyActivityArray struct {
	Item []*CreditsafeCompanyActivity `xml:"item,omitempty"`
}

type CreditsafeStreetAddress struct {
	Address string `xml:"address,omitempty"`

	Country string `xml:"country,omitempty"`
}

type CreditsafeStreetAddressArray struct {
	Item []*CreditsafeStreetAddress `xml:"item,omitempty"`
}

type CreditsafeShareHolder struct {
	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`

	Share_percent float64 `xml:"share_percent,omitempty"`
}

type CreditsafeShareHolderArray struct {
	Item []*CreditsafeShareHolder `xml:"item,omitempty"`
}

type CreditsafeDirector struct {
	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`

	Gender string `xml:"gender,omitempty"`

	Appointment_date time.Time `xml:"appointment_date,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Position string `xml:"position,omitempty"`
}

type CreditsafeDirectorArray struct {
	Item []*CreditsafeDirector `xml:"item,omitempty"`
}

type CreditsafePreviousDirector struct {
	Director *CreditsafeDirector `xml:"director,omitempty"`

	Resignation_date time.Time `xml:"resignation_date,omitempty"`
}

type CreditsafePreviousDirectorArray struct {
	Item []*CreditsafePreviousDirector `xml:"item,omitempty"`
}

type CreditsafeBanker struct {
	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`

	Bank_code string `xml:"bank_code,omitempty"`
}

type CreditsafeBankerArray struct {
	Item []*CreditsafeBanker `xml:"item,omitempty"`
}

type CreditsafeAdvisor struct {
	Auditor_name string `xml:"auditor_name,omitempty"`

	Solicitor_name string `xml:"solicitor_name,omitempty"`
}

type CreditsafeAdvisorArray struct {
	Item []*CreditsafeAdvisor `xml:"item,omitempty"`
}

type CreditsafeEmployeeInformation struct {
	Year uint16 `xml:"year,omitempty"`

	Number_of_employees string `xml:"number_of_employees,omitempty"`
}

type CreditsafeEmployeeInformationArray struct {
	Item []*CreditsafeEmployeeInformation `xml:"item,omitempty"`
}

type CreditsafeProfitAndLossFigures struct {
	Year_end_date time.Time `xml:"year_end_date,omitempty"`

	Number_of_weeks uint32 `xml:"number_of_weeks,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Consolidated_accounts bool `xml:"consolidated_accounts,omitempty"`

	Revenue float64 `xml:"revenue,omitempty"`

	Operating_costs float64 `xml:"operating_costs,omitempty"`

	Operating_profit float64 `xml:"operating_profit,omitempty"`

	Wages_and_salaries float64 `xml:"wages_and_salaries,omitempty"`

	Pension_costs float64 `xml:"pension_costs,omitempty"`

	Depreciation float64 `xml:"depreciation,omitempty"`

	Amortisation float64 `xml:"amortisation,omitempty"`

	Financial_income float64 `xml:"financial_income,omitempty"`

	Financial_expenses float64 `xml:"financial_expenses,omitempty"`

	Extraordinary_income float64 `xml:"extraordinary_income,omitempty"`

	Extraordinary_costs float64 `xml:"extraordinary_costs,omitempty"`

	Profit_before_tax float64 `xml:"profit_before_tax,omitempty"`

	Tax float64 `xml:"tax,omitempty"`

	Profit_after_tax float64 `xml:"profit_after_tax,omitempty"`

	Dividends float64 `xml:"dividends,omitempty"`

	Minority_interests float64 `xml:"minority_interests,omitempty"`

	Other_appropriations float64 `xml:"other_appropriations,omitempty"`

	Retained_profit float64 `xml:"retained_profit,omitempty"`
}

type CreditsafeBalanceSheet struct {
	Year_end_date time.Time `xml:"year_end_date,omitempty"`

	Number_of_weeks uint32 `xml:"number_of_weeks,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Consolidated_accounts bool `xml:"consolidated_accounts,omitempty"`

	Land_and_buildings float64 `xml:"land_and_buildings,omitempty"`

	Plant_and_machinery float64 `xml:"plant_and_machinery,omitempty"`

	Other_tangible_assets float64 `xml:"other_tangible_assets,omitempty"`

	Total_tangible_assets float64 `xml:"total_tangible_assets,omitempty"`

	Goodwill float64 `xml:"goodwill,omitempty"`

	Other_intangible_assets float64 `xml:"other_intangible_assets,omitempty"`

	Total_intangible_assets float64 `xml:"total_intangible_assets,omitempty"`

	Investments float64 `xml:"investments,omitempty"`

	Loans_to_group float64 `xml:"loans_to_group,omitempty"`

	Other_loans float64 `xml:"other_loans,omitempty"`

	Miscellaneous_fixed_assets float64 `xml:"miscellaneous_fixed_assets,omitempty"`

	Total_other_fixed_assets float64 `xml:"total_other_fixed_assets,omitempty"`

	Total_fixed_assets float64 `xml:"total_fixed_assets,omitempty"`

	Raw_materials float64 `xml:"raw_materials,omitempty"`

	Work_in_progress float64 `xml:"work_in_progress,omitempty"`

	Finished_goods float64 `xml:"finished_goods,omitempty"`

	Other_inventories float64 `xml:"other_inventories,omitempty"`

	Total_inventories float64 `xml:"total_inventories,omitempty"`

	Trade_receivables float64 `xml:"trade_receivables,omitempty"`

	Group_receivables float64 `xml:"group_receivables,omitempty"`

	Receivables_due_after1year float64 `xml:"receivables_due_after1year,omitempty"`

	Miscellaneous_receivables float64 `xml:"miscellaneous_receivables,omitempty"`

	Total_receivables float64 `xml:"total_receivables,omitempty"`

	Cash float64 `xml:"cash,omitempty"`

	Other_current_assets float64 `xml:"other_current_assets,omitempty"`

	Total_current_assets float64 `xml:"total_current_assets,omitempty"`

	Total_assets float64 `xml:"total_assets,omitempty"`

	Trade_payables float64 `xml:"trade_payables,omitempty"`

	Bank_liabilities float64 `xml:"bank_liabilities,omitempty"`

	Other_loans_or_finance float64 `xml:"other_loans_or_finance,omitempty"`

	Group_payables float64 `xml:"group_payables,omitempty"`

	Miscellaneous_liabilities float64 `xml:"miscellaneous_liabilities,omitempty"`

	Total_current_liabilities float64 `xml:"total_current_liabilities,omitempty"`

	Trade_payables_due_after1year float64 `xml:"trade_payables_due_after1year,omitempty"`

	Bank_liabilities_due_after1year float64 `xml:"bank_liabilities_due_after1year,omitempty"`

	Other_loans_or_finance_due_after1year float64 `xml:"other_loans_or_finance_due_after1year,omitempty"`

	Group_payables_due_after1year float64 `xml:"group_payables_due_after1year,omitempty"`

	Miscellaneous_liabilities_due_after1year float64 `xml:"miscellaneous_liabilities_due_after1year,omitempty"`

	Total_long_term_liabilities float64 `xml:"total_long_term_liabilities,omitempty"`

	Total_liabilities float64 `xml:"total_liabilities,omitempty"`

	Called_up_share_capital float64 `xml:"called_up_share_capital,omitempty"`

	Share_premium float64 `xml:"share_premium,omitempty"`

	Revenue_reserves float64 `xml:"revenue_reserves,omitempty"`

	Other_reserves float64 `xml:"other_reserves,omitempty"`

	Total_shareholders_equity float64 `xml:"total_shareholders_equity,omitempty"`
}

type CreditsafeOtherFinancials struct {
	Contingent_liabilities string `xml:"contingent_liabilities,omitempty"`

	Working_capital float64 `xml:"working_capital,omitempty"`

	Net_worth float64 `xml:"net_worth,omitempty"`
}

type CreditsafeFinancialRatios struct {
	Pre_tax_profit_margin float64 `xml:"pre_tax_profit_margin,omitempty"`

	Return_on_capital_employed float64 `xml:"return_on_capital_employed,omitempty"`

	Return_on_total_assets_employed float64 `xml:"return_on_total_assets_employed,omitempty"`

	Return_on_net_assets_employed float64 `xml:"return_on_net_assets_employed,omitempty"`

	Sales_or_net_working_capital float64 `xml:"sales_or_net_working_capital,omitempty"`

	Stock_turnover_ratio float64 `xml:"stock_turnover_ratio,omitempty"`

	Debtor_days float64 `xml:"debtor_days,omitempty"`

	Creditor_days float64 `xml:"creditor_days,omitempty"`

	Current_ratio float64 `xml:"current_ratio,omitempty"`

	Liquidity_ratio_or_acid_test float64 `xml:"liquidity_ratio_or_acid_test,omitempty"`

	Current_debt_ratio float64 `xml:"current_debt_ratio,omitempty"`

	Gearing float64 `xml:"gearing,omitempty"`

	Equity_in_percentage float64 `xml:"equity_in_percentage,omitempty"`

	Total_debt_ratio float64 `xml:"total_debt_ratio,omitempty"`
}

type CreditsafeCompanyUpdate struct {
	Company_id string `xml:"company_id,omitempty"`

	Changes *StringArray `xml:"changes,omitempty"`

	Date_last_update time.Time `xml:"date_last_update,omitempty"`
}

type CreditsafeCompanyUpdateArray struct {
	Item []*CreditsafeCompanyUpdate `xml:"item,omitempty"`
}

type CreditsafeCompanyUpdatePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *CreditsafeCompanyUpdateArray `xml:"results,omitempty"`
}

type CreditsafeAdditionalInformation struct {
	Nl *CreditsafeNlAdditionalInformation `xml:"nl,omitempty"`

	Be *CreditsafeBeAdditionalInformation `xml:"be,omitempty"`
}

type CreditsafeNlAdditionalInformation struct {
	Negative_information *CreditsafeNlNegativeInformation `xml:"negative_information,omitempty"`

	Misc *CreditsafeNlMisc `xml:"misc,omitempty"`

	Industry_quartile_analysis *CreditsafeIndustryQuartilyAnalysis `xml:"industry_quartile_analysis,omitempty"`

	Payment_expectations_summary *CreditsafeNlPaymentExpectationsSummary `xml:"payment_expectations_summary,omitempty"`

	Letters_of_liability_information_403 *CreditsafeNlLettersOfLiabilityInformation403 `xml:"letters_of_liability_information_403,omitempty"`
}

type CreditsafeNlNegativeInformation struct {
	Court_data *CreditsafeNlCourtData `xml:"court_data,omitempty"`
}

type CreditsafeNlCourtData struct {
	Facts *CreditsafeNlCourtDataFacts `xml:"facts,omitempty"`

	Receiver_details *CreditsafeNlCourtDataReceiver `xml:"receiver_details,omitempty"`
}

type CreditsafeNlCourtDataFacts struct {
	Court_action bool `xml:"court_action,omitempty"`

	Date_of_bankruptcy time.Time `xml:"date_of_bankruptcy,omitempty"`

	Details string `xml:"details,omitempty"`
}

type CreditsafeNlCourtDataReceiver struct {
	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`
}

type CreditsafeNlMisc struct {
	Ceased_trading_date time.Time `xml:"ceased_trading_date,omitempty"`

	Exporter bool `xml:"exporter,omitempty"`

	Importer bool `xml:"importer,omitempty"`
}

type CreditsafeIndustryQuartilyAnalysis struct {
	Payment_expectation_days *CreditsafeQuartiles `xml:"payment_expectation_days,omitempty"`

	Day_sales_outstanding *CreditsafeQuartiles `xml:"day_sales_outstanding,omitempty"`
}

type CreditsafeQuartiles struct {
	Lower float64 `xml:"lower,omitempty"`

	Median float64 `xml:"median,omitempty"`

	Upper float64 `xml:"upper,omitempty"`
}

type CreditsafeNlPaymentExpectationsSummary struct {
	Suspension_of_payments_mora bool `xml:"suspension_of_payments_mora,omitempty"`

	Payment_expectation_days float64 `xml:"payment_expectation_days,omitempty"`

	Day_sales_outstanding float64 `xml:"day_sales_outstanding,omitempty"`

	Industry_average_payment_expectation_days float64 `xml:"industry_average_payment_expectation_days,omitempty"`

	Industry_average_day_sales_outstanding float64 `xml:"industry_average_day_sales_outstanding,omitempty"`
}

type CreditsafeNlLettersOfLiabilityInformation403 struct {
	Financial_year uint32 `xml:"financial_year,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Company_number uint32 `xml:"company_number,omitempty"`

	Date_submitted time.Time `xml:"date_submitted,omitempty"`

	Letters_of_liability *CreditsafeNlLettersOfLiabilityArray `xml:"letters_of_liability,omitempty"`

	Letters_of_liability_2nd_parent *CreditsafeNlLettersOfLiabilityArray `xml:"letters_of_liability_2nd_parent,omitempty"`
}

type CreditsafeNlLettersOfLiability struct {
	Start_date time.Time `xml:"start_date,omitempty"`

	Submitted_date time.Time `xml:"submitted_date,omitempty"`

	Removal_submitted_date time.Time `xml:"removal_submitted_date,omitempty"`

	Removal_date time.Time `xml:"removal_date,omitempty"`
}

type CreditsafeNlLettersOfLiabilityArray struct {
	Item []*CreditsafeNlLettersOfLiability `xml:"item,omitempty"`
}

type CreditsafeBeAdditionalInformation struct {
	Negative_information *CreditsafeBeNegativeInformation `xml:"negative_information,omitempty"`

	Misc *CreditsafeBeMisc `xml:"misc,omitempty"`

	Industry_quartile_analysis *CreditsafeIndustryQuartilyAnalysis `xml:"industry_quartile_analysis,omitempty"`

	Payment_expectations_summary *CreditsafeBePaymentExpectationsSummary `xml:"payment_expectations_summary,omitempty"`

	Industry_comparison *CreditsafeBeIndustryComparison `xml:"industry_comparison,omitempty"`

	Registered_contractors *CreditsafeBeRegisteredContractorArray `xml:"registered_contractors,omitempty"`

	Events *CreditsafeBeEventArray `xml:"events,omitempty"`
}

type CreditsafeBeNegativeInformation struct {
	Nsso_details *CreditsafeBeNSSODetailArray `xml:"nsso_details,omitempty"`

	Protested_bills *CreditsafeBeProtestedBillArray `xml:"protested_bills,omitempty"`

	Bankruptcy_events *CreditsafeBeBankruptcyEventArray `xml:"bankruptcy_events,omitempty"`
}

type CreditsafeBeNSSODetail struct {
	Business_number string `xml:"business_number,omitempty"`

	Name_of_defendant string `xml:"name_of_defendant,omitempty"`

	Legal_form_of_defendant string `xml:"legal_form_of_defendant,omitempty"`

	Date_of_summons time.Time `xml:"date_of_summons,omitempty"`

	Labour_court string `xml:"labour_court,omitempty"`
}

type CreditsafeBeNSSODetailArray struct {
	Item []*CreditsafeBeNSSODetail `xml:"item,omitempty"`
}

type CreditsafeBeProtestedBill struct {
	Drawee_name string `xml:"drawee_name,omitempty"`

	Drawee_address string `xml:"drawee_address,omitempty"`

	Bill_amount float64 `xml:"bill_amount,omitempty"`

	Bill_currency string `xml:"bill_currency,omitempty"`

	Maturity_of_bill_in_months float64 `xml:"maturity_of_bill_in_months,omitempty"`

	Payment_date time.Time `xml:"payment_date,omitempty"`

	Name_of_drawer string `xml:"name_of_drawer,omitempty"`

	City_of_drawer string `xml:"city_of_drawer,omitempty"`
}

type CreditsafeBeProtestedBillArray struct {
	Item []*CreditsafeBeProtestedBill `xml:"item,omitempty"`
}

type CreditsafeBeBankruptcyEvent struct {
	Bankruptcy_type string `xml:"bankruptcy_type,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Street_of_registered_office string `xml:"street_of_registered_office,omitempty"`

	City_of_registered_office string `xml:"city_of_registered_office,omitempty"`

	Postcode_of_registered_office string `xml:"postcode_of_registered_office,omitempty"`

	Nace_bel_code string `xml:"nace_bel_code,omitempty"`

	Nace_bel_description string `xml:"nace_bel_description,omitempty"`

	Commercial_court string `xml:"commercial_court,omitempty"`

	Source_date time.Time `xml:"source_date,omitempty"`

	Source string `xml:"source,omitempty"`

	Date_of_decreee_of_bankruptcy time.Time `xml:"date_of_decreee_of_bankruptcy,omitempty"`

	Closing_date_of_summons time.Time `xml:"closing_date_of_summons,omitempty"`

	Name_of_trustee string `xml:"name_of_trustee,omitempty"`

	Street_of_trustee string `xml:"street_of_trustee,omitempty"`

	City_of_trustee string `xml:"city_of_trustee,omitempty"`

	Postcode_of_trustee string `xml:"postcode_of_trustee,omitempty"`
}

type CreditsafeBeBankruptcyEventArray struct {
	Item []*CreditsafeBeBankruptcyEvent `xml:"item,omitempty"`
}

type CreditsafeBeMisc struct {
	Fax_number string `xml:"fax_number,omitempty"`
}

type CreditsafeBePaymentExpectationsSummary struct {
	Payment_expectation_days float64 `xml:"payment_expectation_days,omitempty"`

	Day_sales_outstanding float64 `xml:"day_sales_outstanding,omitempty"`
}

type CreditsafeBeIndustryComparison struct {
	Activity_code uint32 `xml:"activity_code,omitempty"`

	Activity_description string `xml:"activity_description,omitempty"`

	Industry_average_payment_expectation_days float64 `xml:"industry_average_payment_expectation_days,omitempty"`

	Industry_average_day_sales_outstanding float64 `xml:"industry_average_day_sales_outstanding,omitempty"`

	Industry_average_credit_rating float64 `xml:"industry_average_credit_rating,omitempty"`
}

type CreditsafeBeRegisteredContractor struct {
	Registration_number uint32 `xml:"registration_number,omitempty"`

	Contractor_description string `xml:"contractor_description,omitempty"`

	Striking_off_date time.Time `xml:"striking_off_date,omitempty"`
}

type CreditsafeBeRegisteredContractorArray struct {
	Item []*CreditsafeBeRegisteredContractor `xml:"item,omitempty"`
}

type CreditsafeBeEvent struct {
	Serial_number uint32 `xml:"serial_number,omitempty"`

	Description string `xml:"description,omitempty"`
}

type CreditsafeBeEventArray struct {
	Item []*CreditsafeBeEvent `xml:"item,omitempty"`
}

type CreditsafeSearchResultV2 struct {
	Companies *CreditsafeCompanyV2Array `xml:"companies,omitempty"`
}

type CreditsafeCompanyV2Array struct {
	Item []*CreditsafeCompanyV2 `xml:"item,omitempty"`
}

type CreditsafeAddressV2 struct {
	City string `xml:"city,omitempty"`

	Country string `xml:"country,omitempty"`

	Full_address string `xml:"full_address,omitempty"`

	House_number string `xml:"house_number,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	Street string `xml:"street,omitempty"`

	Telephone string `xml:"telephone,omitempty"`
}

type CreditsafeAddressV2Array struct {
	Item []*CreditsafeAddressV2 `xml:"item,omitempty"`
}

type CreditsafeSearchCriteria struct {
	Countries *StringArray `xml:"countries,omitempty"`

	Languages *StringArray `xml:"languages,omitempty"`

	Combinations *CreditsafeSearchCriteriaCombinationArray `xml:"combinations,omitempty"`
}

type CreditsafeSearchCriteriaCombination struct {
	Params *CreditsafeSearchCriteriaParamArray `xml:"params,omitempty"`
}

type CreditsafeSearchCriteriaCombinationArray struct {
	Item []*CreditsafeSearchCriteriaCombination `xml:"item,omitempty"`
}

type CreditsafeSearchCriteriaParam struct {
	Name string `xml:"name,omitempty"`

	Min_length int32 `xml:"min_length,omitempty"`

	Max_length int32 `xml:"max_length,omitempty"`

	Validation_regexp string `xml:"validation_regexp,omitempty"`

	Allowed_values *StringArray `xml:"allowed_values,omitempty"`

	Is_required bool `xml:"is_required,omitempty"`
}

type CreditsafeSearchCriteriaParamArray struct {
	Item []*CreditsafeSearchCriteriaParam `xml:"item,omitempty"`
}

type CreditsafeCompanyReportFullV2 struct {
	Additional_information *CreditsafeCompanyAdditionalInformationV2 `xml:"additional_information,omitempty"`

	Company_identification *CreditsafeCompanyIdentificationV2 `xml:"company_identification,omitempty"`

	Company_summary *CreditsafeCompanySummaryV2 `xml:"company_summary,omitempty"`

	Contact_information *CreditsafeCompanyContactInformationV2 `xml:"contact_information,omitempty"`

	Credit_score *CreditsafeCompanyCreditScoreV2 `xml:"credit_score,omitempty"`

	Directors *CreditsafePersonDirectorsV2 `xml:"directors,omitempty"`

	Financial_statements *CreditsafeCompanyFinancialStatementV2Array `xml:"financial_statements,omitempty"`

	Group_structure *CreditsafeCompanyGroupStructureV2 `xml:"group_structure,omitempty"`

	Local_financial_statements *CreditsafeCompanyFinancialStatementV2Array `xml:"local_financial_statements,omitempty"`

	Negative_information *CreditsafeCompanyNegativeInformationV2 `xml:"negative_information,omitempty"`

	Order_number string `xml:"order_number,omitempty"`

	Other_information *CreditsafeCompanyOtherInformationV2 `xml:"other_information,omitempty"`

	Share_capital_structure *CreditsafeCompanyShareCapitalStructureV2 `xml:"share_capital_structure,omitempty"`

	Payment_data *CreditsafeCompanyPaymentDataV2 `xml:"payment_data,omitempty"`

	Company_id string `xml:"company_id,omitempty"`

	Language string `xml:"language,omitempty"`
}

type CreditsafeCompanyReportFullV2Array struct {
	Item []*CreditsafeCompanyReportFullV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyPaymentDataV2 struct {
	Generic *CreditsafeCompanyKeyValueV2Array `xml:"generic,omitempty"`
}

type CreditsafeCompanyPaymentDataV2Array struct {
	Item []*CreditsafeCompanyPaymentDataV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyShareCapitalStructureV2 struct {
	Issued_share_capital string `xml:"issued_share_capital,omitempty"`

	Nominal_share_capital string `xml:"nominal_share_capital,omitempty"`

	Share_holders *CreditsafeCompanyShareHolderV2Array `xml:"share_holders,omitempty"`
}

type CreditsafeCompanyShareCapitalStructureV2Array struct {
	Item []*CreditsafeCompanyShareCapitalStructureV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyShareHolderV2 struct {
	Address *CreditsafeAddressV2 `xml:"address,omitempty"`

	Share_percent int32 `xml:"share_percent,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CreditsafeCompanyShareHolderV2Array struct {
	Item []*CreditsafeCompanyShareHolderV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyOtherInformationV2 struct {
	Bankers *CreditsafeCompanyBankerV2Array `xml:"bankers,omitempty"`

	Advisors *CreditsafeCompanyAdvisorV2Array `xml:"advisors,omitempty"`

	Employees_information *CreditsafeCompanyEmployeeInformationV2Array `xml:"employees_information,omitempty"`
}

type CreditsafeCompanyOtherInformationV2Array struct {
	Item []*CreditsafeCompanyOtherInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyEmployeeInformationV2 struct {
	Number_of_employees int32 `xml:"number_of_employees,omitempty"`

	Year int32 `xml:"year,omitempty"`
}

type CreditsafeCompanyEmployeeInformationV2Array struct {
	Item []*CreditsafeCompanyEmployeeInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdvisorV2 struct {
	Auditor_name string `xml:"auditor_name,omitempty"`

	Solicitor_name string `xml:"solicitor_name,omitempty"`
}

type CreditsafeCompanyAdvisorV2Array struct {
	Item []*CreditsafeCompanyAdvisorV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyBankerV2 struct {
	Address *CreditsafeAddressV2 `xml:"address,omitempty"`

	Bank_code string `xml:"bank_code,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CreditsafeCompanyBankerV2Array struct {
	Item []*CreditsafeCompanyBankerV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationV2 struct {
	Nl *CreditsafeCompanyNegativeInformationNLV2 `xml:"nl,omitempty"`

	De *CreditsafeCompanyNegativeInformationDEV2 `xml:"de,omitempty"`

	Generic *CreditsafeCompanyKeyValueV2Array `xml:"generic,omitempty"`
}

type CreditsafeCompanyNegativeInformationV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEV2 struct {
	Bankruptcy *CreditsafeCompanyNegativeInformationDEBankruptcyV2 `xml:"bankruptcy,omitempty"`

	Insolvency_information *CreditsafeCompanyNegativeInformationDEInsolvencyInformationV2 `xml:"insolvency_information,omitempty"`

	Director_events *CreditsafeCompanyNegativeInformationDEDirectorEventV2Array `xml:"director_events,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationDEV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEDirectorEventV2 struct {
	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Description string `xml:"description,omitempty"`

	Code string `xml:"code,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEDirectorEventV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationDEDirectorEventV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEInsolvencyInformationV2 struct {
	Insolvency_events *CreditsafeCompanyNegativeInformationDEInsolvencyEventV2Array `xml:"insolvency_events,omitempty"`

	Dateof_insolvency string `xml:"dateof_insolvency,omitempty"`

	Trustee *CreditsafeCompanyNegativeInformationDETrusteeV2 `xml:"trustee,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEInsolvencyInformationV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationDEInsolvencyInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationDETrusteeV2 struct {
	Name string `xml:"name,omitempty"`

	Title string `xml:"title,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Code string `xml:"code,omitempty"`

	City string `xml:"city,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Email_address string `xml:"email_address,omitempty"`
}

type CreditsafeCompanyNegativeInformationDETrusteeV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationDETrusteeV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEInsolvencyEventV2 struct {
	Reference_number_court string `xml:"reference_number_court,omitempty"`

	Created_at time.Time `xml:"created_at,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEInsolvencyEventV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationDEInsolvencyEventV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEBankruptcyV2 struct {
	Has_bankruptcy string `xml:"has_bankruptcy,omitempty"`
}

type CreditsafeCompanyNegativeInformationDEBankruptcyV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationDEBankruptcyV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationNLV2 struct {
	Court_items *CreditsafeCompanyNegativeInformationNLCourtItemV2Array `xml:"court_items,omitempty"`
}

type CreditsafeCompanyNegativeInformationNLV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationNLV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyNegativeInformationNLCourtItemV2 struct {
	Court_action string `xml:"court_action,omitempty"`

	Date_of_bankruptcy time.Time `xml:"date_of_bankruptcy,omitempty"`

	Address string `xml:"address,omitempty"`
}

type CreditsafeCompanyNegativeInformationNLCourtItemV2Array struct {
	Item []*CreditsafeCompanyNegativeInformationNLCourtItemV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyGroupStructureV2 struct {
	Ultimate_parent *CreditsafeCompanyCompanyV2 `xml:"ultimate_parent,omitempty"`

	Immediate_parent *CreditsafeCompanyCompanyV2 `xml:"immediate_parent,omitempty"`

	Subsidiary_companies *CreditsafeCompanyCompanyV2Array `xml:"subsidiary_companies,omitempty"`

	Affiliated_companies *CreditsafeCompanyCompanyV2Array `xml:"affiliated_companies,omitempty"`
}

type CreditsafeCompanyGroupStructureV2Array struct {
	Item []*CreditsafeCompanyGroupStructureV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyCompanyV2 struct {
	Address *CreditsafeAddressV2 `xml:"address,omitempty"`

	Creditsafe_number string `xml:"creditsafe_number,omitempty"`

	Date_last_account time.Time `xml:"date_last_account,omitempty"`

	Date_last_change time.Time `xml:"date_last_change,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Vat_numbers *StringArray `xml:"vat_numbers,omitempty"`

	Country string `xml:"country,omitempty"`

	Registration_number string `xml:"registration_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Id string `xml:"id,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Status string `xml:"status,omitempty"`

	Office_type string `xml:"office_type,omitempty"`
}

type CreditsafeCompanyCompanyV2Array struct {
	Item []*CreditsafeCompanyCompanyV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyFinancialStatementV2 struct {
	Balance_sheet *CreditsafeCompanyBalanceSheetV2 `xml:"balance_sheet,omitempty"`

	Consolidated_accounts bool `xml:"consolidated_accounts,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Number_of_weeks int32 `xml:"number_of_weeks,omitempty"`

	Other_financials *CreditsafeCompanyOtherFinancialsV2 `xml:"other_financials,omitempty"`

	Profit_and_loss *CreditsafeCompanyProfitAndLossV2 `xml:"profit_and_loss,omitempty"`

	Ratios *CreditsafeCompanyRatiosV2 `xml:"ratios,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Year_end_date time.Time `xml:"year_end_date,omitempty"`
}

type CreditsafeCompanyFinancialStatementV2Array struct {
	Item []*CreditsafeCompanyFinancialStatementV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyRatiosV2 struct {
	Creditor_days float32 `xml:"creditor_days,omitempty"`

	Current_debt_ratio float32 `xml:"current_debt_ratio,omitempty"`

	Current_ratio float32 `xml:"current_ratio,omitempty"`

	Debtor_days float32 `xml:"debtor_days,omitempty"`

	Equity_in_percentage float32 `xml:"equity_in_percentage,omitempty"`

	Gearing float32 `xml:"gearing,omitempty"`

	Liquidity_ratio_or_acid_test float32 `xml:"liquidity_ratio_or_acid_test,omitempty"`

	Pre_tax_profit_margin float32 `xml:"pre_tax_profit_margin,omitempty"`

	Return_on_capital_employed float32 `xml:"return_on_capital_employed,omitempty"`

	Return_on_net_assets_employed float32 `xml:"return_on_net_assets_employed,omitempty"`

	Return_on_total_assets_employed float32 `xml:"return_on_total_assets_employed,omitempty"`

	Sales_or_net_working_capital float32 `xml:"sales_or_net_working_capital,omitempty"`

	Stock_turnover_ratio float32 `xml:"stock_turnover_ratio,omitempty"`

	Total_debt_ratio float32 `xml:"total_debt_ratio,omitempty"`
}

type CreditsafeCompanyRatiosV2Array struct {
	Item []*CreditsafeCompanyRatiosV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyProfitAndLossV2 struct {
	Depreciation int32 `xml:"depreciation,omitempty"`

	Dividends int32 `xml:"dividends,omitempty"`

	Financial_expenses int32 `xml:"financial_expenses,omitempty"`

	Financial_income int32 `xml:"financial_income,omitempty"`

	Operating_costs int32 `xml:"operating_costs,omitempty"`

	Operating_profit int32 `xml:"operating_profit,omitempty"`

	Other_appropriations int32 `xml:"other_appropriations,omitempty"`

	Pension_costs int32 `xml:"pension_costs,omitempty"`

	Profit_after_tax int32 `xml:"profit_after_tax,omitempty"`

	Profit_before_tax int32 `xml:"profit_before_tax,omitempty"`

	Retained_profit int32 `xml:"retained_profit,omitempty"`

	Revenue int32 `xml:"revenue,omitempty"`

	Tax int32 `xml:"tax,omitempty"`

	Wages_and_salaries int32 `xml:"wages_and_salaries,omitempty"`
}

type CreditsafeCompanyProfitAndLossV2Array struct {
	Item []*CreditsafeCompanyProfitAndLossV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyOtherFinancialsV2 struct {
	Net_worth int32 `xml:"net_worth,omitempty"`

	Working_capital int32 `xml:"working_capital,omitempty"`
}

type CreditsafeCompanyOtherFinancialsV2Array struct {
	Item []*CreditsafeCompanyOtherFinancialsV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyBalanceSheetV2 struct {
	Bank_liabilities int32 `xml:"bank_liabilities,omitempty"`

	Bank_liabilities_due_after1year int32 `xml:"bank_liabilities_due_after1year,omitempty"`

	Called_up_share_capital int32 `xml:"called_up_share_capital,omitempty"`

	Cash int32 `xml:"cash,omitempty"`

	Consolidated_accounts int32 `xml:"consolidated_accounts,omitempty"`

	Finished_goods int32 `xml:"finished_goods,omitempty"`

	Goodwill int32 `xml:"goodwill,omitempty"`

	Group_payables int32 `xml:"group_payables,omitempty"`

	Group_payables_due_after1year int32 `xml:"group_payables_due_after1year,omitempty"`

	Group_receivables int32 `xml:"group_receivables,omitempty"`

	Investments int32 `xml:"investments,omitempty"`

	Land_and_buildings int32 `xml:"land_and_buildings,omitempty"`

	Loans_to_group int32 `xml:"loans_to_group,omitempty"`

	Miscellaneous_fixed_assets int32 `xml:"miscellaneous_fixed_assets,omitempty"`

	Miscellaneous_liabilities int32 `xml:"miscellaneous_liabilities,omitempty"`

	Miscellaneous_liabilities_due_after1_year int32 `xml:"miscellaneous_liabilities_due_after1_year,omitempty"`

	Miscellaneous_receivables int32 `xml:"miscellaneous_receivables,omitempty"`

	Other_current_assets int32 `xml:"other_current_assets,omitempty"`

	Other_intangible_assets int32 `xml:"other_intangible_assets,omitempty"`

	Other_inventories int32 `xml:"other_inventories,omitempty"`

	Other_loans int32 `xml:"other_loans,omitempty"`

	Other_loans_or_finance int32 `xml:"other_loans_or_finance,omitempty"`

	Other_loans_or_finance_due_after1_year int32 `xml:"other_loans_or_finance_due_after1_year,omitempty"`

	Other_reserves int32 `xml:"other_reserves,omitempty"`

	Other_tangible_assets int32 `xml:"other_tangible_assets,omitempty"`

	Plant_and_machinery int32 `xml:"plant_and_machinery,omitempty"`

	Raw_materials int32 `xml:"raw_materials,omitempty"`

	Receivables_due_after1year int32 `xml:"receivables_due_after1year,omitempty"`

	Revenue_reserves int32 `xml:"revenue_reserves,omitempty"`

	Share_premium int32 `xml:"share_premium,omitempty"`

	Total_assets int32 `xml:"total_assets,omitempty"`

	Total_current_assets int32 `xml:"total_current_assets,omitempty"`

	Total_current_liabilities int32 `xml:"total_current_liabilities,omitempty"`

	Total_fixed_assets int32 `xml:"total_fixed_assets,omitempty"`

	Total_intangible_assets int32 `xml:"total_intangible_assets,omitempty"`

	Total_inventories int32 `xml:"total_inventories,omitempty"`

	Total_liabilities int32 `xml:"total_liabilities,omitempty"`

	Total_long_term_liabilities int32 `xml:"total_long_term_liabilities,omitempty"`

	Total_other_fixed_assets int32 `xml:"total_other_fixed_assets,omitempty"`

	Total_receivables int32 `xml:"total_receivables,omitempty"`

	Total_shareholders_equity int32 `xml:"total_shareholders_equity,omitempty"`

	Total_tangible_assets int32 `xml:"total_tangible_assets,omitempty"`

	Trade_payables int32 `xml:"trade_payables,omitempty"`

	Trade_payables_due_after1year int32 `xml:"trade_payables_due_after1year,omitempty"`

	Trade_receivables int32 `xml:"trade_receivables,omitempty"`

	Work_in_progress int32 `xml:"work_in_progress,omitempty"`
}

type CreditsafeCompanyBalanceSheetV2Array struct {
	Item []*CreditsafeCompanyBalanceSheetV2 `xml:"item,omitempty"`
}

type CreditsafePersonDirectorsV2 struct {
	Current_directors *CreditsafePersonDirectorV2Array `xml:"current_directors,omitempty"`

	Previous_directors *CreditsafePersonPreviousDirectorV2Array `xml:"previous_directors,omitempty"`
}

type CreditsafePersonDirectorsV2Array struct {
	Item []*CreditsafePersonDirectorsV2 `xml:"item,omitempty"`
}

type CreditsafePersonPreviousDirectorV2 struct {
	Resignation_date time.Time `xml:"resignation_date,omitempty"`

	Address *CreditsafeAddressV2 `xml:"address,omitempty"`

	Director_type string `xml:"director_type,omitempty"`

	Id string `xml:"id,omitempty"`

	Positions *CreditsafePersonPositionV2Array `xml:"positions,omitempty"`

	Name string `xml:"name,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Gender string `xml:"gender,omitempty"`
}

type CreditsafePersonPreviousDirectorV2Array struct {
	Item []*CreditsafePersonPreviousDirectorV2 `xml:"item,omitempty"`
}

type CreditsafePersonDirectorV2 struct {
	Address *CreditsafeAddressV2 `xml:"address,omitempty"`

	Director_type string `xml:"director_type,omitempty"`

	Id string `xml:"id,omitempty"`

	Positions *CreditsafePersonPositionV2Array `xml:"positions,omitempty"`

	Name string `xml:"name,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Gender string `xml:"gender,omitempty"`
}

type CreditsafePersonDirectorV2Array struct {
	Item []*CreditsafePersonDirectorV2 `xml:"item,omitempty"`
}

type CreditsafePersonPositionV2 struct {
	Name string `xml:"name,omitempty"`

	Date_appointed time.Time `xml:"date_appointed,omitempty"`
}

type CreditsafePersonPositionV2Array struct {
	Item []*CreditsafePersonPositionV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyCreditScoreV2 struct {
	Current_credit_rating *CreditsafeCompanyCreditRatingV2 `xml:"current_credit_rating,omitempty"`

	Previous_credit_rating *CreditsafeCompanyCreditRatingV2 `xml:"previous_credit_rating,omitempty"`
}

type CreditsafeCompanyCreditScoreV2Array struct {
	Item []*CreditsafeCompanyCreditScoreV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyContactInformationV2 struct {
	Main_address *CreditsafeAddressV2 `xml:"main_address,omitempty"`

	Other_addresses *CreditsafeAddressV2Array `xml:"other_addresses,omitempty"`
}

type CreditsafeCompanyContactInformationV2Array struct {
	Item []*CreditsafeCompanyContactInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanySummaryV2 struct {
	Company_status *CreditsafeCompanyStatusV2 `xml:"company_status,omitempty"`

	Credit_rating *CreditsafeCompanyCreditRatingV2 `xml:"credit_rating,omitempty"`

	Latest_shareholders_equity_figure int32 `xml:"latest_shareholders_equity_figure,omitempty"`

	Latest_turnover_figure int32 `xml:"latest_turnover_figure,omitempty"`

	Main_activity *CreditsafeCompanyActivityV2 `xml:"main_activity,omitempty"`

	Country string `xml:"country,omitempty"`

	Number string `xml:"number,omitempty"`

	Company_registration_number string `xml:"company_registration_number,omitempty"`

	Business_name string `xml:"business_name,omitempty"`
}

type CreditsafeCompanySummaryV2Array struct {
	Item []*CreditsafeCompanySummaryV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyCreditRatingV2 struct {
	Common_description string `xml:"common_description,omitempty"`

	Common_value string `xml:"common_value,omitempty"`

	Credit_limit string `xml:"credit_limit,omitempty"`

	Provider_description string `xml:"provider_description,omitempty"`

	Provider_value *CreditsafeValueV2 `xml:"provider_value,omitempty"`
}

type CreditsafeCompanyCreditRatingV2Array struct {
	Item []*CreditsafeCompanyCreditRatingV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyIdentificationV2 struct {
	Basic_information *CreditsafeCompanyBasicInformationV2 `xml:"basic_information,omitempty"`

	Activity_classifications *CreditsafeCompanyActivityClassificationV2Array `xml:"activity_classifications,omitempty"`
}

type CreditsafeCompanyIdentificationV2Array struct {
	Item []*CreditsafeCompanyIdentificationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyActivityClassificationV2 struct {
	Activities *CreditsafeCompanyActivityV2Array `xml:"activities,omitempty"`

	Classification string `xml:"classification,omitempty"`
}

type CreditsafeCompanyActivityClassificationV2Array struct {
	Item []*CreditsafeCompanyActivityClassificationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyActivityV2 struct {
	Classification string `xml:"classification,omitempty"`

	Activity_code string `xml:"activity_code,omitempty"`

	Activity_description string `xml:"activity_description,omitempty"`
}

type CreditsafeCompanyActivityV2Array struct {
	Item []*CreditsafeCompanyActivityV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyBasicInformationV2 struct {
	Company_status *CreditsafeCompanyStatusV2 `xml:"company_status,omitempty"`

	Contact_address *CreditsafeAddressV2 `xml:"contact_address,omitempty"`

	Legal_form *CreditsafeDescriptionV2 `xml:"legal_form,omitempty"`

	Principal_activity *CreditsafeDescriptionV2 `xml:"principal_activity,omitempty"`

	Business_name string `xml:"business_name,omitempty"`

	Company_registration_number string `xml:"company_registration_number,omitempty"`

	Country string `xml:"country,omitempty"`

	Dateof_company_registration time.Time `xml:"dateof_company_registration,omitempty"`

	Dateof_starting_operations time.Time `xml:"dateof_starting_operations,omitempty"`

	Registered_company_name string `xml:"registered_company_name,omitempty"`

	Vat_registration_number string `xml:"vat_registration_number,omitempty"`
}

type CreditsafeCompanyBasicInformationV2Array struct {
	Item []*CreditsafeCompanyBasicInformationV2 `xml:"item,omitempty"`
}

type CreditsafeDescriptionV2 struct {
	Code string `xml:"code,omitempty"`

	Description string `xml:"description,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CreditsafeDescriptionV2Array struct {
	Item []*CreditsafeDescriptionV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyStatusV2 struct {
	Description string `xml:"description,omitempty"`

	Status string `xml:"status,omitempty"`
}

type CreditsafeCompanyStatusV2Array struct {
	Item []*CreditsafeCompanyStatusV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationV2 struct {
	De *CreditsafeCompanyAdditionalInformationDEAdditionalInformationV2 `xml:"de,omitempty"`

	Nl *CreditsafeCompanyAdditionalInformationNLAdditionalInformationV2 `xml:"nl,omitempty"`

	Generic *CreditsafeCompanyKeyValueV2Array `xml:"generic,omitempty"`
}

type CreditsafeCompanyAdditionalInformationV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyKeyValueV2 struct {
	Key string `xml:"key,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CreditsafeCompanyKeyValueV2Array struct {
	Item []*CreditsafeCompanyKeyValueV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLAdditionalInformationV2 struct {
	Payment_expectations_summary *CreditsafeCompanyAdditionalInformationNLPaymentExpectationsSummaryV2 `xml:"payment_expectations_summary,omitempty"`

	Miscellaneous *CreditsafeCompanyAdditionalInformationNLMiscellaneousV2 `xml:"miscellaneous,omitempty"`

	Industy_comparison *CreditsafeCompanyAdditionalInformationNLIndustryComparisonV2 `xml:"industy_comparison,omitempty"`

	Industry_quartile_analysis *CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisV2 `xml:"industry_quartile_analysis,omitempty"`

	Financial_items *CreditsafeCompanyAdditionalInformationNLFinancialItemV2Array `xml:"financial_items,omitempty"`

	Kvk_filings *CreditsafeCompanyAdditionalInformationNLKvkFilingV2Array `xml:"kvk_filings,omitempty"`

	Letters_of_liablility_information403 *CreditsafeCompanyAdditionalInformationNLLetterOfLiabilityInformation403V2Array `xml:"letters_of_liablility_information403,omitempty"`

	Comments *CreditsafeCompanyAdditionalInformationCommentV2Array `xml:"comments,omitempty"`

	Historical_events *CreditsafeCompanyAdditionalInformationHistoricalEventV2Array `xml:"historical_events,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLAdditionalInformationV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLAdditionalInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLLetterOfLiabilityInformation403V2 struct {
	Financial_year int32 `xml:"financial_year,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Company_number string `xml:"company_number,omitempty"`

	Date_submitted time.Time `xml:"date_submitted,omitempty"`

	Letter_of_liability1_start_date time.Time `xml:"letter_of_liability1_start_date,omitempty"`

	Letter_of_liability1_submitted_date time.Time `xml:"letter_of_liability1_submitted_date,omitempty"`

	Letter_of_liability1_removal_submitted time.Time `xml:"letter_of_liability1_removal_submitted,omitempty"`

	Letter_of_liability1_removal_date time.Time `xml:"letter_of_liability1_removal_date,omitempty"`

	Letter_of_liability2_start_date time.Time `xml:"letter_of_liability2_start_date,omitempty"`

	Letter_of_liability2_submitted_date time.Time `xml:"letter_of_liability2_submitted_date,omitempty"`

	Letter_of_liability2_removal_submitted time.Time `xml:"letter_of_liability2_removal_submitted,omitempty"`

	Letter_of_liability2_removal_date time.Time `xml:"letter_of_liability2_removal_date,omitempty"`

	Letter_of_liability1_start_date_parent2 time.Time `xml:"letter_of_liability1_start_date_parent2,omitempty"`

	Letter_of_liability1_submitted_date_parent2 time.Time `xml:"letter_of_liability1_submitted_date_parent2,omitempty"`

	Letter_of_liability1_removal_submitted_parent2 time.Time `xml:"letter_of_liability1_removal_submitted_parent2,omitempty"`

	Letter_of_liability1_removal_date_parent2 time.Time `xml:"letter_of_liability1_removal_date_parent2,omitempty"`

	Letter_of_liability2_start_date_parent2 time.Time `xml:"letter_of_liability2_start_date_parent2,omitempty"`

	Letter_of_liability2_submitted_date_parent2 time.Time `xml:"letter_of_liability2_submitted_date_parent2,omitempty"`

	Letter_of_liability2_removal_submitted_parent2 time.Time `xml:"letter_of_liability2_removal_submitted_parent2,omitempty"`

	Letter_of_liability2_removal_date_parent2 time.Time `xml:"letter_of_liability2_removal_date_parent2,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLLetterOfLiabilityInformation403V2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLLetterOfLiabilityInformation403V2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLKvkFilingV2 struct {
	Date time.Time `xml:"date,omitempty"`

	Event string `xml:"event,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLKvkFilingV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLKvkFilingV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLFinancialItemV2 struct {
	Date_year_end time.Time `xml:"date_year_end,omitempty"`

	Judgement string `xml:"judgement,omitempty"`

	Consolidated_subsidiaries string `xml:"consolidated_subsidiaries,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLFinancialItemV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLFinancialItemV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisV2 struct {
	Payment_expectation_days *CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisDayV2 `xml:"payment_expectation_days,omitempty"`

	Day_sales_outstanding *CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisDayV2 `xml:"day_sales_outstanding,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisDayV2 struct {
	Lower float32 `xml:"lower,omitempty"`

	Median float32 `xml:"median,omitempty"`

	Upper float32 `xml:"upper,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisDayV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLIndustryQuartileAnalysisDayV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLIndustryComparisonV2 struct {
	Industry_average_credit_rating float32 `xml:"industry_average_credit_rating,omitempty"`

	Industry_average_credit_limit float32 `xml:"industry_average_credit_limit,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLIndustryComparisonV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLIndustryComparisonV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLMiscellaneousV2 struct {
	Date_of_legal_form time.Time `xml:"date_of_legal_form,omitempty"`

	Date_of_cessation_trading time.Time `xml:"date_of_cessation_trading,omitempty"`

	Exporter string `xml:"exporter,omitempty"`

	Importer string `xml:"importer,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Branch_number string `xml:"branch_number,omitempty"`

	Negative_rating int32 `xml:"negative_rating,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLMiscellaneousV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLMiscellaneousV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLPaymentExpectationsSummaryV2 struct {
	Suspension_of_payments_mora string `xml:"suspension_of_payments_mora,omitempty"`

	Payment_expectation_days int32 `xml:"payment_expectation_days,omitempty"`

	Day_sales_outstanding int32 `xml:"day_sales_outstanding,omitempty"`

	Industry_average_payment_expectation_days int32 `xml:"industry_average_payment_expectation_days,omitempty"`

	Industry_average_day_sales_outstanding int32 `xml:"industry_average_day_sales_outstanding,omitempty"`
}

type CreditsafeCompanyAdditionalInformationNLPaymentExpectationsSummaryV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationNLPaymentExpectationsSummaryV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEAdditionalInformationV2 struct {
	Miscellaneous *CreditsafeCompanyAdditionalInformationDEMiscellaneousV2 `xml:"miscellaneous,omitempty"`

	Image_accounts *CreditsafeCompanyAdditionalInformationDEImageAccountV2 `xml:"image_accounts,omitempty"`

	Beneficiary_owners *CreditsafeCompanyAdditionalInformationDEOwnerV2 `xml:"beneficiary_owners,omitempty"`

	Turnover_ranges *CreditsafeCompanyAdditionalInformationDETurnoverRangeV2 `xml:"turnover_ranges,omitempty"`

	Comments *CreditsafeCompanyAdditionalInformationCommentV2Array `xml:"comments,omitempty"`

	Historical_events *CreditsafeCompanyAdditionalInformationHistoricalEventV2Array `xml:"historical_events,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEAdditionalInformationV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationDEAdditionalInformationV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationHistoricalEventV2 struct {
	Date time.Time `xml:"date,omitempty"`

	Description string `xml:"description,omitempty"`

	Previous_value string `xml:"previous_value,omitempty"`

	Current_value string `xml:"current_value,omitempty"`
}

type CreditsafeCompanyAdditionalInformationHistoricalEventV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationHistoricalEventV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationCommentV2 struct {
	Text string `xml:"text,omitempty"`

	Sentiment string `xml:"sentiment,omitempty"`
}

type CreditsafeCompanyAdditionalInformationCommentV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationCommentV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDETurnoverRangeV2 struct {
	Year int32 `xml:"year,omitempty"`

	Range_ string `xml:"range,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDETurnoverRangeV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationDETurnoverRangeV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEOwnerV2 struct {
	Name string `xml:"name,omitempty"`

	City string `xml:"city,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Share_percent string `xml:"share_percent,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEOwnerV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationDEOwnerV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEImageAccountV2 struct {
	Serial_number string `xml:"serial_number,omitempty"`

	Start_date time.Time `xml:"start_date,omitempty"`

	End_date time.Time `xml:"end_date,omitempty"`

	Published_date time.Time `xml:"published_date,omitempty"`

	Document_type string `xml:"document_type,omitempty"`

	Financials_type string `xml:"financials_type,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEImageAccountV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationDEImageAccountV2 `xml:"item,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEMiscellaneousV2 struct {
	Turnover_range string `xml:"turnover_range,omitempty"`

	Commercial_register_city string `xml:"commercial_register_city,omitempty"`

	Commercial_register_zip string `xml:"commercial_register_zip,omitempty"`

	Financials_quality string `xml:"financials_quality,omitempty"`

	Current_rating *CreditsafeValueV2 `xml:"current_rating,omitempty"`

	Previous_rating *CreditsafeValueV2 `xml:"previous_rating,omitempty"`

	Premise_type string `xml:"premise_type,omitempty"`

	Small_office string `xml:"small_office,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Company_type string `xml:"company_type,omitempty"`

	Negative_rating string `xml:"negative_rating,omitempty"`

	Complementary_company string `xml:"complementary_company,omitempty"`

	Business_purpose string `xml:"business_purpose,omitempty"`
}

type CreditsafeCompanyAdditionalInformationDEMiscellaneousV2Array struct {
	Item []*CreditsafeCompanyAdditionalInformationDEMiscellaneousV2 `xml:"item,omitempty"`
}

type CreditsafeValueV2 struct {
	Max_value string `xml:"max_value,omitempty"`

	Min_value string `xml:"min_value,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CreditsafeValueV2Array struct {
	Item []*CreditsafeValueV2 `xml:"item,omitempty"`
}

type DNBBusinessReference struct {
	Dnb_key string `xml:"dnb_key,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Address string `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Country string `xml:"country,omitempty"`

	Main_branch_indication bool `xml:"main_branch_indication,omitempty"`

	Confidence_code int32 `xml:"confidence_code,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DNBBusinessReferenceArray struct {
	Item []*DNBBusinessReference `xml:"item,omitempty"`
}

type DNBBusinessReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DNBBusinessReferenceArray `xml:"results,omitempty"`
}

type DNBBusinessReferenceV2 struct {
	Dnb_key string `xml:"dnb_key,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Address string `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Country int32 `xml:"country,omitempty"`

	Country_iso2 string `xml:"country_iso2,omitempty"`

	Main_branch_indication bool `xml:"main_branch_indication,omitempty"`

	Confidence_code int32 `xml:"confidence_code,omitempty"`
}

type DNBBusinessReferenceV2Array struct {
	Item []*DNBBusinessReferenceV2 `xml:"item,omitempty"`
}

type DNBBusinessReferenceV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DNBBusinessReferenceV2Array `xml:"results,omitempty"`
}

type DNBQuickCheck struct {
	Duns string `xml:"duns,omitempty"`

	Dnb_key string `xml:"dnb_key,omitempty"`

	Enquiry_duns string `xml:"enquiry_duns,omitempty"`

	Business_number string `xml:"business_number,omitempty"`

	Business_number_type string `xml:"business_number_type,omitempty"`

	Business_number_description string `xml:"business_number_description,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Address *StringArray `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Region string `xml:"region,omitempty"`

	Subregion string `xml:"subregion,omitempty"`

	Country_calling_code string `xml:"country_calling_code,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Country string `xml:"country,omitempty"`

	Ceo_name string `xml:"ceo_name,omitempty"`

	Ceo_title string `xml:"ceo_title,omitempty"`

	Principal_name string `xml:"principal_name,omitempty"`

	Principal_title string `xml:"principal_title,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Statement_currency string `xml:"statement_currency,omitempty"`

	Debtor_inpossession bool `xml:"debtor_inpossession,omitempty"`

	Criminal_proceedings bool `xml:"criminal_proceedings,omitempty"`

	Suit_judgement bool `xml:"suit_judgement,omitempty"`

	Financial_embarrassment bool `xml:"financial_embarrassment,omitempty"`

	Financial_legal_events bool `xml:"financial_legal_events,omitempty"`

	Operational_events bool `xml:"operational_events,omitempty"`

	Other_events bool `xml:"other_events,omitempty"`

	Disaster bool `xml:"disaster,omitempty"`

	Secured_filings bool `xml:"secured_filings,omitempty"`

	Out_of_business bool `xml:"out_of_business,omitempty"`

	History_indicator string `xml:"history_indicator,omitempty"`

	History_description string `xml:"history_description,omitempty"`

	Income_statement_date string `xml:"income_statement_date,omitempty"`

	Income_statement_date_start string `xml:"income_statement_date_start,omitempty"`

	Income_statement_date_end string `xml:"income_statement_date_end,omitempty"`

	Financial_statement_date string `xml:"financial_statement_date,omitempty"`

	Incorporation_year int32 `xml:"incorporation_year,omitempty"`

	Start_year int32 `xml:"start_year,omitempty"`

	Current_control_year int32 `xml:"current_control_year,omitempty"`

	Local_activity_code string `xml:"local_activity_code,omitempty"`

	Local_activity_code_type string `xml:"local_activity_code_type,omitempty"`

	Local_activity_description string `xml:"local_activity_description,omitempty"`

	Sic_activity_code string `xml:"sic_activity_code,omitempty"`

	Sic_activity_description string `xml:"sic_activity_description,omitempty"`

	Sic_version string `xml:"sic_version,omitempty"`

	Legal_form string `xml:"legal_form,omitempty"`

	Legal_form_description string `xml:"legal_form_description,omitempty"`

	Paydex_norm int32 `xml:"paydex_norm,omitempty"`

	Paydex_score int32 `xml:"paydex_score,omitempty"`

	Dnb_rating string `xml:"dnb_rating,omitempty"`

	Maximum_credit int32 `xml:"maximum_credit,omitempty"`

	Maximum_credit_currency string `xml:"maximum_credit_currency,omitempty"`

	Intangible_assets int32 `xml:"intangible_assets,omitempty"`

	Net_worth int32 `xml:"net_worth,omitempty"`

	Tangible_net_worth *DNBMonetaryAmount `xml:"tangible_net_worth,omitempty"`

	Sales int32 `xml:"sales,omitempty"`

	Annual_sales *DNBMonetaryAmount `xml:"annual_sales,omitempty"`

	Employees_primary *DNBEmployeeCount `xml:"employees_primary,omitempty"`

	Employees_total *DNBEmployeeCount `xml:"employees_total,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DNBBusinessVerification struct {
	Duns string `xml:"duns,omitempty"`

	Dnb_key string `xml:"dnb_key,omitempty"`

	Enquiry_duns string `xml:"enquiry_duns,omitempty"`

	Business_number string `xml:"business_number,omitempty"`

	Business_number_type string `xml:"business_number_type,omitempty"`

	Business_number_description string `xml:"business_number_description,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Address *StringArray `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Region string `xml:"region,omitempty"`

	Subregion string `xml:"subregion,omitempty"`

	Country_calling_code string `xml:"country_calling_code,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Country string `xml:"country,omitempty"`

	Main_branch_indication bool `xml:"main_branch_indication,omitempty"`

	Out_of_business bool `xml:"out_of_business,omitempty"`

	Incorporation_year int32 `xml:"incorporation_year,omitempty"`

	Start_year int32 `xml:"start_year,omitempty"`

	Sic_activity_code string `xml:"sic_activity_code,omitempty"`

	Sic_activity_description string `xml:"sic_activity_description,omitempty"`

	Sic_version string `xml:"sic_version,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DNBEnterpriseManagement struct {
	Quickcheck *DNBQuickCheck `xml:"quickcheck,omitempty"`

	Main_branch_indication bool `xml:"main_branch_indication,omitempty"`

	Indate_indicator bool `xml:"indate_indicator,omitempty"`

	Registration_type string `xml:"registration_type,omitempty"`

	Import_indicator bool `xml:"import_indicator,omitempty"`

	Export_indicator bool `xml:"export_indicator,omitempty"`

	Location_ownership string `xml:"location_ownership,omitempty"`

	Claims_indicator bool `xml:"claims_indicator,omitempty"`

	Paydex_score_earlier int32 `xml:"paydex_score_earlier,omitempty"`

	Average_high_credit int32 `xml:"average_high_credit,omitempty"`

	High_credit int32 `xml:"high_credit,omitempty"`

	Total_payments int32 `xml:"total_payments,omitempty"`

	Headquarters *DNBRelatedBusiness `xml:"headquarters,omitempty"`

	Parent *DNBRelatedBusiness `xml:"parent,omitempty"`

	Domestic_ultimate *DNBRelatedBusiness `xml:"domestic_ultimate,omitempty"`

	Global_ultimate *DNBRelatedBusiness `xml:"global_ultimate,omitempty"`

	Is_subsidiary_location bool `xml:"is_subsidiary_location,omitempty"`

	Location_status string `xml:"location_status,omitempty"`

	Accounts_audited bool `xml:"accounts_audited,omitempty"`

	Statement_consolidated bool `xml:"statement_consolidated,omitempty"`

	Figures_estimated bool `xml:"figures_estimated,omitempty"`

	Figures_forecast bool `xml:"figures_forecast,omitempty"`

	Year_period bool `xml:"year_period,omitempty"`

	Figures_final bool `xml:"figures_final,omitempty"`

	Opening_statement bool `xml:"opening_statement,omitempty"`

	Figures_proforma bool `xml:"figures_proforma,omitempty"`

	Trial_balance bool `xml:"trial_balance,omitempty"`

	Figures_signed bool `xml:"figures_signed,omitempty"`

	Figures_restated bool `xml:"figures_restated,omitempty"`

	Figures_unbalanced bool `xml:"figures_unbalanced,omitempty"`

	Figures_qualified bool `xml:"figures_qualified,omitempty"`

	Accounts_receivable int32 `xml:"accounts_receivable,omitempty"`

	Liquid_assets int32 `xml:"liquid_assets,omitempty"`

	Inventory int32 `xml:"inventory,omitempty"`

	Total_current_assets int32 `xml:"total_current_assets,omitempty"`

	Total_assets int32 `xml:"total_assets,omitempty"`

	Accounts_payable int32 `xml:"accounts_payable,omitempty"`

	Total_current_liabilities int32 `xml:"total_current_liabilities,omitempty"`

	Total_liabilities int32 `xml:"total_liabilities,omitempty"`

	Net_income int32 `xml:"net_income,omitempty"`

	Quick_ratio float64 `xml:"quick_ratio,omitempty"`

	Current_ratio float64 `xml:"current_ratio,omitempty"`

	Previous_net_worth int32 `xml:"previous_net_worth,omitempty"`

	Previous_sales int32 `xml:"previous_sales,omitempty"`

	Previous_statement_date string `xml:"previous_statement_date,omitempty"`

	Line_of_business string `xml:"line_of_business,omitempty"`

	Failure_score *DNBScoreGroupArray `xml:"failure_score,omitempty"`

	Credit_score *DNBScoreGroupArray `xml:"credit_score,omitempty"`

	Global_failure_score *DNBScoreGroupArray `xml:"global_failure_score,omitempty"`
}

type DNBScoreGroup struct {
	Score int32 `xml:"score,omitempty"`

	Score_commentaries *DNBCommentaryArray `xml:"score_commentaries,omitempty"`

	Score_class bool `xml:"score_class,omitempty"`

	National_percentile int32 `xml:"national_percentile,omitempty"`

	Score_override_code string `xml:"score_override_code,omitempty"`

	Incidence_of_default int32 `xml:"incidence_of_default,omitempty"`

	Score_date string `xml:"score_date,omitempty"`

	Industry_norm string `xml:"industry_norm,omitempty"`

	Industry_incidence_of_default string `xml:"industry_incidence_of_default,omitempty"`

	Industry_percentile int32 `xml:"industry_percentile,omitempty"`

	Unnormalized_score string `xml:"unnormalized_score,omitempty"`

	Algorithm_id string `xml:"algorithm_id,omitempty"`
}

type DNBScoreGroupArray struct {
	Item []*DNBScoreGroup `xml:"item,omitempty"`
}

type DNBCommentary struct {
	Code string `xml:"code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DNBCommentaryArray struct {
	Item []*DNBCommentary `xml:"item,omitempty"`
}

type DNBRelatedBusiness struct {
	Duns string `xml:"duns,omitempty"`

	Dnb_key string `xml:"dnb_key,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Country string `xml:"country,omitempty"`
}

type DNBEmployeeCount struct {
	Employees int32 `xml:"employees,omitempty"`

	Employees_text string `xml:"employees_text,omitempty"`

	Estimated bool `xml:"estimated,omitempty"`

	Subsidiaries_included bool `xml:"subsidiaries_included,omitempty"`

	Qualification_indicator string `xml:"qualification_indicator,omitempty"`
}

type DNBMonetaryAmount struct {
	Value int32 `xml:"value,omitempty"`

	Value_text string `xml:"value_text,omitempty"`

	Estimated int32 `xml:"estimated,omitempty"`

	Consolidated bool `xml:"consolidated,omitempty"`

	Currency string `xml:"currency,omitempty"`
}

type DNBMarketing struct {
	Duns string `xml:"duns,omitempty"`

	Business_number string `xml:"business_number,omitempty"`

	Business_number_type int32 `xml:"business_number_type,omitempty"`

	Primary_name string `xml:"primary_name,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Former_duns string `xml:"former_duns,omitempty"`

	Former_primary_name string `xml:"former_primary_name,omitempty"`

	Establishment_address *DNBAddress `xml:"establishment_address,omitempty"`

	Country_calling_code string `xml:"country_calling_code,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Marketability_indication bool `xml:"marketability_indication,omitempty"`
}

type DNBMarketingPlusResult struct {
	Marketing *DNBMarketing `xml:"marketing,omitempty"`

	Marketing_plus *DNBMarketingPlus `xml:"marketing_plus,omitempty"`
}

type DNBMarketingPlusLinkageResult struct {
	Marketing *DNBMarketing `xml:"marketing,omitempty"`

	Marketing_plus *DNBMarketingPlus `xml:"marketing_plus,omitempty"`

	Marketing_plus_linkage *DNBMarketingPlusLinkage `xml:"marketing_plus_linkage,omitempty"`
}

type DNBMarketingPlus struct {
	Correspondence_address *DNBAddress `xml:"correspondence_address,omitempty"`

	Correspondence_deliverability_indication bool `xml:"correspondence_deliverability_indication,omitempty"`

	Establishment_deliverability_indication bool `xml:"establishment_deliverability_indication,omitempty"`

	Location_status string `xml:"location_status,omitempty"`

	Is_subsidiary_location bool `xml:"is_subsidiary_location,omitempty"`

	Legal_form int32 `xml:"legal_form,omitempty"`

	Line_of_business string `xml:"line_of_business,omitempty"`

	Sic_activity_codes *DNBSICCodeArray `xml:"sic_activity_codes,omitempty"`

	Sic8_activity_codes *DNBSICCodeArray `xml:"sic8_activity_codes,omitempty"`

	Local_activity_code string `xml:"local_activity_code,omitempty"`

	Local_activity_code_type int32 `xml:"local_activity_code_type,omitempty"`

	Start_year int32 `xml:"start_year,omitempty"`

	Ceo_name string `xml:"ceo_name,omitempty"`

	Ceo_title string `xml:"ceo_title,omitempty"`

	Executive_names *StringArray `xml:"executive_names,omitempty"`

	Statement_date string `xml:"statement_date,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Annual_sales int64 `xml:"annual_sales,omitempty"`

	Annual_sales_us int64 `xml:"annual_sales_us,omitempty"`

	Annual_sales_indication int32 `xml:"annual_sales_indication,omitempty"`

	Net_worth int64 `xml:"net_worth,omitempty"`

	Net_worth_us int64 `xml:"net_worth_us,omitempty"`

	Net_income int64 `xml:"net_income,omitempty"`

	Net_income_us int64 `xml:"net_income_us,omitempty"`

	Import_indication bool `xml:"import_indication,omitempty"`

	Export_indication bool `xml:"export_indication,omitempty"`

	Agent_indication bool `xml:"agent_indication,omitempty"`

	Employees int32 `xml:"employees,omitempty"`

	Employees_indication int32 `xml:"employees_indication,omitempty"`

	Employees_total int32 `xml:"employees_total,omitempty"`

	Employees_total_indication int32 `xml:"employees_total_indication,omitempty"`

	Employees_total_includes_principals bool `xml:"employees_total_includes_principals,omitempty"`

	Report_date string `xml:"report_date,omitempty"`

	Out_of_business bool `xml:"out_of_business,omitempty"`
}

type DNBMarketingPlusLinkage struct {
	Domestic_ultimate_duns string `xml:"domestic_ultimate_duns,omitempty"`

	Domestic_ultimate_name string `xml:"domestic_ultimate_name,omitempty"`

	Domestic_ultimate_address *DNBAddress `xml:"domestic_ultimate_address,omitempty"`

	Global_ultimate_indication bool `xml:"global_ultimate_indication,omitempty"`

	Global_ultimate_duns string `xml:"global_ultimate_duns,omitempty"`

	Global_ultimate_name string `xml:"global_ultimate_name,omitempty"`

	Global_ultimate_address *DNBAddress `xml:"global_ultimate_address,omitempty"`

	Headquarters_duns string `xml:"headquarters_duns,omitempty"`

	Headquarters_name string `xml:"headquarters_name,omitempty"`

	Headquarters_address *DNBAddress `xml:"headquarters_address,omitempty"`

	Family_member_count int32 `xml:"family_member_count,omitempty"`

	Linkage_report_date string `xml:"linkage_report_date,omitempty"`
}

type DNBAddress struct {
	Address string `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Region string `xml:"region,omitempty"`

	Region_short string `xml:"region_short,omitempty"`

	Region_code int32 `xml:"region_code,omitempty"`

	Country string `xml:"country,omitempty"`

	Country_dnb int32 `xml:"country_dnb,omitempty"`

	Country_iso string `xml:"country_iso,omitempty"`
}

type DNBSICCode struct {
	Code string `xml:"code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DNBSICCodeArray struct {
	Item []*DNBSICCode `xml:"item,omitempty"`
}

type DriveInfo struct {
	Fastest int32 `xml:"fastest,omitempty"`

	Shortest int32 `xml:"shortest,omitempty"`
}

type DutchAddressPostcodeRange struct {
	House_number_from int32 `xml:"house_number_from,omitempty"`

	House_number_to int32 `xml:"house_number_to,omitempty"`

	Neighborhood_code string `xml:"neighborhood_code,omitempty"`

	Letter_combination string `xml:"letter_combination,omitempty"`

	Range_indication string `xml:"range_indication,omitempty"`

	Street_name string `xml:"street_name,omitempty"`

	Street_name_nen string `xml:"street_name_nen,omitempty"`

	City string `xml:"city,omitempty"`

	Municipality string `xml:"municipality,omitempty"`

	Municipality_code int32 `xml:"municipality_code,omitempty"`

	Cebuco_code int32 `xml:"cebuco_code,omitempty"`

	Province string `xml:"province,omitempty"`

	Province_code string `xml:"province_code,omitempty"`
}

type DutchBusinessDossier struct {
	Update_info *DutchBusinessUpdateReference `xml:"update_info,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment_number string `xml:"main_establishment_number,omitempty"`

	Indication_main_establishment bool `xml:"indication_main_establishment,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Chamber_number string `xml:"chamber_number,omitempty"`

	Legal_form_code string `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Indication_organisation_code string `xml:"indication_organisation_code,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Trade_name_45 string `xml:"trade_name_45,omitempty"`

	Trade_name_full string `xml:"trade_name_full,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Correspondence_postcode string `xml:"correspondence_postcode,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Correspondence_house_number int32 `xml:"correspondence_house_number,omitempty"`

	Correspondence_house_number_addition string `xml:"correspondence_house_number_addition,omitempty"`

	Correspondence_country string `xml:"correspondence_country,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Mobile_number string `xml:"mobile_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Contact_title1 string `xml:"contact_title1,omitempty"`

	Contact_title2 string `xml:"contact_title2,omitempty"`

	Contact_initials string `xml:"contact_initials,omitempty"`

	Contact_prefix string `xml:"contact_prefix,omitempty"`

	Contact_surname string `xml:"contact_surname,omitempty"`

	Contact_gender string `xml:"contact_gender,omitempty"`

	Primary_sbi_code string `xml:"primary_sbi_code,omitempty"`

	Secondary_sbi_code1 string `xml:"secondary_sbi_code1,omitempty"`

	Secondary_sbi_code2 string `xml:"secondary_sbi_code2,omitempty"`

	Primary_sbi_code_text string `xml:"primary_sbi_code_text,omitempty"`

	Secondary_sbi_code1_text string `xml:"secondary_sbi_code1_text,omitempty"`

	Secondary_sbi_code2_text string `xml:"secondary_sbi_code2_text,omitempty"`

	Personnel int32 `xml:"personnel,omitempty"`

	Class_personnel int32 `xml:"class_personnel,omitempty"`

	Personnel_fulltime int32 `xml:"personnel_fulltime,omitempty"`

	Class_personnel_fulltime int32 `xml:"class_personnel_fulltime,omitempty"`

	Personnel_reference_date *DutchBusinessDate `xml:"personnel_reference_date,omitempty"`

	Indication_import bool `xml:"indication_import,omitempty"`

	Indication_export bool `xml:"indication_export,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`

	Indication_non_mailing bool `xml:"indication_non_mailing,omitempty"`

	Indication_bankruptcy bool `xml:"indication_bankruptcy,omitempty"`

	Indication_dip bool `xml:"indication_dip,omitempty"`

	Authorized_share_capital int64 `xml:"authorized_share_capital,omitempty"`

	Authorized_share_capital_currency string `xml:"authorized_share_capital_currency,omitempty"`

	Paid_up_share_capital int64 `xml:"paid_up_share_capital,omitempty"`

	Paid_up_share_capital_currency string `xml:"paid_up_share_capital_currency,omitempty"`

	Issued_share_capital int64 `xml:"issued_share_capital,omitempty"`

	Issued_share_capital_currency string `xml:"issued_share_capital_currency,omitempty"`

	Founding_date *DutchBusinessDate `xml:"founding_date,omitempty"`

	Continuation_date *DutchBusinessDate `xml:"continuation_date,omitempty"`

	Establishment_date *DutchBusinessDate `xml:"establishment_date,omitempty"`
}

type DutchBusinessGetConcernRelationsOverviewResult struct {
	References *DutchBusinessGetConcernRelationsOverviewReferenceArray `xml:"references,omitempty"`
}

type DutchBusinessGetConcernRelationsOverviewReference struct {
	Name string `xml:"name,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Levels int32 `xml:"levels,omitempty"`

	Level_of_matched_registration int32 `xml:"level_of_matched_registration,omitempty"`

	Registrations int32 `xml:"registrations,omitempty"`
}

type DutchBusinessGetConcernRelationsOverviewReferenceArray struct {
	Item []*DutchBusinessGetConcernRelationsOverviewReference `xml:"item,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsResult struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Graph *DutchBusinessGetConcernRelationsDetailsGraph `xml:"graph,omitempty"`

	Pdf []byte `xml:"pdf,omitempty"`

	Source []byte `xml:"source,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsGraph struct {
	Nodes *DutchBusinessGetConcernRelationsDetailsGraphNodeArray `xml:"nodes,omitempty"`

	Links *DutchBusinessGetConcernRelationsDetailsGraphLinkArray `xml:"links,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsGraphNode struct {
	Id string `xml:"id,omitempty"`

	Name string `xml:"name,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsGraphNodeArray struct {
	Item []*DutchBusinessGetConcernRelationsDetailsGraphNode `xml:"item,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsGraphLink struct {
	Source string `xml:"source,omitempty"`

	Target string `xml:"target,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsGraphLinkArray struct {
	Item []*DutchBusinessGetConcernRelationsDetailsGraphLink `xml:"item,omitempty"`
}

type DutchBusinessDossierV2 struct {
	Update_info *DutchBusinessUpdateReference `xml:"update_info,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment_number string `xml:"main_establishment_number,omitempty"`

	Indication_main_establishment bool `xml:"indication_main_establishment,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Chamber_number string `xml:"chamber_number,omitempty"`

	Legal_form_code string `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Indication_organisation_code string `xml:"indication_organisation_code,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Trade_name_45 string `xml:"trade_name_45,omitempty"`

	Trade_name_full string `xml:"trade_name_full,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Correspondence_postcode string `xml:"correspondence_postcode,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Correspondence_house_number int32 `xml:"correspondence_house_number,omitempty"`

	Correspondence_house_number_addition string `xml:"correspondence_house_number_addition,omitempty"`

	Correspondence_country string `xml:"correspondence_country,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Mobile_number string `xml:"mobile_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Contact_title1 string `xml:"contact_title1,omitempty"`

	Contact_title2 string `xml:"contact_title2,omitempty"`

	Contact_initials string `xml:"contact_initials,omitempty"`

	Contact_prefix string `xml:"contact_prefix,omitempty"`

	Contact_surname string `xml:"contact_surname,omitempty"`

	Contact_gender string `xml:"contact_gender,omitempty"`

	Primary_sbi_code string `xml:"primary_sbi_code,omitempty"`

	Secondary_sbi_code1 string `xml:"secondary_sbi_code1,omitempty"`

	Secondary_sbi_code2 string `xml:"secondary_sbi_code2,omitempty"`

	Primary_sbi_code_text string `xml:"primary_sbi_code_text,omitempty"`

	Secondary_sbi_code1_text string `xml:"secondary_sbi_code1_text,omitempty"`

	Secondary_sbi_code2_text string `xml:"secondary_sbi_code2_text,omitempty"`

	Personnel int32 `xml:"personnel,omitempty"`

	Class_personnel int32 `xml:"class_personnel,omitempty"`

	Personnel_fulltime int32 `xml:"personnel_fulltime,omitempty"`

	Class_personnel_fulltime int32 `xml:"class_personnel_fulltime,omitempty"`

	Personnel_reference_date *DutchBusinessDateV2 `xml:"personnel_reference_date,omitempty"`

	Personnel_ci int32 `xml:"personnel_ci,omitempty"`

	Class_personnel_ci int32 `xml:"class_personnel_ci,omitempty"`

	Personnel_ci_reference_date *DutchBusinessDateV2 `xml:"personnel_ci_reference_date,omitempty"`

	Indication_import bool `xml:"indication_import,omitempty"`

	Indication_export bool `xml:"indication_export,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`

	Indication_non_mailing bool `xml:"indication_non_mailing,omitempty"`

	Indication_bankruptcy bool `xml:"indication_bankruptcy,omitempty"`

	Indication_dip bool `xml:"indication_dip,omitempty"`

	Authorized_share_capital int64 `xml:"authorized_share_capital,omitempty"`

	Authorized_share_capital_currency string `xml:"authorized_share_capital_currency,omitempty"`

	Paid_up_share_capital int64 `xml:"paid_up_share_capital,omitempty"`

	Paid_up_share_capital_currency string `xml:"paid_up_share_capital_currency,omitempty"`

	Issued_share_capital int64 `xml:"issued_share_capital,omitempty"`

	Issued_share_capital_currency string `xml:"issued_share_capital_currency,omitempty"`

	Founding_date *DutchBusinessDateV2 `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDateV2 `xml:"discontinuation_date,omitempty"`

	Continuation_date *DutchBusinessDateV2 `xml:"continuation_date,omitempty"`

	Establishment_date *DutchBusinessDateV2 `xml:"establishment_date,omitempty"`

	Annual_financial_statement_summary *DutchBusinessAnnualFinancialStatementSummary `xml:"annual_financial_statement_summary,omitempty"`

	Structure *DutchBusinessStructure `xml:"structure,omitempty"`

	Sbi_collection *DutchBusinessSBICollection `xml:"sbi_collection,omitempty"`
}

type DutchBusinessDossierV3 struct {
	Update_info *DutchBusinessUpdateReference `xml:"update_info,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment_number string `xml:"main_establishment_number,omitempty"`

	Indication_main_establishment bool `xml:"indication_main_establishment,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Chamber_number string `xml:"chamber_number,omitempty"`

	Legal_form_code string `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Indication_organisation_code string `xml:"indication_organisation_code,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Trade_name_45 string `xml:"trade_name_45,omitempty"`

	Trade_name_full string `xml:"trade_name_full,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Establishment_address *DutchBusinessFormattedAddress `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessFormattedAddress `xml:"correspondence_address,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Mobile_number string `xml:"mobile_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Contact_title1 string `xml:"contact_title1,omitempty"`

	Contact_title2 string `xml:"contact_title2,omitempty"`

	Contact_initials string `xml:"contact_initials,omitempty"`

	Contact_prefix string `xml:"contact_prefix,omitempty"`

	Contact_surname string `xml:"contact_surname,omitempty"`

	Contact_gender string `xml:"contact_gender,omitempty"`

	Primary_sbi_code string `xml:"primary_sbi_code,omitempty"`

	Secondary_sbi_code1 string `xml:"secondary_sbi_code1,omitempty"`

	Secondary_sbi_code2 string `xml:"secondary_sbi_code2,omitempty"`

	Primary_sbi_code_text string `xml:"primary_sbi_code_text,omitempty"`

	Secondary_sbi_code1_text string `xml:"secondary_sbi_code1_text,omitempty"`

	Secondary_sbi_code2_text string `xml:"secondary_sbi_code2_text,omitempty"`

	Personnel int32 `xml:"personnel,omitempty"`

	Class_personnel int32 `xml:"class_personnel,omitempty"`

	Personnel_fulltime int32 `xml:"personnel_fulltime,omitempty"`

	Class_personnel_fulltime int32 `xml:"class_personnel_fulltime,omitempty"`

	Personnel_reference_date *DutchBusinessDateV3 `xml:"personnel_reference_date,omitempty"`

	Personnel_ci int32 `xml:"personnel_ci,omitempty"`

	Class_personnel_ci int32 `xml:"class_personnel_ci,omitempty"`

	Personnel_ci_reference_date *DutchBusinessDateV3 `xml:"personnel_ci_reference_date,omitempty"`

	Indication_import bool `xml:"indication_import,omitempty"`

	Indication_export bool `xml:"indication_export,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`

	Indication_non_mailing bool `xml:"indication_non_mailing,omitempty"`

	Indication_bankruptcy bool `xml:"indication_bankruptcy,omitempty"`

	Indication_dip bool `xml:"indication_dip,omitempty"`

	Insolvencies *DutchBusinessInsolvencyArray `xml:"insolvencies,omitempty"`

	Authorized_share_capital int64 `xml:"authorized_share_capital,omitempty"`

	Authorized_share_capital_currency string `xml:"authorized_share_capital_currency,omitempty"`

	Paid_up_share_capital int64 `xml:"paid_up_share_capital,omitempty"`

	Paid_up_share_capital_currency string `xml:"paid_up_share_capital_currency,omitempty"`

	Issued_share_capital int64 `xml:"issued_share_capital,omitempty"`

	Issued_share_capital_currency string `xml:"issued_share_capital_currency,omitempty"`

	Founding_date *DutchBusinessDateV3 `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDateV3 `xml:"discontinuation_date,omitempty"`

	Continuation_date *DutchBusinessDateV3 `xml:"continuation_date,omitempty"`

	Establishment_date *DutchBusinessDateV3 `xml:"establishment_date,omitempty"`

	Annual_financial_statement_summary *DutchBusinessAnnualFinancialStatementSummary `xml:"annual_financial_statement_summary,omitempty"`

	Structure *DutchBusinessStructure `xml:"structure,omitempty"`

	Sbi_collection *DutchBusinessSBICollection `xml:"sbi_collection,omitempty"`
}

type DutchBusinessAnnualFinancialStatementSummary struct {
	Year int32 `xml:"year,omitempty"`

	Turnover *DutchBusinessMoneyV2 `xml:"turnover,omitempty"`

	Profit *DutchBusinessMoneyV2 `xml:"profit,omitempty"`

	Assets *DutchBusinessMoneyV2 `xml:"assets,omitempty"`
}

type DutchBusinessStructure struct {
	Ultimate_parent string `xml:"ultimate_parent,omitempty"`

	Number_of_subsidiaries int32 `xml:"number_of_subsidiaries,omitempty"`

	Parent string `xml:"parent,omitempty"`
}

type DutchBusinessExtractDocument struct {
	Document []byte `xml:"document,omitempty"`

	Document_date time.Time `xml:"document_date,omitempty"`
}

type DutchBusinessExtractDocumentData struct {
	Document []byte `xml:"document,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Status string `xml:"status,omitempty"`

	Document_date time.Time `xml:"document_date,omitempty"`

	Legal_entity *DutchBusinessLegalEntity `xml:"legal_entity,omitempty"`

	Enterprise *DutchBusinessEnterprise `xml:"enterprise,omitempty"`

	Partnership *DutchBusinessPartnershipArray `xml:"partnership,omitempty"`

	Establishments *DutchBusinessEstablishmentArray `xml:"establishments,omitempty"`

	Positions *DutchBusinessPositionArray `xml:"positions,omitempty"`
}

type DutchBusinessExtractDocumentDataV2 struct {
	Document []byte `xml:"document,omitempty"`

	Data *DutchBusinessExtractDataV2 `xml:"data,omitempty"`
}

type DutchBusinessExtractDocumentDataV3 struct {
	Document []byte `xml:"document,omitempty"`

	Data *DutchBusinessExtractDataV3 `xml:"data,omitempty"`

	Source []byte `xml:"source,omitempty"`
}

type DutchBusinessExtractDataV2 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Reference_date time.Time `xml:"reference_date,omitempty"`

	Legal_entity *DutchBusinessLegalEntityV2 `xml:"legal_entity,omitempty"`

	Partnership *DutchBusinessPartnershipV2 `xml:"partnership,omitempty"`

	Enterprise *DutchBusinessEnterpriseV2 `xml:"enterprise,omitempty"`

	Establishments *DutchBusinessEstablishmentV2Array `xml:"establishments,omitempty"`

	Positions *DutchBusinessPositionV2Array `xml:"positions,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`
}

type DutchBusinessExtractDataV3 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Reference_date time.Time `xml:"reference_date,omitempty"`

	Legal_entity *DutchBusinessLegalEntityV3 `xml:"legal_entity,omitempty"`

	Partnership *DutchBusinessPartnershipV3 `xml:"partnership,omitempty"`

	Enterprise *DutchBusinessEnterpriseV3 `xml:"enterprise,omitempty"`

	Establishments *DutchBusinessEstablishmentV3Array `xml:"establishments,omitempty"`

	Positions *DutchBusinessPositionV3Array `xml:"positions,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`
}

type DutchBusinessPositions struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Reference_date time.Time `xml:"reference_date,omitempty"`

	Positions *DutchBusinessPositionV2Array `xml:"positions,omitempty"`
}

type DutchBusinessLegalEntityData struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Reference_date time.Time `xml:"reference_date,omitempty"`

	Legal_entity *DutchBusinessLegalEntityV3 `xml:"legal_entity,omitempty"`
}

type DutchBusinessLegalEntity struct {
	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Alternative_name string `xml:"alternative_name,omitempty"`

	Shortened_name string `xml:"shortened_name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Status string `xml:"status,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Legal_form_change string `xml:"legal_form_change,omitempty"`

	Foreign_legal_form_description string `xml:"foreign_legal_form_description,omitempty"`

	Sbi_codes *StringArray `xml:"sbi_codes,omitempty"`

	Sbi_codes_text *StringArray `xml:"sbi_codes_text,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Company_arrangement string `xml:"company_arrangement,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Statutory_seat string `xml:"statutory_seat,omitempty"`

	Registration_date *DutchBusinessDate `xml:"registration_date,omitempty"`

	Founding_date *DutchBusinessDate `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDate `xml:"discontinuation_date,omitempty"`

	Dissolution_date *DutchBusinessDate `xml:"dissolution_date,omitempty"`

	Dissolution_reason string `xml:"dissolution_reason,omitempty"`

	Removal_date *DutchBusinessDate `xml:"removal_date,omitempty"`

	Registration_end_date *DutchBusinessDate `xml:"registration_end_date,omitempty"`

	Legal_entity_end_date *DutchBusinessDate `xml:"legal_entity_end_date,omitempty"`

	Liquidation_closure_date *DutchBusinessDate `xml:"liquidation_closure_date,omitempty"`

	Liquidation_reopening_date *DutchBusinessDate `xml:"liquidation_reopening_date,omitempty"`

	Deed_incorporation_date *DutchBusinessDate `xml:"deed_incorporation_date,omitempty"`

	Deed_last_statutes_amendment_date *DutchBusinessDate `xml:"deed_last_statutes_amendment_date,omitempty"`

	Last_statutes_amendment_date *DutchBusinessDate `xml:"last_statutes_amendment_date,omitempty"`

	Liability string `xml:"liability,omitempty"`

	Merger_description string `xml:"merger_description,omitempty"`

	Annual_report_submission string `xml:"annual_report_submission,omitempty"`

	Authorized_share_capital int64 `xml:"authorized_share_capital,omitempty"`

	Authorized_share_capital_currency string `xml:"authorized_share_capital_currency,omitempty"`

	Issued_share_capital int64 `xml:"issued_share_capital,omitempty"`

	Issued_share_capital_currency string `xml:"issued_share_capital_currency,omitempty"`

	Paid_up_share_capital int64 `xml:"paid_up_share_capital,omitempty"`

	Paid_up_share_capital_currency string `xml:"paid_up_share_capital_currency,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Establishment_country string `xml:"establishment_country,omitempty"`

	Correspondence_postcode string `xml:"correspondence_postcode,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Correspondence_house_number int32 `xml:"correspondence_house_number,omitempty"`

	Correspondence_house_number_addition string `xml:"correspondence_house_number_addition,omitempty"`

	Correspondence_country string `xml:"correspondence_country,omitempty"`

	Duration string `xml:"duration,omitempty"`

	Duration_end_date *DutchBusinessDate `xml:"duration_end_date,omitempty"`

	Shares string `xml:"shares,omitempty"`

	Share_holders string `xml:"share_holders,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`
}

type DutchBusinessLegalEntityV2 struct {
	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Alternative_name string `xml:"alternative_name,omitempty"`

	Shortened_name string `xml:"shortened_name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Legal_form_change string `xml:"legal_form_change,omitempty"`

	Foreign_legal_form_description string `xml:"foreign_legal_form_description,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Company_arrangement string `xml:"company_arrangement,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Statutory_seat string `xml:"statutory_seat,omitempty"`

	Registration_date *DutchBusinessDateV2 `xml:"registration_date,omitempty"`

	Founding_date *DutchBusinessDateV2 `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDateV2 `xml:"discontinuation_date,omitempty"`

	Dissolution_date *DutchBusinessDateV2 `xml:"dissolution_date,omitempty"`

	Dissolution_reason string `xml:"dissolution_reason,omitempty"`

	Removal_date *DutchBusinessDateV2 `xml:"removal_date,omitempty"`

	Registration_end_date *DutchBusinessDateV2 `xml:"registration_end_date,omitempty"`

	Legal_entity_end_date *DutchBusinessDateV2 `xml:"legal_entity_end_date,omitempty"`

	Liquidation_closure_date *DutchBusinessDateV2 `xml:"liquidation_closure_date,omitempty"`

	Liquidation_reopening_date *DutchBusinessDateV2 `xml:"liquidation_reopening_date,omitempty"`

	Deed_incorporation_date *DutchBusinessDateV2 `xml:"deed_incorporation_date,omitempty"`

	Deed_last_statutes_amendment_date *DutchBusinessDateV2 `xml:"deed_last_statutes_amendment_date,omitempty"`

	Last_statutes_amendment_date *DutchBusinessDateV2 `xml:"last_statutes_amendment_date,omitempty"`

	Liability string `xml:"liability,omitempty"`

	Merger_description string `xml:"merger_description,omitempty"`

	Annual_report_submission string `xml:"annual_report_submission,omitempty"`

	Authorized_share_capital *DutchBusinessMoneyV2 `xml:"authorized_share_capital,omitempty"`

	Issued_share_capital *DutchBusinessMoneyV2 `xml:"issued_share_capital,omitempty"`

	Paid_up_share_capital *DutchBusinessMoneyV2 `xml:"paid_up_share_capital,omitempty"`

	Duration string `xml:"duration,omitempty"`

	Duration_end_date *DutchBusinessDateV2 `xml:"duration_end_date,omitempty"`

	Shares string `xml:"shares,omitempty"`

	Share_holders string `xml:"share_holders,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`

	Establishment_address *DutchBusinessAddressV2 `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV2 `xml:"correspondence_address,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`

	Sbi_codes *DutchBusinessSbiCodeV2Array `xml:"sbi_codes,omitempty"`
}

type DutchBusinessLegalEntityV3 struct {
	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Alternative_name string `xml:"alternative_name,omitempty"`

	Shortened_name string `xml:"shortened_name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Legal_form_change string `xml:"legal_form_change,omitempty"`

	Foreign_legal_form_description string `xml:"foreign_legal_form_description,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Company_arrangement string `xml:"company_arrangement,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Statutory_seat string `xml:"statutory_seat,omitempty"`

	Registration_date *DutchBusinessDateV3 `xml:"registration_date,omitempty"`

	Founding_date *DutchBusinessDateV3 `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDateV3 `xml:"discontinuation_date,omitempty"`

	Dissolution_date *DutchBusinessDateV3 `xml:"dissolution_date,omitempty"`

	Dissolution_reason string `xml:"dissolution_reason,omitempty"`

	Removal_date *DutchBusinessDateV3 `xml:"removal_date,omitempty"`

	Registration_end_date *DutchBusinessDateV3 `xml:"registration_end_date,omitempty"`

	Legal_entity_end_date *DutchBusinessDateV3 `xml:"legal_entity_end_date,omitempty"`

	Liquidation_closure_date *DutchBusinessDateV3 `xml:"liquidation_closure_date,omitempty"`

	Liquidation_reopening_date *DutchBusinessDateV3 `xml:"liquidation_reopening_date,omitempty"`

	Deed_incorporation_date *DutchBusinessDateV3 `xml:"deed_incorporation_date,omitempty"`

	Deed_last_statutes_amendment_date *DutchBusinessDateV3 `xml:"deed_last_statutes_amendment_date,omitempty"`

	Last_statutes_amendment_date *DutchBusinessDateV3 `xml:"last_statutes_amendment_date,omitempty"`

	Liability string `xml:"liability,omitempty"`

	Merger_description string `xml:"merger_description,omitempty"`

	Annual_report_submission string `xml:"annual_report_submission,omitempty"`

	Authorized_share_capital *DutchBusinessMoneyV3 `xml:"authorized_share_capital,omitempty"`

	Issued_share_capital *DutchBusinessMoneyV3 `xml:"issued_share_capital,omitempty"`

	Paid_up_share_capital *DutchBusinessMoneyV3 `xml:"paid_up_share_capital,omitempty"`

	Duration string `xml:"duration,omitempty"`

	Duration_end_date *DutchBusinessDateV3 `xml:"duration_end_date,omitempty"`

	Shares string `xml:"shares,omitempty"`

	Share_holders string `xml:"share_holders,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`

	Establishment_address *DutchBusinessAddressV3 `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV3 `xml:"correspondence_address,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`

	Sbi_codes *DutchBusinessSbiCodeV3Array `xml:"sbi_codes,omitempty"`
}

type DutchBusinessOrganizationTree struct {
	Name string `xml:"name,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Tree *DutchBusinessOrganizationNode `xml:"tree,omitempty"`
}

type DutchBusinessOrganizationNode struct {
	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Id string `xml:"id,omitempty"`

	Children *DutchBusinessOrganizationNodeArray `xml:"children,omitempty"`
}

type DutchBusinessOrganizationNodeArray struct {
	Item []*DutchBusinessOrganizationNode `xml:"item,omitempty"`
}

type DutchBusinessPartnership struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Status string `xml:"status,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Founding_date *DutchBusinessDate `xml:"founding_date,omitempty"`

	Dissolution_date *DutchBusinessDate `xml:"dissolution_date,omitempty"`

	Function_start_date *DutchBusinessDate `xml:"function_start_date,omitempty"`

	Duration string `xml:"duration,omitempty"`

	Limited_partnership_capital int64 `xml:"limited_partnership_capital,omitempty"`

	Limited_partnership_capital_currency string `xml:"limited_partnership_capital_currency,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Establishment_country string `xml:"establishment_country,omitempty"`

	Silent_partners int32 `xml:"silent_partners,omitempty"`
}

type DutchBusinessPartnershipArray struct {
	Item []*DutchBusinessPartnership `xml:"item,omitempty"`
}

type DutchBusinessPartnershipV2 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Founding_date *DutchBusinessDateV2 `xml:"founding_date,omitempty"`

	Dissolution_date *DutchBusinessDateV2 `xml:"dissolution_date,omitempty"`

	Function_start_date *DutchBusinessDateV2 `xml:"function_start_date,omitempty"`

	Duration string `xml:"duration,omitempty"`

	Limited_partnership_capital *DutchBusinessMoneyV2 `xml:"limited_partnership_capital,omitempty"`

	Establishment_address *DutchBusinessAddressV2 `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV2 `xml:"correspondence_address,omitempty"`

	Silent_partners int32 `xml:"silent_partners,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`
}

type DutchBusinessPartnershipV3 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Founding_date *DutchBusinessDateV3 `xml:"founding_date,omitempty"`

	Dissolution_date *DutchBusinessDateV3 `xml:"dissolution_date,omitempty"`

	Function_start_date *DutchBusinessDateV3 `xml:"function_start_date,omitempty"`

	Duration string `xml:"duration,omitempty"`

	Limited_partnership_capital *DutchBusinessMoneyV3 `xml:"limited_partnership_capital,omitempty"`

	Establishment_address *DutchBusinessAddressV3 `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV3 `xml:"correspondence_address,omitempty"`

	Silent_partners int32 `xml:"silent_partners,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`
}

type DutchBusinessEnterprise struct {
	Rsin_number string `xml:"rsin_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Status string `xml:"status,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Description string `xml:"description,omitempty"`

	Sbi_codes *StringArray `xml:"sbi_codes,omitempty"`

	Sbi_codes_text *StringArray `xml:"sbi_codes_text,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Establishment_country string `xml:"establishment_country,omitempty"`

	Correspondence_postcode string `xml:"correspondence_postcode,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Correspondence_house_number int32 `xml:"correspondence_house_number,omitempty"`

	Correspondence_house_number_addition string `xml:"correspondence_house_number_addition,omitempty"`

	Correspondence_country string `xml:"correspondence_country,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`

	Establishment_date *DutchBusinessDate `xml:"establishment_date,omitempty"`

	Incorporation_date *DutchBusinessDate `xml:"incorporation_date,omitempty"`

	Founding_date *DutchBusinessDate `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDate `xml:"discontinuation_date,omitempty"`

	Date_since *DutchBusinessDate `xml:"date_since,omitempty"`

	Personnel int32 `xml:"personnel,omitempty"`
}

type DutchBusinessEnterpriseV2 struct {
	Name string `xml:"name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Description string `xml:"description,omitempty"`

	Incorporation_date *DutchBusinessDateV2 `xml:"incorporation_date,omitempty"`

	Founding_date *DutchBusinessDateV2 `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDateV2 `xml:"discontinuation_date,omitempty"`

	Date_since *DutchBusinessDateV2 `xml:"date_since,omitempty"`

	Personnel *DutchBusinessPersonnelInfoV2 `xml:"personnel,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`

	Sbi_codes *DutchBusinessSbiCodeV2Array `xml:"sbi_codes,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DutchBusinessEnterpriseV3 struct {
	Name string `xml:"name,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Description string `xml:"description,omitempty"`

	Incorporation_date *DutchBusinessDateV3 `xml:"incorporation_date,omitempty"`

	Founding_date *DutchBusinessDateV3 `xml:"founding_date,omitempty"`

	Discontinuation_date *DutchBusinessDateV3 `xml:"discontinuation_date,omitempty"`

	Date_since *DutchBusinessDateV3 `xml:"date_since,omitempty"`

	Personnel *DutchBusinessPersonnelInfoV3 `xml:"personnel,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`

	Sbi_codes *DutchBusinessSbiCodeV3Array `xml:"sbi_codes,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DutchBusinessEstablishment struct {
	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment bool `xml:"main_establishment,omitempty"`

	Status string `xml:"status,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Description string `xml:"description,omitempty"`

	Sbi_codes *StringArray `xml:"sbi_codes,omitempty"`

	Sbi_codes_text *StringArray `xml:"sbi_codes_text,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Establishment_country string `xml:"establishment_country,omitempty"`

	Correspondence_postcode string `xml:"correspondence_postcode,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Correspondence_house_number int32 `xml:"correspondence_house_number,omitempty"`

	Correspondence_house_number_addition string `xml:"correspondence_house_number_addition,omitempty"`

	Correspondence_country string `xml:"correspondence_country,omitempty"`

	Curator_postcode string `xml:"curator_postcode,omitempty"`

	Curator_city string `xml:"curator_city,omitempty"`

	Curator_street string `xml:"curator_street,omitempty"`

	Curator_house_number int32 `xml:"curator_house_number,omitempty"`

	Curator_house_number_addition string `xml:"curator_house_number_addition,omitempty"`

	Curator_country string `xml:"curator_country,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`

	Establishment_date *DutchBusinessDate `xml:"establishment_date,omitempty"`

	Date_since *DutchBusinessDate `xml:"date_since,omitempty"`

	Personnel int32 `xml:"personnel,omitempty"`
}

type DutchBusinessEstablishmentArray struct {
	Item []*DutchBusinessEstablishment `xml:"item,omitempty"`
}

type DutchBusinessEstablishmentV2 struct {
	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment bool `xml:"main_establishment,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Description string `xml:"description,omitempty"`

	Establishment_date *DutchBusinessDateV2 `xml:"establishment_date,omitempty"`

	Date_since *DutchBusinessDateV2 `xml:"date_since,omitempty"`

	Personnel *DutchBusinessPersonnelInfoV2 `xml:"personnel,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`

	Sbi_codes *DutchBusinessSbiCodeV2Array `xml:"sbi_codes,omitempty"`

	Establishment_address *DutchBusinessAddressV2 `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV2 `xml:"correspondence_address,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DutchBusinessEstablishmentV2Array struct {
	Item []*DutchBusinessEstablishmentV2 `xml:"item,omitempty"`
}

type DutchBusinessEstablishmentV3 struct {
	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment bool `xml:"main_establishment,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Activity string `xml:"activity,omitempty"`

	Description string `xml:"description,omitempty"`

	Establishment_date *DutchBusinessDateV3 `xml:"establishment_date,omitempty"`

	Date_since *DutchBusinessDateV3 `xml:"date_since,omitempty"`

	Personnel *DutchBusinessPersonnelInfoV3 `xml:"personnel,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`

	Sbi_codes *DutchBusinessSbiCodeV3Array `xml:"sbi_codes,omitempty"`

	Establishment_address *DutchBusinessAddressV3 `xml:"establishment_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV3 `xml:"correspondence_address,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Fax_numbers *StringArray `xml:"fax_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Domain_names *StringArray `xml:"domain_names,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DutchBusinessEstablishmentV3Array struct {
	Item []*DutchBusinessEstablishmentV3 `xml:"item,omitempty"`
}

type DutchBusinessPosition struct {
	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Name string `xml:"name,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Title string `xml:"title,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Function_type string `xml:"function_type,omitempty"`

	Function_description string `xml:"function_description,omitempty"`

	Function_title string `xml:"function_title,omitempty"`

	Function_start_date *DutchBusinessDate `xml:"function_start_date,omitempty"`

	Function_registration_date *DutchBusinessDate `xml:"function_registration_date,omitempty"`

	Function_end_date *DutchBusinessDate `xml:"function_end_date,omitempty"`

	Function_authorization string `xml:"function_authorization,omitempty"`

	Function_authorization_description string `xml:"function_authorization_description,omitempty"`

	Function_authorization_start_date *DutchBusinessDate `xml:"function_authorization_start_date,omitempty"`

	Function_authorization_signing_power string `xml:"function_authorization_signing_power,omitempty"`

	Function_authorization_end_date *DutchBusinessDate `xml:"function_authorization_end_date,omitempty"`

	Authorization_description string `xml:"authorization_description,omitempty"`

	Authorization_start_date *DutchBusinessDate `xml:"authorization_start_date,omitempty"`

	Authorization_end_date *DutchBusinessDate `xml:"authorization_end_date,omitempty"`

	Inauguration string `xml:"inauguration,omitempty"`

	Inauguration_function string `xml:"inauguration_function,omitempty"`

	Inauguration_duration string `xml:"inauguration_duration,omitempty"`

	Inauguration_date *DutchBusinessDate `xml:"inauguration_date,omitempty"`

	Inauguration_body string `xml:"inauguration_body,omitempty"`

	Under_receivership string `xml:"under_receivership,omitempty"`

	Rights_against_third_parties string `xml:"rights_against_third_parties,omitempty"`

	Release_of_covenant string `xml:"release_of_covenant,omitempty"`

	Date_deceased *DutchBusinessDate `xml:"date_deceased,omitempty"`

	Date_of_birth *DutchBusinessDate `xml:"date_of_birth,omitempty"`

	Place_of_birth string `xml:"place_of_birth,omitempty"`

	Country_of_birth string `xml:"country_of_birth,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Country string `xml:"country,omitempty"`

	Telephone_numbers *StringArray `xml:"telephone_numbers,omitempty"`

	Date_since *DutchBusinessDate `xml:"date_since,omitempty"`

	Date_joined *DutchBusinessDate `xml:"date_joined,omitempty"`

	Status string `xml:"status,omitempty"`
}

type DutchBusinessPositionArray struct {
	Item []*DutchBusinessPosition `xml:"item,omitempty"`
}

type DutchBusinessPositionV2 struct {
	Functionary *DutchBusinessPersonV2 `xml:"functionary,omitempty"`

	Organisation *DutchBusinessOrganizationReferenceV2 `xml:"organisation,omitempty"`

	Residential_address *DutchBusinessAddressV2 `xml:"residential_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV2 `xml:"correspondence_address,omitempty"`

	Establishment_address *DutchBusinessAddressV2 `xml:"establishment_address,omitempty"`

	Longest_serving bool `xml:"longest_serving,omitempty"`

	Function_type string `xml:"function_type,omitempty"`

	Function_description string `xml:"function_description,omitempty"`

	Function_title string `xml:"function_title,omitempty"`

	Function_start_date *DutchBusinessDateV2 `xml:"function_start_date,omitempty"`

	Function_registration_date *DutchBusinessDateV2 `xml:"function_registration_date,omitempty"`

	Function_end_date *DutchBusinessDateV2 `xml:"function_end_date,omitempty"`

	Function_authorization string `xml:"function_authorization,omitempty"`

	Function_authorization_description string `xml:"function_authorization_description,omitempty"`

	Function_authorization_start_date *DutchBusinessDateV2 `xml:"function_authorization_start_date,omitempty"`

	Function_authorization_signing_power string `xml:"function_authorization_signing_power,omitempty"`

	Function_authorization_end_date *DutchBusinessDateV2 `xml:"function_authorization_end_date,omitempty"`

	Authorization_description string `xml:"authorization_description,omitempty"`

	Authorization_establishment_number string `xml:"authorization_establishment_number,omitempty"`

	Authorization_start_date *DutchBusinessDateV2 `xml:"authorization_start_date,omitempty"`

	Authorization_end_date *DutchBusinessDateV2 `xml:"authorization_end_date,omitempty"`

	Authorization_constraints *StringArray `xml:"authorization_constraints,omitempty"`

	Inauguration string `xml:"inauguration,omitempty"`

	Inauguration_function string `xml:"inauguration_function,omitempty"`

	Inauguration_duration string `xml:"inauguration_duration,omitempty"`

	Inauguration_date *DutchBusinessDateV2 `xml:"inauguration_date,omitempty"`

	Inauguration_body string `xml:"inauguration_body,omitempty"`

	Under_receivership string `xml:"under_receivership,omitempty"`

	Rights_against_third_parties string `xml:"rights_against_third_parties,omitempty"`

	Release_of_covenant string `xml:"release_of_covenant,omitempty"`

	Date_since *DutchBusinessDateV2 `xml:"date_since,omitempty"`

	Date_joined *DutchBusinessDateV2 `xml:"date_joined,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`
}

type DutchBusinessPositionV2Array struct {
	Item []*DutchBusinessPositionV2 `xml:"item,omitempty"`
}

type DutchBusinessPositionV3 struct {
	Functionary *DutchBusinessPersonV3 `xml:"functionary,omitempty"`

	Organisation *DutchBusinessOrganizationReferenceV3 `xml:"organisation,omitempty"`

	Residential_address *DutchBusinessAddressV3 `xml:"residential_address,omitempty"`

	Correspondence_address *DutchBusinessAddressV3 `xml:"correspondence_address,omitempty"`

	Establishment_address *DutchBusinessAddressV3 `xml:"establishment_address,omitempty"`

	Longest_serving bool `xml:"longest_serving,omitempty"`

	Function_type string `xml:"function_type,omitempty"`

	Function_description string `xml:"function_description,omitempty"`

	Function_title string `xml:"function_title,omitempty"`

	Function_start_date *DutchBusinessDateV3 `xml:"function_start_date,omitempty"`

	Function_registration_date *DutchBusinessDateV3 `xml:"function_registration_date,omitempty"`

	Function_end_date *DutchBusinessDateV3 `xml:"function_end_date,omitempty"`

	Function_authorization string `xml:"function_authorization,omitempty"`

	Function_authorization_description string `xml:"function_authorization_description,omitempty"`

	Function_authorization_start_date *DutchBusinessDateV3 `xml:"function_authorization_start_date,omitempty"`

	Function_authorization_signing_power string `xml:"function_authorization_signing_power,omitempty"`

	Function_authorization_end_date *DutchBusinessDateV3 `xml:"function_authorization_end_date,omitempty"`

	Authorization_description string `xml:"authorization_description,omitempty"`

	Authorization_establishment_number string `xml:"authorization_establishment_number,omitempty"`

	Authorization_start_date *DutchBusinessDateV3 `xml:"authorization_start_date,omitempty"`

	Authorization_end_date *DutchBusinessDateV3 `xml:"authorization_end_date,omitempty"`

	Authorization_constraints *StringArray `xml:"authorization_constraints,omitempty"`

	Inauguration string `xml:"inauguration,omitempty"`

	Inauguration_function string `xml:"inauguration_function,omitempty"`

	Inauguration_duration string `xml:"inauguration_duration,omitempty"`

	Inauguration_date *DutchBusinessDateV3 `xml:"inauguration_date,omitempty"`

	Inauguration_body string `xml:"inauguration_body,omitempty"`

	Under_receivership string `xml:"under_receivership,omitempty"`

	Rights_against_third_parties string `xml:"rights_against_third_parties,omitempty"`

	Release_of_covenant string `xml:"release_of_covenant,omitempty"`

	Date_since *DutchBusinessDateV3 `xml:"date_since,omitempty"`

	Date_joined *DutchBusinessDateV3 `xml:"date_joined,omitempty"`

	Remarks *StringArray `xml:"remarks,omitempty"`
}

type DutchBusinessPositionV3Array struct {
	Item []*DutchBusinessPositionV3 `xml:"item,omitempty"`
}

type DutchBusinessPersonV2 struct {
	Full_name string `xml:"full_name,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Title string `xml:"title,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Gender string `xml:"gender,omitempty"`

	Date_deceased *DutchBusinessDateV2 `xml:"date_deceased,omitempty"`

	Date_of_birth *DutchBusinessDateV2 `xml:"date_of_birth,omitempty"`

	Place_of_birth string `xml:"place_of_birth,omitempty"`

	Country_of_birth string `xml:"country_of_birth,omitempty"`
}

type DutchBusinessPersonV3 struct {
	Full_name string `xml:"full_name,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Title string `xml:"title,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Gender string `xml:"gender,omitempty"`

	Date_deceased *DutchBusinessDateV3 `xml:"date_deceased,omitempty"`

	Date_of_birth *DutchBusinessDateV3 `xml:"date_of_birth,omitempty"`

	Place_of_birth string `xml:"place_of_birth,omitempty"`

	Country_of_birth string `xml:"country_of_birth,omitempty"`
}

type DutchBusinessAdditionalPositionFunctionary struct {
	Full_name string `xml:"full_name,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Date_deceased time.Time `xml:"date_deceased,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Place_of_birth string `xml:"place_of_birth,omitempty"`

	Country_of_birth string `xml:"country_of_birth,omitempty"`
}

type DutchBusinessOrganizationReferenceV2 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Name string `xml:"name,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DutchBusinessOrganizationReferenceV3 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Registration string `xml:"registration,omitempty"`

	Name string `xml:"name,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`
}

type DutchBusinessMoneyV2 struct {
	Amount int64 `xml:"amount,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Formatted string `xml:"formatted,omitempty"`
}

type DutchBusinessMoneyV3 struct {
	Amount int64 `xml:"amount,omitempty"`

	Currency string `xml:"currency,omitempty"`

	Formatted string `xml:"formatted,omitempty"`
}

type DutchBusinessPersonnelInfoV2 struct {
	Fulltime int32 `xml:"fulltime,omitempty"`

	Parttime int32 `xml:"parttime,omitempty"`

	Total int32 `xml:"total,omitempty"`
}

type DutchBusinessPersonnelInfoV3 struct {
	Fulltime int32 `xml:"fulltime,omitempty"`

	Parttime int32 `xml:"parttime,omitempty"`

	Total int32 `xml:"total,omitempty"`
}

type DutchBusinessFormattedAddress struct {
	Original *DutchBusinessAddressV3 `xml:"original,omitempty"`

	Official *DutchBusinessAddressV3 `xml:"official,omitempty"`

	Formatted *DutchBusinessAddressV3 `xml:"formatted,omitempty"`
}

type DutchBusinessAddressV2 struct {
	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Country string `xml:"country,omitempty"`
}

type DutchBusinessAddressV3 struct {
	Address string `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Country string `xml:"country,omitempty"`
}

type DutchBusinessSBICodeInfo struct {
	Section *DutchBusinessSBISection `xml:"section,omitempty"`

	Sbi_codes *DutchBusinessSBICodeArray `xml:"sbi_codes,omitempty"`
}

type DutchBusinessSBISection struct {
	Section_code string `xml:"section_code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchBusinessSBICode struct {
	Sbi_code string `xml:"sbi_code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchBusinessSBICodeArray struct {
	Item []*DutchBusinessSBICode `xml:"item,omitempty"`
}

type DutchBusinessSbiCodeV2 struct {
	Sbi_code string `xml:"sbi_code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchBusinessSbiCodeV2Array struct {
	Item []*DutchBusinessSbiCodeV2 `xml:"item,omitempty"`
}

type DutchBusinessSbiCodeV3 struct {
	Sbi_code string `xml:"sbi_code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchBusinessSbiCodeV3Array struct {
	Item []*DutchBusinessSbiCodeV3 `xml:"item,omitempty"`
}

type DutchBusinessInsolvency struct {
	Insolvency_id string `xml:"insolvency_id,omitempty"`

	Events *DutchBusinessInsolvencyPublicationArray `xml:"events,omitempty"`
}

type DutchBusinessInsolvencyArray struct {
	Item []*DutchBusinessInsolvency `xml:"item,omitempty"`
}

type DutchBusinessInsolvencyPublication struct {
	Publication_id string `xml:"publication_id,omitempty"`

	Description string `xml:"description,omitempty"`

	Date *DutchBusinessDateV3 `xml:"date,omitempty"`
}

type DutchBusinessInsolvencyPublicationArray struct {
	Item []*DutchBusinessInsolvencyPublication `xml:"item,omitempty"`
}

type DutchBusinessReference struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`
}

type DutchBusinessReferenceArray struct {
	Item []*DutchBusinessReference `xml:"item,omitempty"`
}

type DutchBusinessReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessReferenceArray `xml:"results,omitempty"`
}

type DutchBusinessReferenceV2 struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Match_type string `xml:"match_type,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`
}

type DutchBusinessReferenceV2Array struct {
	Item []*DutchBusinessReferenceV2 `xml:"item,omitempty"`
}

type DutchBusinessReferenceV2PagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessReferenceV2Array `xml:"results,omitempty"`
}

type DutchBusinessEstablishmentReference struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Match_type string `xml:"match_type,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Indication_main_establishment bool `xml:"indication_main_establishment,omitempty"`
}

type DutchBusinessEstablishmentReferenceArray struct {
	Item []*DutchBusinessEstablishmentReference `xml:"item,omitempty"`
}

type DutchBusinessEstablishmentReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessEstablishmentReferenceArray `xml:"results,omitempty"`
}

type DutchBusinessUpdateSubscription struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Up_to_date bool `xml:"up_to_date,omitempty"`

	Pending_updates *StringArray `xml:"pending_updates,omitempty"`
}

type DutchBusinessUpdateSubscriptionArray struct {
	Item []*DutchBusinessUpdateSubscription `xml:"item,omitempty"`
}

type DutchBusinessUpdateSubscriptionPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessUpdateSubscriptionArray `xml:"results,omitempty"`
}

type DutchBusinessUpdateReference struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Date_last_update time.Time `xml:"date_last_update,omitempty"`
}

type DutchBusinessUpdateReferenceArray struct {
	Item []*DutchBusinessUpdateReference `xml:"item,omitempty"`
}

type DutchBusinessUpdateReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessUpdateReferenceArray `xml:"results,omitempty"`
}

type DutchBusinessVatNumber struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Vat_number string `xml:"vat_number,omitempty"`

	Date_last_update time.Time `xml:"date_last_update,omitempty"`
}

type DutchBusinessDate struct {
	Year int32 `xml:"Year,omitempty"`

	Month int32 `xml:"Month,omitempty"`

	Day int32 `xml:"Day,omitempty"`
}

type DutchBusinessDateV2 struct {
	Year int32 `xml:"year,omitempty"`

	Month int32 `xml:"month,omitempty"`

	Day int32 `xml:"day,omitempty"`
}

type DutchBusinessDateV3 struct {
	Year int32 `xml:"year,omitempty"`

	Month int32 `xml:"month,omitempty"`

	Day int32 `xml:"day,omitempty"`
}

type DutchBusinessExtractHistory struct {
	Forecast *DutchBusinessExtractChangeForecast `xml:"forecast,omitempty"`

	References *DutchBusinessExtractReferenceArray `xml:"references,omitempty"`
}

type DutchBusinessExtractReferenceArray struct {
	Item []*DutchBusinessExtractReference `xml:"item,omitempty"`
}

type DutchBusinessExtractChangeForecast struct {
	Change_set *DutchBusinessExtractChangeSet `xml:"change_set,omitempty"`
}

type DutchBusinessExtractReference struct {
	Extract_id string `xml:"extract_id,omitempty"`

	Document_date time.Time `xml:"document_date,omitempty"`

	Change_set *DutchBusinessExtractChangeSet `xml:"change_set,omitempty"`
}

type DutchBusinessExtractChangeSet struct {
	Period *DutchBusinessPeriod `xml:"period,omitempty"`

	Changes *StringArray `xml:"changes,omitempty"`
}

type DutchBusinessDossierHistory struct {
	References *DutchBusinessDossierReferenceArray `xml:"references,omitempty"`
}

type DutchBusinessDossierReferenceArray struct {
	Item []*DutchBusinessDossierReference `xml:"item,omitempty"`
}

type DutchBusinessDossierReference struct {
	Update_date time.Time `xml:"update_date,omitempty"`

	Change_set *DutchBusinessChangeSet `xml:"change_set,omitempty"`
}

type DutchBusinessChangeSet struct {
	Period *DutchBusinessPeriod `xml:"period,omitempty"`

	Changes *StringArray `xml:"changes,omitempty"`
}

type DutchBusinessPeriod struct {
	Start_date time.Time `xml:"start_date,omitempty"`

	End_date time.Time `xml:"end_date,omitempty"`
}

type DutchBusinessNewsItem struct {
	Title string `xml:"title,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Short string `xml:"short,omitempty"`

	Url string `xml:"url,omitempty"`

	Source string `xml:"source,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Sentiment float64 `xml:"sentiment,omitempty"`

	Topics *StringArray `xml:"topics,omitempty"`

	Companies *DutchBusinessNewsCompanyArray `xml:"companies,omitempty"`
}

type DutchBusinessNewsItemArray struct {
	Item []*DutchBusinessNewsItem `xml:"item,omitempty"`
}

type DutchBusinessNewsItemPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessNewsItemArray `xml:"results,omitempty"`
}

type DutchBusinessNewsCompany struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Name string `xml:"name,omitempty"`
}

type DutchBusinessNewsCompanyArray struct {
	Item []*DutchBusinessNewsCompany `xml:"item,omitempty"`
}

type DutchBusinessUBOInvestigationToken struct {
	Token string `xml:"token,omitempty"`
}

type DutchBusinessUBOInvestigationStatus struct {
	Status string `xml:"status,omitempty"`

	Date_since time.Time `xml:"date_since,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchBusinessUBOInvestigationResult struct {
	Name string `xml:"name,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Persons *DutchBusinessUBOPersonArray `xml:"persons,omitempty"`

	Organisations *DutchBusinessUBOOrganisationArray `xml:"organisations,omitempty"`

	Tree *DutchBusinessUBOOrganizationNode `xml:"tree,omitempty"`
}

type DutchBusinessUBOOrganizationNode struct {
	Level int32 `xml:"level,omitempty"`

	Id string `xml:"id,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Role_type string `xml:"role_type,omitempty"`

	Registered_in string `xml:"registered_in,omitempty"`

	Name string `xml:"name,omitempty"`

	Parents *DutchBusinessUBOOrganizationNodeArray `xml:"parents,omitempty"`
}

type DutchBusinessUBOOrganizationNodeArray struct {
	Item []*DutchBusinessUBOOrganizationNode `xml:"item,omitempty"`
}

type DutchBusinessUBOPerson struct {
	Key string `xml:"key,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Middle_names string `xml:"middle_names,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Gender string `xml:"gender,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Place_of_birth string `xml:"place_of_birth,omitempty"`

	Country_of_birth string `xml:"country_of_birth,omitempty"`

	Date_of_death time.Time `xml:"date_of_death,omitempty"`

	Roles *DutchBusinessUBORoleArray `xml:"roles,omitempty"`
}

type DutchBusinessUBOPersonArray struct {
	Item []*DutchBusinessUBOPerson `xml:"item,omitempty"`
}

type DutchBusinessUBORole struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Title string `xml:"title,omitempty"`

	Role_type string `xml:"role_type,omitempty"`

	Competency string `xml:"competency,omitempty"`

	Start_role_date time.Time `xml:"start_role_date,omitempty"`
}

type DutchBusinessUBORoleArray struct {
	Item []*DutchBusinessUBORole `xml:"item,omitempty"`
}

type DutchBusinessUBOOrganisation struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Legal_form_code string `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	City string `xml:"city,omitempty"`

	Extract *DutchBusinessUBOExtract `xml:"extract,omitempty"`
}

type DutchBusinessUBOOrganisationArray struct {
	Item []*DutchBusinessUBOOrganisation `xml:"item,omitempty"`
}

type DutchBusinessUBOExtract struct {
	Document string `xml:"document,omitempty"`

	Source string `xml:"source,omitempty"`

	Data *DutchBusinessExtractDocumentData `xml:"data,omitempty"`

	Extract_date time.Time `xml:"extract_date,omitempty"`
}

type DutchBusinessUBOCostsInvestigationResult struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOCostsInvestigationResult"`

	Token string `xml:"token,omitempty"`

	Status string `xml:"status,omitempty"`

	Receipt *DutchBusinessUBOReceipt `xml:"receipt,omitempty"`

	Payment *DutchBusinessUBOPayment `xml:"payment,omitempty"`
}

type DutchBusinessUBOReceipt struct {
	Total_price string `xml:"total_price,omitempty"`

	Items *DutchBusinessUBOReceiptItemArray `xml:"items,omitempty"`
}

type DutchBusinessUBOReceiptItem struct {
	Description string `xml:"description,omitempty"`

	Quantity int32 `xml:"quantity,omitempty"`

	Unit_price string `xml:"unit_price,omitempty"`

	Total_price string `xml:"total_price,omitempty"`
}

type DutchBusinessUBOReceiptItemArray struct {
	Item []*DutchBusinessUBOReceiptItem `xml:"item,omitempty"`
}

type DutchBusinessUBOClassificationResult struct {
	Token string `xml:"token,omitempty"`

	Scenario string `xml:"scenario,omitempty"`

	Suspects *DutchBusinessUBOClassifiedSuspectArray `xml:"suspects,omitempty"`
}

type DutchBusinessUBOClassifiedSuspect struct {
	Key string `xml:"key,omitempty"`

	Person *DutchBusinessUBOClassifiedPerson `xml:"person,omitempty"`

	Role *DutchBusinessUBOClassifiedRole `xml:"role,omitempty"`

	Classification string `xml:"classification,omitempty"`
}

type DutchBusinessUBOClassifiedSuspectArray struct {
	Item []*DutchBusinessUBOClassifiedSuspect `xml:"item,omitempty"`
}

type DutchBusinessUBOClassifiedPerson struct {
	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`
}

type DutchBusinessUBOClassifiedRole struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Organisation_name string `xml:"organisation_name,omitempty"`

	Title string `xml:"title,omitempty"`

	Function_type string `xml:"function_type,omitempty"`
}

type DutchBusinessUBOPayment struct {
	Pre string `xml:"pre,omitempty"`

	Cost string `xml:"cost,omitempty"`

	Refund string `xml:"refund,omitempty"`
}

type DutchBusinessSBICollection struct {
	Original *DutchBusinessSBICodeArray `xml:"original,omitempty"`

	Company_info *DutchBusinessSBICodeArray `xml:"company_info,omitempty"`
}

type DutchBusinessSBICollectionArray struct {
	Item []*DutchBusinessSBICollection `xml:"item,omitempty"`
}

type DutchBusinessAnnualFinancialStatement struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Year int32 `xml:"year,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Document []byte `xml:"document,omitempty"`
}

type DutchBusinessGetLeiResult struct {
	Lei_code string `xml:"lei_code,omitempty"`

	Business_entity *DutchBusinessLeiBusinessEntity `xml:"business_entity,omitempty"`

	Registration *DutchBusinessLeiRegistration `xml:"registration,omitempty"`
}

type DutchBusinessLeiRegistration struct {
	Last_update time.Time `xml:"last_update,omitempty"`

	Status string `xml:"status,omitempty"`

	Assignment_date time.Time `xml:"assignment_date,omitempty"`

	Renewal_date time.Time `xml:"renewal_date,omitempty"`

	Managing_lou string `xml:"managing_lou,omitempty"`

	Validation_sources string `xml:"validation_sources,omitempty"`
}

type DutchBusinessLeiBusinessEntity struct {
	Expiration_date time.Time `xml:"expiration_date,omitempty"`

	Expiration_reason string `xml:"expiration_reason,omitempty"`

	Successor *DutchBusinessLeiSuccessor `xml:"successor,omitempty"`

	Company_name string `xml:"company_name,omitempty"`

	Legal_status string `xml:"legal_status,omitempty"`

	Legal_form *DutchBusinessLeiLegalForm `xml:"legal_form,omitempty"`

	Headquarters_address *DutchBusinessLeiAddress `xml:"headquarters_address,omitempty"`

	Legal_address *DutchBusinessLeiAddress `xml:"legal_address,omitempty"`

	Registration_authority *DutchBusinessLeiRegistrationAuthority `xml:"registration_authority,omitempty"`
}

type DutchBusinessLeiLegalForm struct {
	Legal_form_code string `xml:"legal_form_code,omitempty"`
}

type DutchBusinessLeiRegistrationAuthority struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Registration_authority_id string `xml:"registration_authority_id,omitempty"`
}

type DutchBusinessLeiAddress struct {
	Street_and_number string `xml:"street_and_number,omitempty"`

	City string `xml:"city,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	Region string `xml:"region,omitempty"`

	Country string `xml:"country,omitempty"`
}

type DutchBusinessLeiSuccessor struct {
	Name string `xml:"name,omitempty"`

	Lei_code string `xml:"lei_code,omitempty"`
}

type DutchBusinessNonMarketingIndicator struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Non_marketing_indicator bool `xml:"non_marketing_indicator,omitempty"`

	Since time.Time `xml:"since,omitempty"`
}

type DutchBusinessGetAdditionalPositionsResult struct {
	Functionary *DutchBusinessAdditionalPositionFunctionary `xml:"functionary,omitempty"`

	Current_positions *DutchBusinessAdditionalPositionArray `xml:"current_positions,omitempty"`

	Previous_positions *DutchBusinessAdditionalPositionArray `xml:"previous_positions,omitempty"`

	Possible_current_positions *DutchBusinessAdditionalPositionArray `xml:"possible_current_positions,omitempty"`

	Possible_previous_positions *DutchBusinessAdditionalPositionArray `xml:"possible_previous_positions,omitempty"`
}

type DutchBusinessAdditionalPosition struct {
	Start_date time.Time `xml:"start_date,omitempty"`

	End_date time.Time `xml:"end_date,omitempty"`

	End_date_source string `xml:"end_date_source,omitempty"`

	Role string `xml:"role,omitempty"`

	Function string `xml:"function,omitempty"`

	Authorization_description string `xml:"authorization_description,omitempty"`

	Historical_authorizations *DutchBusinessPositionAuthorizationArray `xml:"historical_authorizations,omitempty"`

	Company *DutchBusinessPositionCompany `xml:"company,omitempty"`
}

type DutchBusinessAdditionalPositionArray struct {
	Item []*DutchBusinessAdditionalPosition `xml:"item,omitempty"`
}

type DutchBusinessPositionCompany struct {
	Kvk string `xml:"kvk,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Trade_name_full string `xml:"trade_name_full,omitempty"`

	Primary_sbi_code int32 `xml:"primary_sbi_code,omitempty"`

	Financial_status string `xml:"financial_status,omitempty"`
}

type DutchBusinessPositionAuthorization struct {
	Start_date time.Time `xml:"start_date,omitempty"`

	End_date time.Time `xml:"end_date,omitempty"`

	End_date_source string `xml:"end_date_source,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchBusinessPositionAuthorizationArray struct {
	Item []*DutchBusinessPositionAuthorization `xml:"item,omitempty"`
}

type DutchBusinessLookALikeCompany struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Name string `xml:"name,omitempty"`

	Similarity_score float64 `xml:"similarity_score,omitempty"`

	Individual_features *DutchBusinessLookALikeIndividualFeatureArray `xml:"individual_features,omitempty"`

	Profile_url string `xml:"profile_url,omitempty"`
}

type DutchBusinessLookALikeCompanyArray struct {
	Item []*DutchBusinessLookALikeCompany `xml:"item,omitempty"`
}

type DutchBusinessLookALikeCompanyPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *DutchBusinessLookALikeCompanyArray `xml:"results,omitempty"`
}

type DutchBusinessLookALikeIndividualFeature struct {
	Name string `xml:"name,omitempty"`

	Score float64 `xml:"score,omitempty"`

	Weight float64 `xml:"weight,omitempty"`
}

type DutchBusinessLookALikeIndividualFeatureArray struct {
	Item []*DutchBusinessLookALikeIndividualFeature `xml:"item,omitempty"`
}

type DutchVehicle struct {
	Basic_data *DutchVehicleData `xml:"basic_data,omitempty"`

	Status *DutchVehicleStatus `xml:"status,omitempty"`

	Registration *DutchVehicleRegistration `xml:"registration,omitempty"`

	Details *DutchVehicleFeature `xml:"details,omitempty"`

	Tax *DutchVehicleTaxData `xml:"tax,omitempty"`

	Technical_weight *DutchVehicleTechnicalWeight `xml:"technical_weight,omitempty"`

	Engine *DutchVehicleEngine `xml:"engine,omitempty"`

	Environmental_impact *DutchVehicleEnvironmentArray `xml:"environmental_impact,omitempty"`

	Axle_information *DutchVehicleAxleArray `xml:"axle_information,omitempty"`

	Type_approval *DutchVehicleTypeApproval `xml:"type_approval,omitempty"`

	Body_approval *DutchVehicleBodyApprovalArray `xml:"body_approval,omitempty"`

	Remark *DutchVehicleRemarkArray `xml:"remark,omitempty"`

	Recall *DutchVehicleRecall `xml:"recall,omitempty"`
}

type DutchVehicleV2 struct {
	Basic_data *DutchVehicleData `xml:"basic_data,omitempty"`

	Status *DutchVehicleStatus `xml:"status,omitempty"`

	Registration *DutchVehicleRegistration `xml:"registration,omitempty"`

	Details *DutchVehicleFeature `xml:"details,omitempty"`

	Tax *DutchVehicleTaxData `xml:"tax,omitempty"`

	Technical_weight *DutchVehicleTechnicalWeight `xml:"technical_weight,omitempty"`

	Engine *DutchVehicleEngine `xml:"engine,omitempty"`

	Environmental_impact *DutchVehicleEnvironmentArray `xml:"environmental_impact,omitempty"`

	Axle_information *DutchVehicleAxleArray `xml:"axle_information,omitempty"`

	Type_approval *DutchVehicleTypeApproval `xml:"type_approval,omitempty"`

	Body_approval *DutchVehicleBodyApprovalArray `xml:"body_approval,omitempty"`

	Remark *DutchVehicleRemarkArray `xml:"remark,omitempty"`

	Recall *DutchVehicleRecall `xml:"recall,omitempty"`

	Mileage *DutchVehicleMileage `xml:"mileage,omitempty"`
}

type DutchVehicleOwnerHistory struct {
	Ownerships *DutchVehicleOwnershipArray `xml:"ownerships,omitempty"`
}

type DutchVehiclePurchaseReference struct {
	Basic_data *DutchVehicleData `xml:"basic_data,omitempty"`

	Reference *DutchVehicleReference `xml:"reference,omitempty"`

	Purchase_data *DutchVehiclePurchase `xml:"purchase_data,omitempty"`
}

type DutchVehicleMileage struct {
	Registration_year int32 `xml:"registration_year,omitempty"`

	Indication string `xml:"indication,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchVehicleData struct {
	License_plate string `xml:"license_plate,omitempty"`

	Location_vin_number string `xml:"location_vin_number,omitempty"`

	Catalog_price int32 `xml:"catalog_price,omitempty"`

	Category string `xml:"category,omitempty"`

	Brand_code string `xml:"brand_code,omitempty"`

	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Body_style string `xml:"body_style,omitempty"`

	Color string `xml:"color,omitempty"`

	Secondary_color string `xml:"secondary_color,omitempty"`
}

type DutchVehicleReference struct {
	Version string `xml:"version,omitempty"`

	Date_first_admission time.Time `xml:"date_first_admission,omitempty"`

	Year_manufactored int32 `xml:"year_manufactored,omitempty"`

	Fuel_type string `xml:"fuel_type,omitempty"`

	Weight int32 `xml:"weight,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`
}

type DutchVehicleStatus struct {
	Insured bool `xml:"insured,omitempty"`

	Awaiting_inspection bool `xml:"awaiting_inspection,omitempty"`

	Missed bool `xml:"missed,omitempty"`

	Stolen bool `xml:"stolen,omitempty"`

	Exported bool `xml:"exported,omitempty"`

	Suspended bool `xml:"suspended,omitempty"`
}

type DutchVehicleRegistration struct {
	Date_apk_due_date time.Time `xml:"date_apk_due_date,omitempty"`

	Date_first_issuing time.Time `xml:"date_first_issuing,omitempty"`

	Date_first_admission time.Time `xml:"date_first_admission,omitempty"`

	Date_suspension_end_date time.Time `xml:"date_suspension_end_date,omitempty"`

	Date_tachograph_expiration time.Time `xml:"date_tachograph_expiration,omitempty"`

	Date_latest_name_registration time.Time `xml:"date_latest_name_registration,omitempty"`

	Ownerships int32 `xml:"ownerships,omitempty"`
}

type DutchVehicleOwnership struct {
	Type_ string `xml:"type,omitempty"`

	Date_start time.Time `xml:"date_start,omitempty"`

	Date_end time.Time `xml:"date_end,omitempty"`

	Duration int32 `xml:"duration,omitempty"`
}

type DutchVehicleOwnershipArray struct {
	Item []*DutchVehicleOwnership `xml:"item,omitempty"`
}

type DutchVehicleFeature struct {
	Length int32 `xml:"length,omitempty"`

	Width int32 `xml:"width,omitempty"`

	Doors int32 `xml:"doors,omitempty"`

	Seats int32 `xml:"seats,omitempty"`

	Standing_room int32 `xml:"standing_room,omitempty"`

	Weight int32 `xml:"weight,omitempty"`

	Axles int32 `xml:"axles,omitempty"`

	Wheels int32 `xml:"wheels,omitempty"`

	Wheel_base int32 `xml:"wheel_base,omitempty"`

	Distance_front_to_clutch int32 `xml:"distance_front_to_clutch,omitempty"`

	Distance_clutch_to_rear int32 `xml:"distance_clutch_to_rear,omitempty"`

	Speed_limit_difference int32 `xml:"speed_limit_difference,omitempty"`
}

type DutchVehicleTaxData struct {
	Amount int32 `xml:"amount,omitempty"`

	Base string `xml:"base,omitempty"`
}

type DutchVehiclePurchase struct {
	Reference_date time.Time `xml:"reference_date,omitempty"`

	Price_data *DutchVehiclePrice `xml:"price_data,omitempty"`
}

type DutchVehiclePrice struct {
	Purchase_price float64 `xml:"purchase_price,omitempty"`

	Vehicle_purchase_tax float64 `xml:"vehicle_purchase_tax,omitempty"`

	Value_added_tax float64 `xml:"value_added_tax,omitempty"`

	Net_purchase_price float64 `xml:"net_purchase_price,omitempty"`

	Calculation_method string `xml:"calculation_method,omitempty"`
}

type DutchVehicleMarketValue struct {
	Valuations *DutchVehiclePriceValuationArray `xml:"valuations,omitempty"`
}

type DutchVehiclePriceValuation struct {
	Brand string `xml:"brand,omitempty"`

	Model string `xml:"model,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Retail int32 `xml:"retail,omitempty"`

	Exchange int32 `xml:"exchange,omitempty"`

	Trade int32 `xml:"trade,omitempty"`

	Calculation_method string `xml:"calculation_method,omitempty"`
}

type DutchVehiclePriceValuationArray struct {
	Item []*DutchVehiclePriceValuation `xml:"item,omitempty"`
}

type DutchVehicleTechnicalWeight struct {
	Mass_ready int32 `xml:"mass_ready,omitempty"`

	Gross_vehicle_mass int32 `xml:"gross_vehicle_mass,omitempty"`

	Maximum_mass_payload int32 `xml:"maximum_mass_payload,omitempty"`

	Maximum_mass_unbraked int32 `xml:"maximum_mass_unbraked,omitempty"`

	Maximum_mass_braked int32 `xml:"maximum_mass_braked,omitempty"`

	Maximum_mass_trailer_braked int32 `xml:"maximum_mass_trailer_braked,omitempty"`

	Maximum_mass_self_braked int32 `xml:"maximum_mass_self_braked,omitempty"`

	Maximum_mass_axle_braked int32 `xml:"maximum_mass_axle_braked,omitempty"`
}

type DutchVehicleRemark struct {
	Vehicle_code string `xml:"vehicle_code,omitempty"`

	Variable_code string `xml:"variable_code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchVehicleRemarkArray struct {
	Item []*DutchVehicleRemark `xml:"item,omitempty"`
}

type DutchVehicleRecall struct {
	Status string `xml:"status,omitempty"`

	Defects string `xml:"defects,omitempty"`
}

type DutchVehicleEngine struct {
	Cylinder_capacity int32 `xml:"cylinder_capacity,omitempty"`

	Cylinders int32 `xml:"cylinders,omitempty"`

	G3_indication bool `xml:"g3_indication,omitempty"`

	Energy_label string `xml:"energy_label,omitempty"`
}

type DutchVehicleEnvironment struct {
	Fuel_type string `xml:"fuel_type,omitempty"`

	Fuel_code string `xml:"fuel_code,omitempty"`

	Usage_city float64 `xml:"usage_city,omitempty"`

	Usage_highway float64 `xml:"usage_highway,omitempty"`

	Usage_combined float64 `xml:"usage_combined,omitempty"`

	Max_power float64 `xml:"max_power,omitempty"`

	Max_power_continuous float64 `xml:"max_power_continuous,omitempty"`

	Noise_level_stationary int32 `xml:"noise_level_stationary,omitempty"`

	Noise_level_stationary_rpm int32 `xml:"noise_level_stationary_rpm,omitempty"`

	Noise_level_driving int32 `xml:"noise_level_driving,omitempty"`

	Emission_code string `xml:"emission_code,omitempty"`

	Co2_emission_combined int32 `xml:"co2_emission_combined,omitempty"`

	Co2_emission_weighted int32 `xml:"co2_emission_weighted,omitempty"`

	Particulate_emission float64 `xml:"particulate_emission,omitempty"`

	Particulate_filter string `xml:"particulate_filter,omitempty"`

	Emission_particles_light float64 `xml:"emission_particles_light,omitempty"`

	Emission_particles_heavy float64 `xml:"emission_particles_heavy,omitempty"`
}

type DutchVehicleEnvironmentArray struct {
	Item []*DutchVehicleEnvironment `xml:"item,omitempty"`
}

type DutchVehicleAxle struct {
	Number int32 `xml:"number,omitempty"`

	Indication_driven_axle bool `xml:"indication_driven_axle,omitempty"`

	Indication_bogie_lift bool `xml:"indication_bogie_lift,omitempty"`

	Max_axle_load int32 `xml:"max_axle_load,omitempty"`

	Max_axle_load_technical int32 `xml:"max_axle_load_technical,omitempty"`

	Location_code string `xml:"location_code,omitempty"`

	Track_width int32 `xml:"track_width,omitempty"`

	Road_behavior_code string `xml:"road_behavior_code,omitempty"`

	Road_behavior_description string `xml:"road_behavior_description,omitempty"`
}

type DutchVehicleAxleArray struct {
	Item []*DutchVehicleAxle `xml:"item,omitempty"`
}

type DutchVehicleTypeApproval struct {
	Type_code string `xml:"type_code,omitempty"`

	Version_code string `xml:"version_code,omitempty"`

	Type_approval_mark string `xml:"type_approval_mark,omitempty"`
}

type DutchVehicleBodyApproval struct {
	Type_id int32 `xml:"type_id,omitempty"`

	Type_code string `xml:"type_code,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DutchVehicleBodyApprovalArray struct {
	Item []*DutchVehicleBodyApproval `xml:"item,omitempty"`
}

type EDRScore struct {
	Person_score int32 `xml:"person_score,omitempty"`

	Region_score int32 `xml:"region_score,omitempty"`

	Edr_region_score int32 `xml:"edr_region_score,omitempty"`

	Weighted_score int32 `xml:"weighted_score,omitempty"`
}

type RDCoordinates struct {
	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`
}

type RDCoordinatesMatch struct {
	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`

	Citymatch string `xml:"citymatch,omitempty"`

	Streetmatch string `xml:"streetmatch,omitempty"`
}

type LatLonCoordinates struct {
	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type LatLonCoordinatesMatch struct {
	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`

	Citymatch string `xml:"citymatch,omitempty"`

	Streetmatch string `xml:"streetmatch,omitempty"`
}

type LatLonCoordinatesInternationalAddress struct {
	Postcode string `xml:"postcode,omitempty"`

	Street string `xml:"street,omitempty"`

	City string `xml:"city,omitempty"`

	Province string `xml:"province,omitempty"`

	Country_iso2 string `xml:"country_iso2,omitempty"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type LatLonCoordinatesInternationalAddressArray struct {
	Item []*LatLonCoordinatesInternationalAddress `xml:"item,omitempty"`
}

type GeoLocationAddress struct {
	Houseno_from int32 `xml:"houseno_from,omitempty"`

	Houseno_to int32 `xml:"houseno_to,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	City string `xml:"city,omitempty"`

	District string `xml:"district,omitempty"`

	Province string `xml:"province,omitempty"`

	Distance int32 `xml:"distance,omitempty"`

	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type GeoLocationAddressV2 struct {
	House_no int32 `xml:"house_no,omitempty"`

	House_no_addition string `xml:"house_no_addition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	City string `xml:"city,omitempty"`

	District string `xml:"district,omitempty"`

	Province string `xml:"province,omitempty"`

	Distance int32 `xml:"distance,omitempty"`

	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type GeoLocationInternationalAddress struct {
	Postcode string `xml:"postcode,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	City string `xml:"city,omitempty"`

	Province string `xml:"province,omitempty"`

	Country string `xml:"country,omitempty"`

	Country_iso3 string `xml:"country_iso3,omitempty"`

	Distance int32 `xml:"distance,omitempty"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type SortedPostcode struct {
	Postcode string `xml:"postcode,omitempty"`

	Distance int32 `xml:"distance,omitempty"`

	Precision string `xml:"precision,omitempty"`
}

type SortedPostcodeArray struct {
	Item []*SortedPostcode `xml:"item,omitempty"`
}

type SortedPostcodePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *SortedPostcodeArray `xml:"results,omitempty"`
}

type GraydonReference struct {
	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Street string `xml:"street,omitempty"`

	House_no int32 `xml:"house_no,omitempty"`

	House_no_addition string `xml:"house_no_addition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Residence string `xml:"residence,omitempty"`

	Country_iso2 string `xml:"country_iso2,omitempty"`

	Url string `xml:"url,omitempty"`

	Email string `xml:"email,omitempty"`
}

type GraydonReferenceArray struct {
	Item []*GraydonReference `xml:"item,omitempty"`
}

type GraydonCreditReport struct {
	Report_normal *GraydonCreditReportNormal `xml:"report_normal,omitempty"`

	Report_calamity *GraydonCreditReportCalamity `xml:"report_calamity,omitempty"`

	Report_alarm *GraydonCreditReportAlarm `xml:"report_alarm,omitempty"`

	Document []byte `xml:"document,omitempty"`

	Alarm *GCRAlarm `xml:"alarm,omitempty"`
}

type GraydonCreditReportNormal struct {
	Company *GCRCompanyItself `xml:"company,omitempty"`

	Company_relations *GCRCompanyRelations `xml:"company_relations,omitempty"`

	Indirect_company_relations *GCRCompanyRelations `xml:"indirect_company_relations,omitempty"`
}

type GraydonCreditReportCalamity struct {
	Company *GCRCompanyItselfCalamity `xml:"company,omitempty"`

	Company_relations *GCRCompanyRelationsSimple `xml:"company_relations,omitempty"`
}

type GraydonCreditReportAlarm struct {
	Company *GCRCompanyItselfAlarm `xml:"company,omitempty"`
}

type GraydonCreditReportQuickscan struct {
	Pd_rating string `xml:"pd_rating,omitempty"`

	Credit_flag_code string `xml:"credit_flag_code,omitempty"`

	Credit_flag_text string `xml:"credit_flag_text,omitempty"`

	Alarm *GCRAlarm `xml:"alarm,omitempty"`
}

type GraydonCreditReportRatings struct {
	Credit_advice_amount string `xml:"credit_advice_amount,omitempty"`

	Credit_advice_currency string `xml:"credit_advice_currency,omitempty"`

	Credit_advice_factors *GCRCreditFactorArray `xml:"credit_advice_factors,omitempty"`

	Pd_rating string `xml:"pd_rating,omitempty"`

	Ratings *GCRRatingArray `xml:"ratings,omitempty"`

	Payment_information *GCRPaymentInformationArray `xml:"payment_information,omitempty"`

	Alarm *GCRAlarm `xml:"alarm,omitempty"`
}

type GraydonCreditReportVatNumber struct {
	Vat_number string `xml:"vat_number,omitempty"`

	Alarm *GCRAlarm `xml:"alarm,omitempty"`
}

type GraydonCreditReportCompanyLiaisons struct {
	Concern_liaisons *GCRLiaisonArray `xml:"concern_liaisons,omitempty"`

	Companies *GCRLiaisonCompanyArray `xml:"companies,omitempty"`

	Alarm *GCRAlarm `xml:"alarm,omitempty"`
}

type GraydonCreditReportManagement struct {
	Company_management *GCRCompanyManagement `xml:"company_management,omitempty"`

	Persons *GCRPersonArray `xml:"persons,omitempty"`

	Alarm *GCRAlarm `xml:"alarm,omitempty"`
}

type GCRCompanyItself struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Contact_details *GCRContactDetails `xml:"contact_details,omitempty"`

	Official_data *GCROfficialData `xml:"official_data,omitempty"`

	History *GCRHistory `xml:"history,omitempty"`

	Positions_in_other_companies *GCRJobTitleArray `xml:"positions_in_other_companies,omitempty"`

	Annual_figures *GCRAnnualFiguresArray `xml:"annual_figures,omitempty"`

	Financial_details *GCRFinancialDetails `xml:"financial_details,omitempty"`

	Financial_calamities *GCRFinancialCalamityArray `xml:"financial_calamities,omitempty"`

	Personnel *GCRPersonnel `xml:"personnel,omitempty"`

	Credit_advice_data *GCRCreditAdviceData `xml:"credit_advice_data,omitempty"`

	Sectors_of_industry *GCRSectorOfIndustryArray `xml:"sectors_of_industry,omitempty"`

	Calamity *GCRCalamity `xml:"calamity,omitempty"`

	Payment_information *GCRPaymentInformationArray `xml:"payment_information,omitempty"`

	Declarations_of_liability *GCRDeclarationOfLiabilityArray `xml:"declarations_of_liability,omitempty"`

	Concern_liaisons *GCRLiaisonArray `xml:"concern_liaisons,omitempty"`

	Company_management *GCRCompanyManagement `xml:"company_management,omitempty"`
}

type GCRCompanyItselfCalamity struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Contact_details *GCRContactDetails `xml:"contact_details,omitempty"`

	Official_data *GCROfficialData `xml:"official_data,omitempty"`

	History *GCRHistory `xml:"history,omitempty"`

	Sectors_of_industry *GCRSectorOfIndustryArray `xml:"sectors_of_industry,omitempty"`

	Share_holders *GCRShareHolderArray `xml:"share_holders,omitempty"`

	Branch_offices *GCRBranchOfficeArray `xml:"branch_offices,omitempty"`

	Participations *GCRParticipationArray `xml:"participations,omitempty"`

	Import_export *GCRImportExport `xml:"import_export,omitempty"`

	Special_company_information *GCRSpecialCompanyInformationArray `xml:"special_company_information,omitempty"`

	Financial_calamities *GCRFinancialCalamityArray `xml:"financial_calamities,omitempty"`

	Concern_liaisons *GCRLiaisonArray `xml:"concern_liaisons,omitempty"`

	Company_management *GCRCompanyManagement `xml:"company_management,omitempty"`
}

type GCRCompanyItselfAlarm struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Contact_details *GCRAlarmContactDetails `xml:"contact_details,omitempty"`

	Discontinuance *GCRDiscontinuance `xml:"discontinuance,omitempty"`

	Special_company_information *GCRSpecialCompanyInformationArray `xml:"special_company_information,omitempty"`
}

type GCRAlarmContactDetails struct {
	Company_name string `xml:"company_name,omitempty"`

	Trade_names *GCRTradeNameArray `xml:"trade_names,omitempty"`

	Addresses *GCRAddressArray `xml:"addresses,omitempty"`
}

type GCRContactDetails struct {
	Company_name string `xml:"company_name,omitempty"`

	Trade_names *GCRTradeNameArray `xml:"trade_names,omitempty"`

	Addresses *GCRAddressArray `xml:"addresses,omitempty"`

	Telephone_numbers *GCRTelephoneNumberArray `xml:"telephone_numbers,omitempty"`
}

type GCRTradeName struct {
	Is_current bool `xml:"is_current,omitempty"`

	Name string `xml:"name,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`
}

type GCRTradeNameArray struct {
	Item []*GCRTradeName `xml:"item,omitempty"`
}

type GCRAddress struct {
	Is_current bool `xml:"is_current,omitempty"`

	Address_code int32 `xml:"address_code,omitempty"`

	Address_text string `xml:"address_text,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`

	Street string `xml:"street,omitempty"`

	House_no string `xml:"house_no,omitempty"`

	House_no_addition string `xml:"house_no_addition,omitempty"`

	Pobox string `xml:"pobox,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Residence string `xml:"residence,omitempty"`

	Country_code int32 `xml:"country_code,omitempty"`

	Country_text string `xml:"country_text,omitempty"`
}

type GCRAddressArray struct {
	Item []*GCRAddress `xml:"item,omitempty"`
}

type GCRTelephoneNumber struct {
	Telephone_code int32 `xml:"telephone_code,omitempty"`

	Telephone_text string `xml:"telephone_text,omitempty"`

	Country_number string `xml:"country_number,omitempty"`

	Net_number string `xml:"net_number,omitempty"`

	Subscriber_number string `xml:"subscriber_number,omitempty"`
}

type GCRTelephoneNumberArray struct {
	Item []*GCRTelephoneNumber `xml:"item,omitempty"`
}

type GCROfficialData struct {
	Chamber_of_commerce *GCRChamberOfCommerce `xml:"chamber_of_commerce,omitempty"`

	Vat_number string `xml:"vat_number,omitempty"`

	Present_legal_form_code int32 `xml:"present_legal_form_code,omitempty"`

	Present_legal_form_text string `xml:"present_legal_form_text,omitempty"`

	Present_legal_form_fine_code int32 `xml:"present_legal_form_fine_code,omitempty"`

	Present_legal_form_fine_text string `xml:"present_legal_form_fine_text,omitempty"`

	Number_of_non_executive_partners string `xml:"number_of_non_executive_partners,omitempty"`

	Publication_duty_code string `xml:"publication_duty_code,omitempty"`

	Publication_duty_text string `xml:"publication_duty_text,omitempty"`

	Capital *GCRCapitalArray `xml:"capital,omitempty"`

	Founded *GCRFoundedArray `xml:"founded,omitempty"`

	Discontinuance *GCRDiscontinuance `xml:"discontinuance,omitempty"`
}

type GCRChamberOfCommerce struct {
	Chamber_no string `xml:"chamber_no,omitempty"`

	Dossier_no string `xml:"dossier_no,omitempty"`

	Sub_dossier_no string `xml:"sub_dossier_no,omitempty"`

	Administering_chamber_no string `xml:"administering_chamber_no,omitempty"`

	Trade_register_location string `xml:"trade_register_location,omitempty"`
}

type GCRCapital struct {
	Is_current bool `xml:"is_current,omitempty"`

	Date_change time.Time `xml:"date_change,omitempty"`

	Social_capital string `xml:"social_capital,omitempty"`

	Capital_of_commandite_partnership string `xml:"capital_of_commandite_partnership,omitempty"`

	Issued_capital string `xml:"issued_capital,omitempty"`

	Paid_up_capital string `xml:"paid_up_capital,omitempty"`

	Monetary_amount_specified_in_corporate_deed string `xml:"monetary_amount_specified_in_corporate_deed,omitempty"`
}

type GCRCapitalArray struct {
	Item []*GCRCapital `xml:"item,omitempty"`
}

type GCRFounded struct {
	Date_from time.Time `xml:"date_from,omitempty"`

	Residence string `xml:"residence,omitempty"`

	Legal_form_code int32 `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Start_date time.Time `xml:"start_date,omitempty"`

	Founded_at time.Time `xml:"founded_at,omitempty"`
}

type GCRFoundedArray struct {
	Item []*GCRFounded `xml:"item,omitempty"`
}

type GCRDiscontinuance struct {
	Type_ string `xml:"type,omitempty"`

	Date string `xml:"date,omitempty"`

	Reason string `xml:"reason,omitempty"`
}

type GCRPersonnel struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Employees_class_code int32 `xml:"employees_class_code,omitempty"`

	Employees_class_text string `xml:"employees_class_text,omitempty"`

	Number_of_employees *GCRNumberOfEmployeesArray `xml:"number_of_employees,omitempty"`
}

type GCRNumberOfEmployees struct {
	Year int32 `xml:"year,omitempty"`

	Number_of_employees float32 `xml:"number_of_employees,omitempty"`
}

type GCRNumberOfEmployeesArray struct {
	Item []*GCRNumberOfEmployees `xml:"item,omitempty"`
}

type GCRSectorOfIndustry struct {
	Is_primary bool `xml:"is_primary,omitempty"`

	Sector_code string `xml:"sector_code,omitempty"`

	Sector_text string `xml:"sector_text,omitempty"`
}

type GCRSectorOfIndustryArray struct {
	Item []*GCRSectorOfIndustry `xml:"item,omitempty"`
}

type GCRAnnualFigures struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Company_annual_accounts *GCRCompanyAnnualAccountArray `xml:"company_annual_accounts,omitempty"`

	Bank_and_insurer_data *GCRBankAndInsurerDataArray `xml:"bank_and_insurer_data,omitempty"`
}

type GCRAnnualFiguresArray struct {
	Item []*GCRAnnualFigures `xml:"item,omitempty"`
}

type GCRCompanyAnnualAccount struct {
	Financial_year *GYear `xml:"financial_year,omitempty"`

	Month_end_financial_year int32 `xml:"month_end_financial_year,omitempty"`

	Day_end_financial_year int32 `xml:"day_end_financial_year,omitempty"`

	Length_financial_year_in_months int32 `xml:"length_financial_year_in_months,omitempty"`

	Reliable bool `xml:"reliable,omitempty"`

	Account_code string `xml:"account_code,omitempty"`

	Account_text string `xml:"account_text,omitempty"`

	Filing *GCRFiling `xml:"filing,omitempty"`

	Financial_year_graydon *GYear `xml:"financial_year_graydon,omitempty"`

	Currency_code string `xml:"currency_code,omitempty"`

	Scale_code string `xml:"scale_code,omitempty"`

	Scale_text string `xml:"scale_text,omitempty"`

	Date_of_drawing_up time.Time `xml:"date_of_drawing_up,omitempty"`

	Balance_sheet_code string `xml:"balance_sheet_code,omitempty"`

	Balance_sheet_text string `xml:"balance_sheet_text,omitempty"`

	Balance_sheet *GCRBalanceSheet `xml:"balance_sheet,omitempty"`

	Profit_and_loss_account *GCRItemArray `xml:"profit_and_loss_account,omitempty"`

	Supplement_to_annual_account *GCRItemArray `xml:"supplement_to_annual_account,omitempty"`

	Ratios *GCRRatioArray `xml:"ratios,omitempty"`
}

type GCRCompanyAnnualAccountArray struct {
	Item []*GCRCompanyAnnualAccount `xml:"item,omitempty"`
}

type GCRBankAndInsurerData struct {
	Financial_year *GYear `xml:"financial_year,omitempty"`

	Month_end_financial_year int32 `xml:"month_end_financial_year,omitempty"`

	Day_end_financial_year int32 `xml:"day_end_financial_year,omitempty"`

	Length_financial_year_in_months int32 `xml:"length_financial_year_in_months,omitempty"`

	Reliable bool `xml:"reliable,omitempty"`

	Account_code string `xml:"account_code,omitempty"`

	Account_text string `xml:"account_text,omitempty"`

	Filing *GCRFiling `xml:"filing,omitempty"`

	Financial_year_graydon *GYear `xml:"financial_year_graydon,omitempty"`

	Currency_code string `xml:"currency_code,omitempty"`

	Scale_code string `xml:"scale_code,omitempty"`

	Scale_text string `xml:"scale_text,omitempty"`

	Date_of_drawing_up time.Time `xml:"date_of_drawing_up,omitempty"`

	Balance_sheet_code string `xml:"balance_sheet_code,omitempty"`

	Balance_sheet_text string `xml:"balance_sheet_text,omitempty"`

	Bank_and_insurer_key_figures *GCRBankAndInsurerKeyFigures `xml:"bank_and_insurer_key_figures,omitempty"`

	Ratios *GCRRatioArray `xml:"ratios,omitempty"`
}

type GCRBankAndInsurerDataArray struct {
	Item []*GCRBankAndInsurerData `xml:"item,omitempty"`
}

type GCRBankAndInsurerKeyFigures struct {
	Balance_sheet_total string `xml:"balance_sheet_total,omitempty"`

	Equity_capital string `xml:"equity_capital,omitempty"`

	Group_equity string `xml:"group_equity,omitempty"`

	Liable_equity string `xml:"liable_equity,omitempty"`

	Technical_provisions string `xml:"technical_provisions,omitempty"`

	Entrusted_funds string `xml:"entrusted_funds,omitempty"`

	Debt_securities string `xml:"debt_securities,omitempty"`

	Credits string `xml:"credits,omitempty"`

	Attracted_bankers_funds string `xml:"attracted_bankers_funds,omitempty"`

	Issued_bankers_funds string `xml:"issued_bankers_funds,omitempty"`

	Interest string `xml:"interest,omitempty"`

	Premium_income string `xml:"premium_income,omitempty"`

	Other_income string `xml:"other_income,omitempty"`

	Total_benefits string `xml:"total_benefits,omitempty"`

	Yield_on_investments string `xml:"yield_on_investments,omitempty"`

	Operating_investments string `xml:"operating_investments,omitempty"`

	Holding_gain_of_accounts_receivable string `xml:"holding_gain_of_accounts_receivable,omitempty"`

	Result_before_taxes string `xml:"result_before_taxes,omitempty"`

	Group_profit string `xml:"group_profit,omitempty"`

	Net_profit string `xml:"net_profit,omitempty"`
}

type GCRHistory struct {
	Originated_from *GCROriginatedFromArray `xml:"originated_from,omitempty"`

	Continuation *GCRContinuationArray `xml:"continuation,omitempty"`

	Company_continued_under *GCRCompanyContinuedUnderArray `xml:"company_continued_under,omitempty"`
}

type GCROriginatedFrom struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Originated_from_code int32 `xml:"originated_from_code,omitempty"`

	Originated_from_text string `xml:"originated_from_text,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`
}

type GCROriginatedFromArray struct {
	Item []*GCROriginatedFrom `xml:"item,omitempty"`
}

type GCRContinuation struct {
	Date_from time.Time `xml:"date_from,omitempty"`

	Residence string `xml:"residence,omitempty"`

	Legal_form_code int32 `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`
}

type GCRContinuationArray struct {
	Item []*GCRContinuation `xml:"item,omitempty"`
}

type GCRCompanyContinuedUnder struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Continued_under_code int32 `xml:"continued_under_code,omitempty"`

	Continued_under_text string `xml:"continued_under_text,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`
}

type GCRCompanyContinuedUnderArray struct {
	Item []*GCRCompanyContinuedUnder `xml:"item,omitempty"`
}

type GCRCreditAdviceData struct {
	Including_personal_judgement bool `xml:"including_personal_judgement,omitempty"`

	Credit_advice *GCRCreditAdvice `xml:"credit_advice,omitempty"`

	Credit_advice_mother *GCRCreditAdviceMother `xml:"credit_advice_mother,omitempty"`

	Credit_rating *GCRCreditRating `xml:"credit_rating,omitempty"`
}

type GCRCreditAdvice struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Credit_amount string `xml:"credit_amount,omitempty"`

	Credit_currency string `xml:"credit_currency,omitempty"`

	History_advised_limits *GCRHistoryAdvisedLimitArray `xml:"history_advised_limits,omitempty"`

	Credit_factors *GCRCreditFactorArray `xml:"credit_factors,omitempty"`
}

type GCRHistoryAdvisedLimit struct {
	Value string `xml:"value,omitempty"`

	Year_month *GYearMonth `xml:"year_month,omitempty"`
}

type GCRHistoryAdvisedLimitArray struct {
	Item []*GCRHistoryAdvisedLimit `xml:"item,omitempty"`
}

type GCRCreditFactor struct {
	Type_ string `xml:"type,omitempty"`

	Description string `xml:"description,omitempty"`

	Adjustment int32 `xml:"adjustment,omitempty"`
}

type GCRCreditFactorArray struct {
	Item []*GCRCreditFactor `xml:"item,omitempty"`
}

type GCRCreditAdviceMother struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Amount string `xml:"amount,omitempty"`
}

type GCRCreditRating struct {
	Y_rating string `xml:"y_rating,omitempty"`

	Pd_rating string `xml:"pd_rating,omitempty"`

	Pd_percentage float32 `xml:"pd_percentage,omitempty"`

	Erc_score string `xml:"erc_score,omitempty"`

	Credit_flag_code string `xml:"credit_flag_code,omitempty"`

	Credit_flag_text string `xml:"credit_flag_text,omitempty"`

	Ratings *GCRRatingArray `xml:"ratings,omitempty"`

	Y_rating_history *GCRRatingHistoryArray `xml:"y_rating_history,omitempty"`

	Pd_rating_history *GCRRatingHistoryArray `xml:"pd_rating_history,omitempty"`

	Pd_percentage_history *GCRPercentageHistoryArray `xml:"pd_percentage_history,omitempty"`
}

type GCRRating struct {
	Type_ string `xml:"type,omitempty"`

	Description string `xml:"description,omitempty"`

	Value string `xml:"value,omitempty"`
}

type GCRRatingArray struct {
	Item []*GCRRating `xml:"item,omitempty"`
}

type GCRRatingHistory struct {
	Value string `xml:"value,omitempty"`

	Year_month *GYearMonth `xml:"year_month,omitempty"`
}

type GCRRatingHistoryArray struct {
	Item []*GCRRatingHistory `xml:"item,omitempty"`
}

type GCRPercentageHistory struct {
	Value float32 `xml:"value,omitempty"`

	Year_month *GYearMonth `xml:"year_month,omitempty"`
}

type GCRPercentageHistoryArray struct {
	Item []*GCRPercentageHistory `xml:"item,omitempty"`
}

type GCRCalamity struct {
	Calamity_events *GCRCalamityEventArray `xml:"calamity_events,omitempty"`

	Other_calamities *GCRCalamityOtherArray `xml:"other_calamities,omitempty"`
}

type GCRCalamityEvent struct {
	Event_code int32 `xml:"event_code,omitempty"`

	Event_text string `xml:"event_text,omitempty"`

	Event_date time.Time `xml:"event_date,omitempty"`

	Expiration_date time.Time `xml:"expiration_date,omitempty"`

	Damage_amount string `xml:"damage_amount,omitempty"`

	Damage_currency string `xml:"damage_currency,omitempty"`

	Insurance_code int32 `xml:"insurance_code,omitempty"`

	Insurance_text string `xml:"insurance_text,omitempty"`

	Instigator_graydon_company_id int32 `xml:"instigator_graydon_company_id,omitempty"`
}

type GCRCalamityEventArray struct {
	Item []*GCRCalamityEvent `xml:"item,omitempty"`
}

type GCRCalamityOther struct {
	Calamity_code int32 `xml:"calamity_code,omitempty"`

	Calamity_text string `xml:"calamity_text,omitempty"`

	Expiration_date time.Time `xml:"expiration_date,omitempty"`
}

type GCRCalamityOtherArray struct {
	Item []*GCRCalamityOther `xml:"item,omitempty"`
}

type GCRBranchOffice struct {
	Branch_office_id string `xml:"branch_office_id,omitempty"`

	Business_name string `xml:"business_name,omitempty"`

	Trade_names *GCRTradeNameArray `xml:"trade_names,omitempty"`

	Addresses *GCRAddressArray `xml:"addresses,omitempty"`

	Telephone_numbers *GCRTelephoneNumberArray `xml:"telephone_numbers,omitempty"`

	Email_addresses *StringArray `xml:"email_addresses,omitempty"`

	Estabishment_date time.Time `xml:"estabishment_date,omitempty"`

	Discontinuance *GCRDiscontinuance `xml:"discontinuance,omitempty"`
}

type GCRBranchOfficeArray struct {
	Item []*GCRBranchOffice `xml:"item,omitempty"`
}

type GCRPaymentInformation struct {
	Value float32 `xml:"value,omitempty"`

	Year_month *GYearMonth `xml:"year_month,omitempty"`
}

type GCRPaymentInformationArray struct {
	Item []*GCRPaymentInformation `xml:"item,omitempty"`
}

type GCRFinancialDetails struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Turnovers *GCRFinancialDetailValueArray `xml:"turnovers,omitempty"`

	Results *GCRFinancialDetailValueArray `xml:"results,omitempty"`
}

type GCRFinancialDetailValue struct {
	Financial_code string `xml:"financial_code,omitempty"`

	Financial_text string `xml:"financial_text,omitempty"`

	Year int32 `xml:"year,omitempty"`

	Amount float64 `xml:"amount,omitempty"`

	Amount_currency string `xml:"amount_currency,omitempty"`

	Forecast float64 `xml:"forecast,omitempty"`
}

type GCRFinancialDetailValueArray struct {
	Item []*GCRFinancialDetailValue `xml:"item,omitempty"`
}

type GCRDeclarationOfLiability struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Is_current bool `xml:"is_current,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`
}

type GCRDeclarationOfLiabilityArray struct {
	Item []*GCRDeclarationOfLiability `xml:"item,omitempty"`
}

type GCRParticipation struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Is_current bool `xml:"is_current,omitempty"`

	Percentage_of_shares float32 `xml:"percentage_of_shares,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`
}

type GCRParticipationArray struct {
	Item []*GCRParticipation `xml:"item,omitempty"`
}

type GCRShareHolder struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Is_current bool `xml:"is_current,omitempty"`

	Percentage_of_shares float32 `xml:"percentage_of_shares,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`
}

type GCRShareHolderArray struct {
	Item []*GCRShareHolder `xml:"item,omitempty"`
}

type GCRImportExport struct {
	Import_countries *GCRCountryArray `xml:"import_countries,omitempty"`

	Export_countries *GCRCountryArray `xml:"export_countries,omitempty"`
}

type GCRCountry struct {
	Country_code int32 `xml:"country_code,omitempty"`

	Country_text string `xml:"country_text,omitempty"`
}

type GCRCountryArray struct {
	Item []*GCRCountry `xml:"item,omitempty"`
}

type GCRSpecialCompanyInformation struct {
	Special_company_code string `xml:"special_company_code,omitempty"`

	Special_company_text string `xml:"special_company_text,omitempty"`

	Texts *GCRTextSequenceArray `xml:"texts,omitempty"`
}

type GCRSpecialCompanyInformationArray struct {
	Item []*GCRSpecialCompanyInformation `xml:"item,omitempty"`
}

type GCRTextSequence struct {
	Sequence_number int32 `xml:"sequence_number,omitempty"`

	Text string `xml:"text,omitempty"`
}

type GCRTextSequenceArray struct {
	Item []*GCRTextSequence `xml:"item,omitempty"`
}

type GCRLiaison struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Liaison_code string `xml:"liaison_code,omitempty"`

	Liaison_text string `xml:"liaison_text,omitempty"`
}

type GCRLiaisonArray struct {
	Item []*GCRLiaison `xml:"item,omitempty"`
}

type GCRCompanyManagement struct {
	Management *GCRManagementArray `xml:"management,omitempty"`

	First_directors *GCRFirstDirectorArray `xml:"first_directors,omitempty"`
}

type GCRManagement struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Is_current bool `xml:"is_current,omitempty"`

	Job_titles *GCRJobTitleArray `xml:"job_titles,omitempty"`
}

type GCRManagementArray struct {
	Item []*GCRManagement `xml:"item,omitempty"`
}

type GCRJobTitle struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Is_current bool `xml:"is_current,omitempty"`

	Job_title_code int32 `xml:"job_title_code,omitempty"`

	Job_title_text string `xml:"job_title_text,omitempty"`

	Authority_code int32 `xml:"authority_code,omitempty"`

	Authority_text string `xml:"authority_text,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`
}

type GCRJobTitleArray struct {
	Item []*GCRJobTitle `xml:"item,omitempty"`
}

type GCRFirstDirector struct {
	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Is_current bool `xml:"is_current,omitempty"`

	Job_title_code int32 `xml:"job_title_code,omitempty"`

	Job_title_text string `xml:"job_title_text,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Date_until time.Time `xml:"date_until,omitempty"`
}

type GCRFirstDirectorArray struct {
	Item []*GCRFirstDirector `xml:"item,omitempty"`
}

type GCRCompanyRelations struct {
	Companies *GCRCompanyArray `xml:"companies,omitempty"`

	Persons *GCRPersonArray `xml:"persons,omitempty"`
}

type GCRCompanyRelationsSimple struct {
	Companies *GCRCompanySimpleArray `xml:"companies,omitempty"`

	Persons *GCRPersonSimpleArray `xml:"persons,omitempty"`
}

type GCRCompanySimple struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Contact_details *GCRContactDetails `xml:"contact_details,omitempty"`

	Official_data *GCROfficialData `xml:"official_data,omitempty"`

	History *GCRHistory `xml:"history,omitempty"`

	Sectors_of_industry *GCRSectorOfIndustryArray `xml:"sectors_of_industry,omitempty"`
}

type GCRCompanySimpleArray struct {
	Item []*GCRCompanySimple `xml:"item,omitempty"`
}

type GCRCompany struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Contact_details *GCRContactDetails `xml:"contact_details,omitempty"`

	Official_data *GCROfficialData `xml:"official_data,omitempty"`

	History *GCRHistory `xml:"history,omitempty"`

	Sectors_of_industry *GCRSectorOfIndustryArray `xml:"sectors_of_industry,omitempty"`

	Positions_in_other_companies *GCRJobTitleArray `xml:"positions_in_other_companies,omitempty"`

	Annual_figures *GCRAnnualFiguresArray `xml:"annual_figures,omitempty"`

	Financial_details *GCRFinancialDetails `xml:"financial_details,omitempty"`

	Financial_calamities *GCRFinancialCalamityArray `xml:"financial_calamities,omitempty"`
}

type GCRCompanyArray struct {
	Item []*GCRCompany `xml:"item,omitempty"`
}

type GCRLiaisonCompany struct {
	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Contact_details *GCRLiaisonContactDetails `xml:"contact_details,omitempty"`

	Official_data *GCRLiaisonOfficialData `xml:"official_data,omitempty"`

	Financial_calamities *GCRFinancialCalamityArray `xml:"financial_calamities,omitempty"`
}

type GCRLiaisonCompanyArray struct {
	Item []*GCRLiaisonCompany `xml:"item,omitempty"`
}

type GCRLiaisonContactDetails struct {
	Company_name string `xml:"company_name,omitempty"`

	Addresses *GCRAddressArray `xml:"addresses,omitempty"`

	Telephone_numbers *GCRTelephoneNumberArray `xml:"telephone_numbers,omitempty"`
}

type GCRLiaisonOfficialData struct {
	Chamber_of_commerce *GCRChamberOfCommerce `xml:"chamber_of_commerce,omitempty"`

	Present_legal_form_code int32 `xml:"present_legal_form_code,omitempty"`

	Present_legal_form_text string `xml:"present_legal_form_text,omitempty"`
}

type GCRFiling struct {
	Filing_code string `xml:"filing_code,omitempty"`

	Filing_text string `xml:"filing_text,omitempty"`

	Filing_date time.Time `xml:"filing_date,omitempty"`

	Chamber_of_commerce *GCRChamberOfCommerce `xml:"chamber_of_commerce,omitempty"`
}

type GCRBalanceSheet struct {
	Items *GCRItemArray `xml:"items,omitempty"`

	Assets *GCRAssets `xml:"assets,omitempty"`

	Liabilities *GCRLiabilities `xml:"liabilities,omitempty"`
}

type GCRItem struct {
	Main_code string `xml:"main_code,omitempty"`

	Code string `xml:"code,omitempty"`

	Sub_code string `xml:"sub_code,omitempty"`

	Description string `xml:"description,omitempty"`

	Value float64 `xml:"value,omitempty"`

	Free_text string `xml:"free_text,omitempty"`
}

type GCRItemArray struct {
	Item []*GCRItem `xml:"item,omitempty"`
}

type GCRAssets struct {
	Fixed_assets_total *GCRItem `xml:"fixed_assets_total,omitempty"`

	Fixed_assets_items *GCRItemArray `xml:"fixed_assets_items,omitempty"`

	Current_assets_total *GCRItem `xml:"current_assets_total,omitempty"`

	Current_assets_items *GCRItemArray `xml:"current_assets_items,omitempty"`

	Balance_sheet_total float64 `xml:"balance_sheet_total,omitempty"`
}

type GCRLiabilities struct {
	Liabilities_items *GCRItemArray `xml:"liabilities_items,omitempty"`

	Balance_sheet_total float64 `xml:"balance_sheet_total,omitempty"`
}

type GCRRatio struct {
	Ratio_code string `xml:"ratio_code,omitempty"`

	Ratio_text string `xml:"ratio_text,omitempty"`

	Value float64 `xml:"value,omitempty"`
}

type GCRRatioArray struct {
	Item []*GCRRatio `xml:"item,omitempty"`
}

type GCRPerson struct {
	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Personal_details *GCRPersonalDetails `xml:"personal_details,omitempty"`

	Telephone_number *GCRTelephoneNumber `xml:"telephone_number,omitempty"`

	Address *GCRAddress `xml:"address,omitempty"`

	Birth_information *GCRBirthInformation `xml:"birth_information,omitempty"`

	Job_titles *GCRJobTitleArray `xml:"job_titles,omitempty"`

	Financial_calamities *GCRFinancialCalamityArray `xml:"financial_calamities,omitempty"`
}

type GCRPersonArray struct {
	Item []*GCRPerson `xml:"item,omitempty"`
}

type GCRPersonSimple struct {
	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Personal_details *GCRPersonalDetails `xml:"personal_details,omitempty"`

	Address *GCRAddress `xml:"address,omitempty"`

	Birth_information *GCRBirthInformation `xml:"birth_information,omitempty"`
}

type GCRPersonSimpleArray struct {
	Item []*GCRPersonSimple `xml:"item,omitempty"`
}

type GCRPersonalDetails struct {
	Gender_code string `xml:"gender_code,omitempty"`

	Gender_text string `xml:"gender_text,omitempty"`

	Titles *GCRTitleArray `xml:"titles,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefixes string `xml:"prefixes,omitempty"`

	Name string `xml:"name,omitempty"`
}

type GCRTitle struct {
	Title_code int32 `xml:"title_code,omitempty"`

	Title_text string `xml:"title_text,omitempty"`
}

type GCRTitleArray struct {
	Item []*GCRTitle `xml:"item,omitempty"`
}

type GCRBirthInformation struct {
	Date time.Time `xml:"date,omitempty"`

	Residence string `xml:"residence,omitempty"`

	Country_code int32 `xml:"country_code,omitempty"`

	Country_text string `xml:"country_text,omitempty"`
}

type GCRFinancialCalamity struct {
	Bankruptcy *GCRBankruptcy `xml:"bankruptcy,omitempty"`

	Moratorium *GCRMoratorium `xml:"moratorium,omitempty"`

	Debt_restructuring_private_persons *GCRDebtRestructuringPrivatePersons `xml:"debt_restructuring_private_persons,omitempty"`
}

type GCRFinancialCalamityArray struct {
	Item []*GCRFinancialCalamity `xml:"item,omitempty"`
}

type GCRBankruptcy struct {
	Official_registration_number int32 `xml:"official_registration_number,omitempty"`

	Currently_active bool `xml:"currently_active,omitempty"`

	Date_of_verdict time.Time `xml:"date_of_verdict,omitempty"`

	Definite_discontinuance_code int32 `xml:"definite_discontinuance_code,omitempty"`

	Definite_discontinuance_text string `xml:"definite_discontinuance_text,omitempty"`

	Definite_discontinuance_date time.Time `xml:"definite_discontinuance_date,omitempty"`

	Curators *GCRCuratorArray `xml:"curators,omitempty"`
}

type GCRCurator struct {
	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Date_when_active_in_job_title time.Time `xml:"date_when_active_in_job_title,omitempty"`
}

type GCRCuratorArray struct {
	Item []*GCRCurator `xml:"item,omitempty"`
}

type GCRMoratorium struct {
	Official_registration_number int32 `xml:"official_registration_number,omitempty"`

	Currently_active bool `xml:"currently_active,omitempty"`

	Date_from time.Time `xml:"date_from,omitempty"`

	Provisional_verdict_date time.Time `xml:"provisional_verdict_date,omitempty"`

	Definite_verdict_date time.Time `xml:"definite_verdict_date,omitempty"`

	Moratorium_term_in_months int32 `xml:"moratorium_term_in_months,omitempty"`

	Prolongation_date time.Time `xml:"prolongation_date,omitempty"`

	Prolongation_in_months int32 `xml:"prolongation_in_months,omitempty"`

	Vote_code int32 `xml:"vote_code,omitempty"`

	Vote_text string `xml:"vote_text,omitempty"`

	Vote_date time.Time `xml:"vote_date,omitempty"`

	Definite_discontinuance_code int32 `xml:"definite_discontinuance_code,omitempty"`

	Definite_discontinuance_text string `xml:"definite_discontinuance_text,omitempty"`

	Definite_discontinuance_date time.Time `xml:"definite_discontinuance_date,omitempty"`

	Receivers *GCRReceiverArray `xml:"receivers,omitempty"`
}

type GCRReceiver struct {
	Graydon_person_id int32 `xml:"graydon_person_id,omitempty"`

	Date_when_active_in_job_title time.Time `xml:"date_when_active_in_job_title,omitempty"`
}

type GCRReceiverArray struct {
	Item []*GCRReceiver `xml:"item,omitempty"`
}

type GCRDebtRestructuringPrivatePersons struct {
	Official_registration_number int32 `xml:"official_registration_number,omitempty"`

	Currently_active bool `xml:"currently_active,omitempty"`

	Provisional_verdict_date time.Time `xml:"provisional_verdict_date,omitempty"`

	Definite_verdict_date time.Time `xml:"definite_verdict_date,omitempty"`

	Debt_restructuring_term_in_months int32 `xml:"debt_restructuring_term_in_months,omitempty"`

	Provisional_discontinuance_code int32 `xml:"provisional_discontinuance_code,omitempty"`

	Provisional_discontinuance_text string `xml:"provisional_discontinuance_text,omitempty"`

	Provisional_discontinuance_date time.Time `xml:"provisional_discontinuance_date,omitempty"`

	Definite_discontinuance_code int32 `xml:"definite_discontinuance_code,omitempty"`

	Definite_discontinuance_text string `xml:"definite_discontinuance_text,omitempty"`

	Definite_discontinuance_date time.Time `xml:"definite_discontinuance_date,omitempty"`

	Receivers *GCRReceiverArray `xml:"receivers,omitempty"`
}

type GCRAlarm struct {
	Alarm_code int32 `xml:"alarm_code,omitempty"`

	Alarm_text string `xml:"alarm_text,omitempty"`
}

type InsolvencyCase struct {
	Insolvency_id string `xml:"insolvency_id,omitempty"`

	Magistrate string `xml:"magistrate,omitempty"`

	Previous_insolvency_id string `xml:"previous_insolvency_id,omitempty"`

	Previous_magistrate string `xml:"previous_magistrate,omitempty"`

	Is_pre_hgk_published bool `xml:"is_pre_hgk_published,omitempty"`

	Pre_hgk_insolvency_id string `xml:"pre_hgk_insolvency_id,omitempty"`

	Dept_restructuring_number string `xml:"dept_restructuring_number,omitempty"`

	Removal_date time.Time `xml:"removal_date,omitempty"`

	Legal_subject *InsolvencyLegalSubject `xml:"legal_subject,omitempty"`

	Treating_authority *InsolvencyTreatingAuthority `xml:"treating_authority,omitempty"`

	Curators *InsolvencyCurators `xml:"curators,omitempty"`

	Receivers *InsolvencyReceivers `xml:"receivers,omitempty"`

	Events *InsolvencyPublications `xml:"events,omitempty"`

	Reports *InsolvencyReportArray `xml:"reports,omitempty"`

	Affected_companies *InsolvencyAffectedCompanyArray `xml:"affected_companies,omitempty"`
}

type InsolvencyTreatingAuthority struct {
	Code int32 `xml:"code,omitempty"`

	Name string `xml:"name,omitempty"`

	Establishment_code int32 `xml:"establishment_code,omitempty"`

	Establishment_name string `xml:"establishment_name,omitempty"`
}

type InsolvencyLegalSubject struct {
	Natural_person *InsolvencyConcernedPerson `xml:"natural_person,omitempty"`

	Legal_entity *InsolvencyAffectedCompany `xml:"legal_entity,omitempty"`
}

type InsolvencyConcernedPerson struct {
	First_name string `xml:"first_name,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Addresses *InsolvencyAddresses `xml:"addresses,omitempty"`

	Birth_date time.Time `xml:"birth_date,omitempty"`

	Birth_city string `xml:"birth_city,omitempty"`

	Birth_country string `xml:"birth_country,omitempty"`

	Deceased_date time.Time `xml:"deceased_date,omitempty"`
}

type InsolvencyLegalPerformingPerson struct {
	Titles string `xml:"titles,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Contact_info *InsolvencyAddress `xml:"contact_info,omitempty"`

	Start_date time.Time `xml:"start_date,omitempty"`

	End_date time.Time `xml:"end_date,omitempty"`
}

type InsolvencyLegalPerformingPersonArray struct {
	Item []*InsolvencyLegalPerformingPerson `xml:"item,omitempty"`
}

type InsolvencyAddresses struct {
	Hidden bool `xml:"hidden,omitempty"`

	Address *InsolvencyAddressArray `xml:"address,omitempty"`
}

type InsolvencyPublications struct {
	Is_legacy bool `xml:"is_legacy,omitempty"`

	Event *InsolvencyPublicationArray `xml:"event,omitempty"`
}

type InsolvencyAddress struct {
	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	Addition string `xml:"addition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Start_date time.Time `xml:"start_date,omitempty"`
}

type InsolvencyAddressArray struct {
	Item []*InsolvencyAddress `xml:"item,omitempty"`
}

type InsolvencyPublication struct {
	Publication_id string `xml:"publication_id,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Description string `xml:"description,omitempty"`

	Code string `xml:"code,omitempty"`

	Authority_code int32 `xml:"authority_code,omitempty"`

	Court_session *InsolvencyCourtSession `xml:"court_session,omitempty"`
}

type InsolvencyPublicationArray struct {
	Item []*InsolvencyPublication `xml:"item,omitempty"`
}

type InsolvencyCourtSession struct {
	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	Addition string `xml:"addition,omitempty"`

	City string `xml:"city,omitempty"`
}

type InsolvencyCurators struct {
	Curator *InsolvencyLegalPerformingPersonArray `xml:"curator,omitempty"`
}

type InsolvencyReceivers struct {
	Receiver *InsolvencyLegalPerformingPersonArray `xml:"receiver,omitempty"`
}

type InsolvencyReport struct {
	Report_id string `xml:"report_id,omitempty"`

	Title string `xml:"title,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Is_final bool `xml:"is_final,omitempty"`
}

type InsolvencyReportArray struct {
	Item []*InsolvencyReport `xml:"item,omitempty"`
}

type InsolvencyAffectedCompany struct {
	Trade_name string `xml:"trade_name,omitempty"`

	Coc_number string `xml:"coc_number,omitempty"`

	Coc_registered_city string `xml:"coc_registered_city,omitempty"`

	Addresses *InsolvencyAddresses `xml:"addresses,omitempty"`

	Is_former bool `xml:"is_former,omitempty"`
}

type InsolvencyAffectedCompanyArray struct {
	Item []*InsolvencyAffectedCompany `xml:"item,omitempty"`
}

type InsolvencyPublicationList struct {
	Date time.Time `xml:"date,omitempty"`

	Publications *StringArray `xml:"publications,omitempty"`
}

type InternationalAddressSearchV2Result struct {
	Validation_status string `xml:"validation_status,omitempty"`

	Result *InternationalV2Array `xml:"result,omitempty"`

	Country_iso3 string `xml:"country_iso3,omitempty"`
}

type InternationalV2 struct {
	Address *InternationalFormattedAddress `xml:"address,omitempty"`

	Element_match_status string `xml:"element_match_status,omitempty"`

	Element_result_status string `xml:"element_result_status,omitempty"`
}

type InternationalV2Array struct {
	Item []*InternationalV2 `xml:"item,omitempty"`
}

type International struct {
	Matchrate float32 `xml:"matchrate,omitempty"`

	Street string `xml:"street,omitempty"`

	Housenr string `xml:"housenr,omitempty"`

	Pobox string `xml:"pobox,omitempty"`

	Locality string `xml:"locality,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Province string `xml:"province,omitempty"`

	Country string `xml:"country,omitempty"`
}

type InternationalArray struct {
	Item []*International `xml:"item,omitempty"`
}

type InternationalPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *InternationalArray `xml:"results,omitempty"`
}

type InternationalFormattedAddress struct {
	Matchrate float32 `xml:"matchrate,omitempty"`

	Street string `xml:"street,omitempty"`

	Housenr string `xml:"housenr,omitempty"`

	Pobox string `xml:"pobox,omitempty"`

	Locality string `xml:"locality,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Province string `xml:"province,omitempty"`

	Country string `xml:"country,omitempty"`

	Countryspecific_locality string `xml:"countryspecific_locality,omitempty"`

	Delivery_address string `xml:"delivery_address,omitempty"`

	Formatted_address string `xml:"formatted_address,omitempty"`

	Organization string `xml:"organization,omitempty"`

	Department string `xml:"department,omitempty"`

	Contact string `xml:"contact,omitempty"`

	Building string `xml:"building,omitempty"`
}

type InternationalFormattedAddressArray struct {
	Item []*InternationalFormattedAddress `xml:"item,omitempty"`
}

type InternationalFormattedAddressPagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *InternationalFormattedAddressArray `xml:"results,omitempty"`
}

type KadasterCoordinates struct {
	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`

	Within_structure bool `xml:"within_structure,omitempty"`
}

type KoopsommenOverzicht struct {
	Eind_datum time.Time `xml:"eind_datum,omitempty"`

	Straatnaam string `xml:"straatnaam,omitempty"`

	Huisnr_van int32 `xml:"huisnr_van,omitempty"`

	Huisnr_tm int32 `xml:"huisnr_tm,omitempty"`

	Reeksindicatie string `xml:"reeksindicatie,omitempty"`

	Gemeente string `xml:"gemeente,omitempty"`

	Koopsommen *KoopsomArray `xml:"koopsommen,omitempty"`
}

type KoopsommenOverzichtV2 struct {
	Document []byte `xml:"document,omitempty"`

	Overzicht *KoopsommenOverzicht `xml:"overzicht,omitempty"`
}

type Koopsom struct {
	Koopdatum time.Time `xml:"koopdatum,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Bedrag int32 `xml:"bedrag,omitempty"`

	Oppervlakte int32 `xml:"oppervlakte,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Meer_onroerende_goederen bool `xml:"meer_onroerende_goederen,omitempty"`
}

type KoopsomArray struct {
	Item []*Koopsom `xml:"item,omitempty"`
}

type Base64BinaryArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ base64BinaryArray"`

	Item [][]byte `xml:"item,omitempty"`
}

type UittrekselKadastraleKaart struct {
	Kaart []byte `xml:"kaart,omitempty"`

	Kadastrale_aanduidingen *KadastraleAanduidingArray `xml:"kadastrale_aanduidingen,omitempty"`

	Afbeeldingen *Base64BinaryArray `xml:"afbeeldingen,omitempty"`
}

type KadasterUittrekselKadastraleKaartV2 struct {
	Kaart []byte `xml:"kaart,omitempty"`

	Afbeeldingen *Base64BinaryArray `xml:"afbeeldingen,omitempty"`
}

type KadasterKadastraleKaart struct {
	Kaart []byte `xml:"kaart,omitempty"`

	Schaal int32 `xml:"schaal,omitempty"`

	Afbeeldingen *Base64BinaryArray `xml:"afbeeldingen,omitempty"`
}

type KadasterUittrekselKadastraleKaartResultaat struct {
	Overzicht *KadasterOverzicht `xml:"overzicht,omitempty"`

	Kadastrale_kaart *KadasterUittrekselKadastraleKaartV2 `xml:"kadastrale_kaart,omitempty"`
}

type KadasterUittrekselKadastraleKaartResultaatV2 struct {
	Overzicht *KadasterOverzichtV2 `xml:"overzicht,omitempty"`

	Kadastrale_kaart *KadasterUittrekselKadastraleKaartV2 `xml:"kadastrale_kaart,omitempty"`
}

type KadasterKadastraleKaartResultaat struct {
	Overzicht *KadasterOverzicht `xml:"overzicht,omitempty"`

	Kadastrale_kaart *KadasterKadastraleKaart `xml:"kadastrale_kaart,omitempty"`
}

type KadasterKadastraleKaartResultaatV2 struct {
	Overzicht *DocumentOverzicht `xml:"overzicht,omitempty"`

	Kadastrale_kaart *KadasterKadastraleKaart `xml:"kadastrale_kaart,omitempty"`
}

type KadastraleAanduiding struct {
	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`
}

type KadastraleAanduidingArray struct {
	Item []*KadastraleAanduiding `xml:"item,omitempty"`
}

type DocumentOverzicht struct {
	Onroerende_zaken *OverzichtOnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`
}

type Overzicht struct {
	Onroerende_zaken *OverzichtOnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`

	Rechten *OverzichtRechtArray `xml:"rechten,omitempty"`

	Personen *OverzichtPersoonArray `xml:"personen,omitempty"`
}

type OverzichtV2 struct {
	Onroerende_zaken *OverzichtOnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`

	Rechten *OverzichtRechtArray `xml:"rechten,omitempty"`

	Personen *OverzichtPersoonV2Array `xml:"personen,omitempty"`
}

type KadasterOverzicht struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzicht"`

	Onroerende_zaken *KadasterOverzichtOnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`

	Rechten *KadasterOverzichtRechtArray `xml:"rechten,omitempty"`

	Personen *KadasterOverzichtPersoonArray `xml:"personen,omitempty"`
}

type KadasterOverzichtV2 struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtV2"`

	Onroerende_zaken *KadasterOverzichtOnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`

	Rechten *KadasterOverzichtRechtArray `xml:"rechten,omitempty"`

	Personen *KadasterOverzichtPersoonV2Array `xml:"personen,omitempty"`
}

type OverzichtOnroerendeZaak struct {
	Object_id string `xml:"object_id,omitempty"`

	Kadastrale_aanduiding *KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Locatie *LocatieBinnenland `xml:"locatie,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`
}

type OverzichtOnroerendeZaakArray struct {
	Item []*OverzichtOnroerendeZaak `xml:"item,omitempty"`
}

type KadasterOverzichtOnroerendeZaak struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtOnroerendeZaak"`

	Object_id string `xml:"object_id,omitempty"`

	Kadastrale_aanduiding *KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Locatie *KadasterLocatieBinnenland `xml:"locatie,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type KadasterOverzichtOnroerendeZaakArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtOnroerendeZaakArray"`

	Item []*KadasterOverzichtOnroerendeZaak `xml:"item,omitempty"`
}

type KadasterBeknoptOnroerendeZaak struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterBeknoptOnroerendeZaak"`

	Object_id string `xml:"object_id,omitempty"`

	Kadastrale_aanduiding *KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Voorlopige_aantekening bool `xml:"voorlopige_aantekening,omitempty"`

	Meerdere_splitsingen bool `xml:"meerdere_splitsingen,omitempty"`

	Gerechtigde *KadasterPersoon `xml:"gerechtigde,omitempty"`

	Onverwerkte_documenten *KadasterHypotheekStukLijst `xml:"onverwerkte_documenten,omitempty"`

	Grootte float64 `xml:"grootte,omitempty"`

	Vastgestelde_grootte bool `xml:"vastgestelde_grootte,omitempty"`
}

type KadasterBeknoptOnroerendeZaakArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterBeknoptOnroerendeZaakArray"`

	Item []*KadasterBeknoptOnroerendeZaak `xml:"item,omitempty"`
}

type OverzichtRecht struct {
	Object_id string `xml:"object_id,omitempty"`

	Gerechtigde_id string `xml:"gerechtigde_id,omitempty"`
}

type OverzichtRechtArray struct {
	Item []*OverzichtRecht `xml:"item,omitempty"`
}

type KadasterOverzichtRecht struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtRecht"`

	Object_id string `xml:"object_id,omitempty"`

	Gerechtigde_akr string `xml:"gerechtigde_akr,omitempty"`

	Gerechtigde_kadastrale_identificatie string `xml:"gerechtigde_kadastrale_identificatie,omitempty"`
}

type KadasterOverzichtRechtArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtRechtArray"`

	Item []*KadasterOverzichtRecht `xml:"item,omitempty"`
}

type OverzichtPersoon struct {
	Id string `xml:"id,omitempty"`

	Natuurlijk_persoon *OverzichtNatuurlijkPersoon `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *OverzichtNietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`
}

type OverzichtPersoonArray struct {
	Item []*OverzichtPersoon `xml:"item,omitempty"`
}

type OverzichtPersoonV2 struct {
	Id string `xml:"id,omitempty"`

	Natuurlijk_persoon *OverzichtNatuurlijkPersoon `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *OverzichtNietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`

	Melding string `xml:"melding,omitempty"`
}

type OverzichtPersoonV2Array struct {
	Item []*OverzichtPersoonV2 `xml:"item,omitempty"`
}

type KadasterOverzichtPersoon struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtPersoon"`

	Akr_subject_nr string `xml:"akr_subject_nr,omitempty"`

	Kadastrale_identificatie string `xml:"kadastrale_identificatie,omitempty"`

	Natuurlijk_persoon *KadasterOverzichtNatuurlijkPersoon `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *KadasterOverzichtNietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`
}

type KadasterOverzichtPersoonArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtPersoonArray"`

	Item []*KadasterOverzichtPersoon `xml:"item,omitempty"`
}

type KadasterOverzichtPersoonV2 struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtPersoonV2"`

	Akr_subject_nr string `xml:"akr_subject_nr,omitempty"`

	Kadastrale_identificatie string `xml:"kadastrale_identificatie,omitempty"`

	Natuurlijk_persoon *KadasterOverzichtNatuurlijkPersoon `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *KadasterOverzichtNietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`

	Melding string `xml:"melding,omitempty"`
}

type KadasterOverzichtPersoonV2Array struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtPersoonV2Array"`

	Item []*KadasterOverzichtPersoonV2 `xml:"item,omitempty"`
}

type OverzichtNatuurlijkPersoon struct {
	Geslachtsnaam string `xml:"geslachtsnaam,omitempty"`

	Voorvoegsel_geslachtsnaam string `xml:"voorvoegsel_geslachtsnaam,omitempty"`

	Voornamen string `xml:"voornamen,omitempty"`
}

type KadasterOverzichtNatuurlijkPersoon struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtNatuurlijkPersoon"`

	Bsn string `xml:"bsn,omitempty"`

	Geslachtsnaam string `xml:"geslachtsnaam,omitempty"`

	Voorvoegsel_geslachtsnaam string `xml:"voorvoegsel_geslachtsnaam,omitempty"`

	Voornamen string `xml:"voornamen,omitempty"`

	Geslacht string `xml:"geslacht,omitempty"`
}

type OverzichtNietNatuurlijkPersoon struct {
	Naam string `xml:"naam,omitempty"`

	Code string `xml:"code,omitempty"`
}

type KadasterOverzichtNietNatuurlijkPersoon struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterOverzichtNietNatuurlijkPersoon"`

	Naam string `xml:"naam,omitempty"`

	Rechtsvorm string `xml:"rechtsvorm,omitempty"`

	Bedrijven_nummer string `xml:"bedrijven_nummer,omitempty"`
}

type BerichtObjectDocumentResultaat struct {
	Overzicht *DocumentOverzicht `xml:"overzicht,omitempty"`

	Eigendomsbericht *BerichtObjectDocument `xml:"eigendomsbericht,omitempty"`
}

type BerichtObjectDocument struct {
	Datum_bijgewerkt time.Time `xml:"datum_bijgewerkt,omitempty"`

	Object_id string `xml:"object_id,omitempty"`

	Kadastrale_aanduiding *KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Locaties *LocatieBinnenlandLijst `xml:"locaties,omitempty"`

	Document []byte `xml:"document,omitempty"`

	Afbeeldingen *Base64BinaryArray `xml:"afbeeldingen,omitempty"`
}

type KadasterBerichtObjectDocument struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterBerichtObjectDocument"`

	Document []byte `xml:"document,omitempty"`

	Afbeeldingen *Base64BinaryArray `xml:"afbeeldingen,omitempty"`
}

type BerichtObjectResultaat struct {
	Overzicht *Overzicht `xml:"overzicht,omitempty"`

	Eigendomsbericht *BerichtObject `xml:"eigendomsbericht,omitempty"`
}

type BerichtObjectResultaatV2 struct {
	Overzicht *OverzichtV2 `xml:"overzicht,omitempty"`

	Eigendomsbericht *BerichtObjectV2 `xml:"eigendomsbericht,omitempty"`
}

type BerichtObject struct {
	Datum_bijgewerkt time.Time `xml:"datum_bijgewerkt,omitempty"`

	Onroerende_zaak *OnroerendeZaak `xml:"onroerende_zaak,omitempty"`

	Rechten_lijst *RechtenLijst `xml:"rechten_lijst,omitempty"`

	Personen *PersoonArray `xml:"personen,omitempty"`

	Stukken *StukArray `xml:"stukken,omitempty"`

	Gemeentestukken *GemeenteStukArray `xml:"gemeentestukken,omitempty"`

	Document []byte `xml:"document,omitempty"`
}

type BerichtObjectV2 struct {
	Datum_bijgewerkt time.Time `xml:"datum_bijgewerkt,omitempty"`

	Onroerende_zaak *OnroerendeZaak `xml:"onroerende_zaak,omitempty"`

	Rechten_lijst *RechtenLijst `xml:"rechten_lijst,omitempty"`

	Personen *PersoonV2Array `xml:"personen,omitempty"`

	Stukken *StukArray `xml:"stukken,omitempty"`

	Gemeentestukken *GemeenteStukArray `xml:"gemeentestukken,omitempty"`

	Document []byte `xml:"document,omitempty"`
}

type OnroerendeZaak struct {
	Object_id string `xml:"object_id,omitempty"`

	Kadastrale_aanduiding *KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Locaties *LocatieBinnenlandLijst `xml:"locaties,omitempty"`

	Datum_ontstaan time.Time `xml:"datum_ontstaan,omitempty"`

	Onroerende_zaak_historie *KadastraleAanduidingLijst `xml:"onroerende_zaak_historie,omitempty"`

	Perceel_details *PerceelDetails `xml:"perceel_details,omitempty"`

	Koopsom string `xml:"koopsom,omitempty"`

	Koopjaar string `xml:"koopjaar,omitempty"`

	Indicatie_meer_onroerende_zaken bool `xml:"indicatie_meer_onroerende_zaken,omitempty"`

	Landinrichtingsrente *Landinrichtingsrente `xml:"landinrichtingsrente,omitempty"`

	Indicatie_voorgenomen_splitsing bool `xml:"indicatie_voorgenomen_splitsing,omitempty"`

	Indicatie_voorlopige_aantekening bool `xml:"indicatie_voorlopige_aantekening,omitempty"`

	Niet_volledig_verwerkt_stukken_lijst *StukReferentiesLijst `xml:"niet_volledig_verwerkt_stukken_lijst,omitempty"`

	Mandeligheid *Mandeligheid `xml:"mandeligheid,omitempty"`

	Aantekeningen_kadastraal_object *AantekeningenKadastraalObject `xml:"aantekeningen_kadastraal_object,omitempty"`

	Pr_beperkingen_kadastraal_object *PRBeperkingenKadastraalObject `xml:"pr_beperkingen_kadastraal_object,omitempty"`
}

type KadastraleAanduidingLijst struct {
	Kadastrale_aanduidingen *KadastraleAanduidingArray `xml:"kadastrale_aanduidingen,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type PerceelDetails struct {
	Perceel *KadastraalPerceel `xml:"perceel,omitempty"`

	Appartementsrecht *VerenigingVanEigenaren `xml:"appartementsrecht,omitempty"`
}

type KadastraalPerceel struct {
	X_coordinaat string `xml:"x_coordinaat,omitempty"`

	Y_coordinaat string `xml:"y_coordinaat,omitempty"`

	Grootte_perceel string `xml:"grootte_perceel,omitempty"`

	Indicatie_grootte_geschat bool `xml:"indicatie_grootte_geschat,omitempty"`
}

type VerenigingVanEigenaren struct {
	Naam string `xml:"naam,omitempty"`

	Bestuurder_id string `xml:"bestuurder_id,omitempty"`
}

type Landinrichtingsrente struct {
	Bedrag string `xml:"bedrag,omitempty"`

	Eindjaar string `xml:"eindjaar,omitempty"`
}

type Mandeligheid struct {
	Mandelige_percelen *KadastraleAanduidingLijst `xml:"mandelige_percelen,omitempty"`

	Hoofd_percelen *KadastraleAanduidingLijst `xml:"hoofd_percelen,omitempty"`
}

type AantekeningenKadastraalObject struct {
	Aantekeningen_lijst *AantekeningKadastraalObjectLijst `xml:"aantekeningen_lijst,omitempty"`

	Melding string `xml:"melding,omitempty"`
}

type AantekeningKadastraalObject struct {
	Stuk_id string `xml:"stuk_id,omitempty"`

	Betrokkene_id string `xml:"betrokkene_id,omitempty"`

	Omschrijving_aantekening string `xml:"omschrijving_aantekening,omitempty"`

	Beschrijving_aantekening string `xml:"beschrijving_aantekening,omitempty"`
}

type AantekeningKadastraalObjectArray struct {
	Item []*AantekeningKadastraalObject `xml:"item,omitempty"`
}

type AantekeningKadastraalObjectLijst struct {
	Aantekeningen *AantekeningKadastraalObjectArray `xml:"aantekeningen,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type PRBeperkingenKadastraalObject struct {
	Pr_beperkingen_lijst *PRBeperkingLijst `xml:"pr_beperkingen_lijst,omitempty"`

	Melding string `xml:"melding,omitempty"`
}

type PRBeperking struct {
	Stuk_id string `xml:"stuk_id,omitempty"`

	Betrokkene_id string `xml:"betrokkene_id,omitempty"`

	Gemeente_id string `xml:"gemeente_id,omitempty"`

	Gemeentelijke_registratie_id string `xml:"gemeentelijke_registratie_id,omitempty"`

	Omschrijving_pr_beperking string `xml:"omschrijving_pr_beperking,omitempty"`

	Beschrijving_pr_beperking string `xml:"beschrijving_pr_beperking,omitempty"`

	Splitsing_pr_beperking string `xml:"splitsing_pr_beperking,omitempty"`
}

type PRBeperkingArray struct {
	Item []*PRBeperking `xml:"item,omitempty"`
}

type PRBeperkingLijst struct {
	Pr_beperkingen *PRBeperkingArray `xml:"pr_beperkingen,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type Recht struct {
	Object_id string `xml:"object_id,omitempty"`

	Gerechtigde_id string `xml:"gerechtigde_id,omitempty"`

	Omschrijving_recht string `xml:"omschrijving_recht,omitempty"`

	Aandeel_in_recht string `xml:"aandeel_in_recht,omitempty"`

	Totaal_aandeel_in_recht string `xml:"totaal_aandeel_in_recht,omitempty"`

	Einddatum_recht time.Time `xml:"einddatum_recht,omitempty"`

	Indicatie_recht_betrokken_in_splitsing bool `xml:"indicatie_recht_betrokken_in_splitsing,omitempty"`

	Ontleend_aan_stukken_lijst *StukMetKadastraleAanduidingLijst `xml:"ontleend_aan_stukken_lijst,omitempty"`

	Mogelijk_van_belang_stukken_lijst *StukReferentiesLijst `xml:"mogelijk_van_belang_stukken_lijst,omitempty"`

	Aantekeningen_recht_lijst *AantekeningRechtLijst `xml:"aantekeningen_recht_lijst,omitempty"`
}

type RechtArray struct {
	Item []*Recht `xml:"item,omitempty"`
}

type RechtenLijst struct {
	Rechten *RechtArray `xml:"rechten,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type AantekeningRecht struct {
	Stuk_id string `xml:"stuk_id,omitempty"`

	Betrokkene_id string `xml:"betrokkene_id,omitempty"`

	Omschrijving_aantekening string `xml:"omschrijving_aantekening,omitempty"`

	Beschrijving_aantekening string `xml:"beschrijving_aantekening,omitempty"`

	Einddatum_aantekening time.Time `xml:"einddatum_aantekening,omitempty"`
}

type AantekeningRechtArray struct {
	Item []*AantekeningRecht `xml:"item,omitempty"`
}

type AantekeningRechtLijst struct {
	Aantekeningen_recht *AantekeningRechtArray `xml:"aantekeningen_recht,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type Persoon struct {
	Id string `xml:"id,omitempty"`

	Natuurlijk_persoon *NatuurlijkPersoon `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *NietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`

	Aantekening_beschikkingsbevoegdheid bool `xml:"aantekening_beschikkingsbevoegdheid,omitempty"`

	Indicatie_meer_onroerende_zaken bool `xml:"indicatie_meer_onroerende_zaken,omitempty"`

	Woonlocatie *Locatie `xml:"woonlocatie,omitempty"`

	Postlocatie *Locatie `xml:"postlocatie,omitempty"`

	Niet_volledig_verwerkt_stukken_lijst *StukReferentiesLijst `xml:"niet_volledig_verwerkt_stukken_lijst,omitempty"`

	Mogelijk_van_belang_stukken_lijst *StukReferentiesLijst `xml:"mogelijk_van_belang_stukken_lijst,omitempty"`
}

type PersoonArray struct {
	Item []*Persoon `xml:"item,omitempty"`
}

type PersoonV2 struct {
	Id string `xml:"id,omitempty"`

	Natuurlijk_persoon *NatuurlijkPersoonV2 `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *NietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`

	Aantekening_beschikkingsbevoegdheid bool `xml:"aantekening_beschikkingsbevoegdheid,omitempty"`

	Indicatie_meer_onroerende_zaken bool `xml:"indicatie_meer_onroerende_zaken,omitempty"`

	Woonlocatie *Locatie `xml:"woonlocatie,omitempty"`

	Postlocatie *Locatie `xml:"postlocatie,omitempty"`

	Niet_volledig_verwerkt_stukken_lijst *StukReferentiesLijst `xml:"niet_volledig_verwerkt_stukken_lijst,omitempty"`

	Mogelijk_van_belang_stukken_lijst *StukReferentiesLijst `xml:"mogelijk_van_belang_stukken_lijst,omitempty"`
}

type PersoonV2Array struct {
	Item []*PersoonV2 `xml:"item,omitempty"`
}

type KadasterPersoon struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterPersoon"`

	Akr_subject_nr string `xml:"akr_subject_nr,omitempty"`

	Kadastrale_identificatie string `xml:"kadastrale_identificatie,omitempty"`

	Natuurlijk_persoon *KadasterNatuurlijkPersoon `xml:"natuurlijk_persoon,omitempty"`

	Niet_natuurlijk_persoon *KadasterNietNatuurlijkPersoon `xml:"niet_natuurlijk_persoon,omitempty"`

	Aantekening_beschikkingsbevoegdheid bool `xml:"aantekening_beschikkingsbevoegdheid,omitempty"`

	Woonlocatie *KadasterLocatie `xml:"woonlocatie,omitempty"`

	Postlocatie *KadasterLocatie `xml:"postlocatie,omitempty"`

	Domicilie_keuzes *KadasterDomicilieKeuzeLijst `xml:"domicilie_keuzes,omitempty"`
}

type KadasterPersoonArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterPersoonArray"`

	Item []*KadasterPersoon `xml:"item,omitempty"`
}

type KadasterDomicilieKeuzeLijst struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterDomicilieKeuzeLijst"`

	Domicilie_keuzes *KadasterDomicilieKeuzeArray `xml:"domicilie_keuzes,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type KadasterDomicilieKeuze struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterDomicilieKeuze"`

	Voorletters string `xml:"voorletters,omitempty"`

	Voorvoegsels string `xml:"voorvoegsels,omitempty"`

	Achternaam string `xml:"achternaam,omitempty"`

	Adres string `xml:"adres,omitempty"`
}

type KadasterDomicilieKeuzeArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterDomicilieKeuzeArray"`

	Item []*KadasterDomicilieKeuze `xml:"item,omitempty"`
}

type NatuurlijkPersoon struct {
	Geslachtsnaam string `xml:"geslachtsnaam,omitempty"`

	Voorvoegsel_geslachtsnaam string `xml:"voorvoegsel_geslachtsnaam,omitempty"`

	Voornamen string `xml:"voornamen,omitempty"`

	Geslacht string `xml:"geslacht,omitempty"`

	Geboorteplaats string `xml:"geboorteplaats,omitempty"`

	Geboortedatum time.Time `xml:"geboortedatum,omitempty"`

	Datum_overlijden time.Time `xml:"datum_overlijden,omitempty"`

	Indicatie_overleden bool `xml:"indicatie_overleden,omitempty"`

	Indicatie_conform_gba bool `xml:"indicatie_conform_gba,omitempty"`

	Relatie_partnerid string `xml:"relatie_partnerid,omitempty"`

	Omschrijving_relatie string `xml:"omschrijving_relatie,omitempty"`
}

type NatuurlijkPersoonV2 struct {
	Adellijke_titel_geslachtsnaam string `xml:"adellijke_titel_geslachtsnaam,omitempty"`

	Geslachtsnaam string `xml:"geslachtsnaam,omitempty"`

	Voorvoegsel_geslachtsnaam string `xml:"voorvoegsel_geslachtsnaam,omitempty"`

	Adellijke_titel string `xml:"adellijke_titel,omitempty"`

	Voornamen string `xml:"voornamen,omitempty"`

	Geslacht string `xml:"geslacht,omitempty"`

	Geboorteplaats string `xml:"geboorteplaats,omitempty"`

	Geboortedatum time.Time `xml:"geboortedatum,omitempty"`

	Datum_overlijden time.Time `xml:"datum_overlijden,omitempty"`

	Indicatie_overleden bool `xml:"indicatie_overleden,omitempty"`

	Indicatie_conform_gba bool `xml:"indicatie_conform_gba,omitempty"`

	Relatie_partnerid string `xml:"relatie_partnerid,omitempty"`

	Omschrijving_relatie string `xml:"omschrijving_relatie,omitempty"`
}

type KadasterNatuurlijkPersoon struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterNatuurlijkPersoon"`

	Bsn string `xml:"bsn,omitempty"`

	Geslachtsnaam string `xml:"geslachtsnaam,omitempty"`

	Voorvoegsel_geslachtsnaam string `xml:"voorvoegsel_geslachtsnaam,omitempty"`

	Voornamen string `xml:"voornamen,omitempty"`

	Geslacht string `xml:"geslacht,omitempty"`

	Geboorteplaats string `xml:"geboorteplaats,omitempty"`

	Geboortedatum time.Time `xml:"geboortedatum,omitempty"`

	Datum_overlijden time.Time `xml:"datum_overlijden,omitempty"`

	Indicatie_overleden bool `xml:"indicatie_overleden,omitempty"`

	Burgerlijke_staat string `xml:"burgerlijke_staat,omitempty"`
}

type NietNatuurlijkPersoon struct {
	Naam string `xml:"naam,omitempty"`

	Code string `xml:"code,omitempty"`

	Statutaire_zetel string `xml:"statutaire_zetel,omitempty"`

	Leden_lijst *LidLijst `xml:"leden_lijst,omitempty"`
}

type KadasterNietNatuurlijkPersoon struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterNietNatuurlijkPersoon"`

	Naam string `xml:"naam,omitempty"`

	Rechtsvorm string `xml:"rechtsvorm,omitempty"`

	Statutaire_zetel string `xml:"statutaire_zetel,omitempty"`

	Bedrijven_nummer string `xml:"bedrijven_nummer,omitempty"`
}

type Locatie struct {
	Locatie_binnenland *LocatieBinnenland `xml:"locatie_binnenland,omitempty"`

	Locatie_buitenland *LocatieBuitenland `xml:"locatie_buitenland,omitempty"`
}

type KadasterLocatie struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterLocatie"`

	Locatie_binnenland *KadasterLocatieBinnenland `xml:"locatie_binnenland,omitempty"`

	Locatie_buitenland *KadasterLocatieBuitenland `xml:"locatie_buitenland,omitempty"`
}

type LocatieBinnenland struct {
	Gemeente string `xml:"gemeente,omitempty"`

	Straatnaam string `xml:"straatnaam,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`
}

type LocatieBinnenlandArray struct {
	Item []*LocatieBinnenland `xml:"item,omitempty"`
}

type KadasterLocatieBinnenland struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterLocatieBinnenland"`

	Plaats string `xml:"plaats,omitempty"`

	Straatnaam string `xml:"straatnaam,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`

	Positie string `xml:"positie,omitempty"`
}

type KadasterLocatieBinnenlandArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterLocatieBinnenlandArray"`

	Item []*KadasterLocatieBinnenland `xml:"item,omitempty"`
}

type LocatieBinnenlandLijst struct {
	Locaties_binnenland *LocatieBinnenlandArray `xml:"locaties_binnenland,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type LocatieBuitenland struct {
	Adres_buitenland1 string `xml:"adres_buitenland1,omitempty"`

	Adres_buitenland2 string `xml:"adres_buitenland2,omitempty"`

	Adres_buitenland3 string `xml:"adres_buitenland3,omitempty"`

	Land string `xml:"land,omitempty"`
}

type KadasterLocatieBuitenland struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterLocatieBuitenland"`

	Adres string `xml:"adres,omitempty"`

	Woonplaats string `xml:"woonplaats,omitempty"`

	Regio string `xml:"regio,omitempty"`

	Land string `xml:"land,omitempty"`
}

type Lid struct {
	Id string `xml:"id,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`
}

type LidArray struct {
	Item []*Lid `xml:"item,omitempty"`
}

type LidLijst struct {
	Leden *LidArray `xml:"leden,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type Stuk struct {
	Id string `xml:"id,omitempty"`

	Aanduiding_soort_register string `xml:"aanduiding_soort_register,omitempty"`

	Deel string `xml:"deel,omitempty"`

	Nummer string `xml:"nummer,omitempty"`

	Reeks string `xml:"reeks,omitempty"`

	Datum_aanbieding time.Time `xml:"datum_aanbieding,omitempty"`
}

type StukArray struct {
	Item []*Stuk `xml:"item,omitempty"`
}

type StukMetKadastraleAanduiding struct {
	Id string `xml:"id,omitempty"`

	Kadastrale_aanduiding *KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`
}

type StukMetKadastraleAanduidingArray struct {
	Item []*StukMetKadastraleAanduiding `xml:"item,omitempty"`
}

type StukMetKadastraleAanduidingLijst struct {
	Stukken *StukMetKadastraleAanduidingArray `xml:"stukken,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type GemeenteStuk struct {
	Gemeentelijke_registratie_id string `xml:"gemeentelijke_registratie_id,omitempty"`

	Nummer string `xml:"nummer,omitempty"`

	Datum_in_werking time.Time `xml:"datum_in_werking,omitempty"`
}

type GemeenteStukArray struct {
	Item []*GemeenteStuk `xml:"item,omitempty"`
}

type StukReferentiesLijst struct {
	Stukken *StringArray `xml:"stukken,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type KadasterHypothecairBerichtResultaat struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtResultaat"`

	Overzicht *KadasterOverzicht `xml:"overzicht,omitempty"`

	Hypothecairbericht *KadasterHypothecairBerichtObject `xml:"hypothecairbericht,omitempty"`
}

type KadasterHypothecairBerichtResultaatV2 struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtResultaatV2"`

	Overzicht *KadasterOverzichtV2 `xml:"overzicht,omitempty"`

	Hypothecairbericht *KadasterHypothecairBerichtObject `xml:"hypothecairbericht,omitempty"`
}

type KadasterHypothecairBerichtObject struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtObject"`

	Datum_bijgewerkt string `xml:"datum_bijgewerkt,omitempty"`

	Onroerende_zaak *KadasterBeknoptOnroerendeZaak `xml:"onroerende_zaak,omitempty"`

	Hypotheek_stuk_lijst *KadasterHypotheekStukLijst `xml:"hypotheek_stuk_lijst,omitempty"`

	Hypothecair_document *KadasterBerichtObjectDocument `xml:"hypothecair_document,omitempty"`
}

type KadasterHypotheekStukLijst struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypotheekStukLijst"`

	Hypotheek_stukken *KadasterHypotheekStukArray `xml:"hypotheek_stukken,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type KadasterHypotheekStuk struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypotheekStuk"`

	Stuk *Stuk `xml:"stuk,omitempty"`

	Stukdelen *KadasterStukdeelArray `xml:"stukdelen,omitempty"`

	Datum_aanbieding string `xml:"datum_aanbieding,omitempty"`

	Datum_ondertekening string `xml:"datum_ondertekening,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Rectificatie_verzocht bool `xml:"rectificatie_verzocht,omitempty"`

	Eenzijdig_opzegbaar bool `xml:"eenzijdig_opzegbaar,omitempty"`

	Gefiatteerd bool `xml:"gefiatteerd,omitempty"`
}

type KadasterHypotheekStukArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypotheekStukArray"`

	Item []*KadasterHypotheekStuk `xml:"item,omitempty"`
}

type KadasterStukdeel struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterStukdeel"`

	Id string `xml:"id,omitempty"`

	Aard_stukdeel_omschrijving string `xml:"aard_stukdeel_omschrijving,omitempty"`

	Hypotheek *KadasterHypotheek `xml:"hypotheek,omitempty"`

	Schuldeisers_beslagleggers *KadasterPersoonLijst `xml:"schuldeisers_beslagleggers,omitempty"`

	Vordering *KadasterBedrag `xml:"vordering,omitempty"`

	Rechten *KadasterStukdeelRechtArray `xml:"rechten,omitempty"`

	Onroerende_zaken *KadasterBeknoptOnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Stukdeel_relaties *KadasterStukdeelReferentiesLijst `xml:"stukdeel_relaties,omitempty"`

	Hoort_bij *KadasterGerelateerdStukdeelArray `xml:"hoort_bij,omitempty"`
}

type KadasterStukdeelArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterStukdeelArray"`

	Item []*KadasterStukdeel `xml:"item,omitempty"`
}

type KadasterGerelateerdStukdeel struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterGerelateerdStukdeel"`

	Aard_stukdeel_omschrijving string `xml:"aard_stukdeel_omschrijving,omitempty"`

	Stuk *Stuk `xml:"stuk,omitempty"`

	Gefiatteerd bool `xml:"gefiatteerd,omitempty"`
}

type KadasterGerelateerdStukdeelArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterGerelateerdStukdeelArray"`

	Item []*KadasterGerelateerdStukdeel `xml:"item,omitempty"`
}

type KadasterStukdeelReferentiesLijst struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterStukdeelReferentiesLijst"`

	Stukdelen *StringArray `xml:"stukdelen,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type KadasterStukdeelRecht struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterStukdeelRecht"`

	Id string `xml:"id,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Aandeel_in_recht string `xml:"aandeel_in_recht,omitempty"`

	Totaal_aandeel_in_recht string `xml:"totaal_aandeel_in_recht,omitempty"`

	Object_id string `xml:"object_id,omitempty"`

	Gerelateerd_recht string `xml:"gerelateerd_recht,omitempty"`
}

type KadasterStukdeelRechtArray struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterStukdeelRechtArray"`

	Item []*KadasterStukdeelRecht `xml:"item,omitempty"`
}

type KadasterPersoonLijst struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterPersoonLijst"`

	Personen *KadasterPersoonArray `xml:"personen,omitempty"`

	Indicatie_compleet bool `xml:"indicatie_compleet,omitempty"`
}

type KadasterHypotheek struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypotheek"`

	Hoofdsom *KadasterBedrag `xml:"hoofdsom,omitempty"`

	Rentevoet float32 `xml:"rentevoet,omitempty"`

	Meerdere_rentevoet bool `xml:"meerdere_rentevoet,omitempty"`
}

type KadasterBedrag struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterBedrag"`

	Valuta string `xml:"valuta,omitempty"`

	Bedrag int64 `xml:"bedrag,omitempty"`
}

type KadasterBronDocument struct {
	Document []byte `xml:"document,omitempty"`

	Images *Base64BinaryArray `xml:"images,omitempty"`
}

type KadasterValueList struct {
	List *KadasterListValueArray `xml:"list,omitempty"`
}

type KadasterListValueArray struct {
	Item []*KadasterListValue `xml:"item,omitempty"`
}

type KadasterListValue struct {
	Key string `xml:"key,omitempty"`

	Value string `xml:"value,omitempty"`
}

type KadasterV2KadastraalBerichtResponse struct {
	Bericht *KadasterV2KadastraalBericht `xml:"bericht,omitempty"`

	Document []byte `xml:"document,omitempty"`

	Overgegaan_in *KadasterV2OvergegaanIn `xml:"overgegaan_in,omitempty"`

	Postcode_treffers *KadasterV2PostcodeTreffers `xml:"postcode_treffers,omitempty"`
}

type KadasterV2HypothecairBerichtResponse struct {
	Bericht *KadasterV2HypothecairBericht `xml:"bericht,omitempty"`

	Document []byte `xml:"document,omitempty"`

	Overgegaan_in *KadasterV2OvergegaanIn `xml:"overgegaan_in,omitempty"`

	Postcode_treffers *KadasterV2PostcodeTreffers `xml:"postcode_treffers,omitempty"`
}

type KadasterV2UittrekselKadastraleKaartResponse struct {
	Document []byte `xml:"document,omitempty"`

	Overgegaan_in *KadasterV2OvergegaanIn `xml:"overgegaan_in,omitempty"`

	Postcode_treffers *KadasterV2PostcodeTreffers `xml:"postcode_treffers,omitempty"`
}

type KadasterV2OvergegaanIn struct {
	Productienummer string `xml:"productienummer,omitempty"`

	Tijdstip_samenstelling_bericht time.Time `xml:"tijdstip_samenstelling_bericht,omitempty"`

	Actualiteit_tijdstip_hypothecair time.Time `xml:"actualiteit_tijdstip_hypothecair,omitempty"`

	Actualiteit_tijdstip_kadastraal time.Time `xml:"actualiteit_tijdstip_kadastraal,omitempty"`

	Signalering_tijdstip_beslagen time.Time `xml:"signalering_tijdstip_beslagen,omitempty"`

	Signalering_tijdstip_hypothecair time.Time `xml:"signalering_tijdstip_hypothecair,omitempty"`

	Signalering_tijdstip_kadastraal time.Time `xml:"signalering_tijdstip_kadastraal,omitempty"`

	Betreft *KadasterV2OnroerendeZaak `xml:"betreft,omitempty"`

	Tenaamstellingen *KadasterV2TenaamstellingArray `xml:"tenaamstellingen,omitempty"`

	Zakelijk_rechten *KadasterV2ZakelijkRechtArray `xml:"zakelijk_rechten,omitempty"`
}

type KadasterV2PostcodeTreffers struct {
	Productienummer string `xml:"productienummer,omitempty"`

	Tijdstip_samenstelling_bericht time.Time `xml:"tijdstip_samenstelling_bericht,omitempty"`

	Actualiteit_tijdstip_hypothecair time.Time `xml:"actualiteit_tijdstip_hypothecair,omitempty"`

	Actualiteit_tijdstip_kadastraal time.Time `xml:"actualiteit_tijdstip_kadastraal,omitempty"`

	Signalering_tijdstip_beslagen time.Time `xml:"signalering_tijdstip_beslagen,omitempty"`

	Signalering_tijdstip_hypothecair time.Time `xml:"signalering_tijdstip_hypothecair,omitempty"`

	Signalering_tijdstip_kadastraal time.Time `xml:"signalering_tijdstip_kadastraal,omitempty"`

	Betreft *KadasterV2GevraagdGebied `xml:"betreft,omitempty"`
}

type KadasterV2BronDocumentResponse struct {
	Document []byte `xml:"document,omitempty"`
}

type KadasterV2KadastraalBericht struct {
	Productienummer string `xml:"productienummer,omitempty"`

	Tijdstip_samenstelling_bericht time.Time `xml:"tijdstip_samenstelling_bericht,omitempty"`

	Actualiteit_tijdstip_hypothecair time.Time `xml:"actualiteit_tijdstip_hypothecair,omitempty"`

	Actualiteit_tijdstip_kadastraal time.Time `xml:"actualiteit_tijdstip_kadastraal,omitempty"`

	Signalering_tijdstip_beslagen time.Time `xml:"signalering_tijdstip_beslagen,omitempty"`

	Signalering_tijdstip_hypothecair time.Time `xml:"signalering_tijdstip_hypothecair,omitempty"`

	Signalering_tijdstip_kadastraal time.Time `xml:"signalering_tijdstip_kadastraal,omitempty"`

	Betreft *KadasterV2OnroerendeZaak `xml:"betreft,omitempty"`

	Tenaamstellingen *KadasterV2TenaamstellingArray `xml:"tenaamstellingen,omitempty"`

	Zakelijk_rechten *KadasterV2ZakelijkRechtArray `xml:"zakelijk_rechten,omitempty"`

	Aantekeningen *KadasterV2AantekeningArray `xml:"aantekeningen,omitempty"`

	Erfpachtcanon *KadasterV2ErfpachtCanonArray `xml:"erfpachtcanon,omitempty"`

	Stukken *KadasterV2Stukken `xml:"stukken,omitempty"`
}

type KadasterV2Tenaamstelling struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aandeel *KadasterV2Breuk `xml:"aandeel,omitempty"`

	Burgerlijke_staat_ten_tijde_van_verkrijging string `xml:"burgerlijke_staat_ten_tijde_van_verkrijging,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Verklaring_inzake_derden_bescherming string `xml:"verklaring_inzake_derden_bescherming,omitempty"`

	Verkregen_namens_samenwerkingsverband string `xml:"verkregen_namens_samenwerkingsverband,omitempty"`

	Betrokken_partner *KadasterV2Personen `xml:"betrokken_partner,omitempty"`

	Betrokken_samenwerkingsverband *KadasterV2Personen `xml:"betrokken_samenwerkingsverband,omitempty"`

	Van *KadasterV2ZakelijkRecht `xml:"van,omitempty"`

	Geldt_voor *KadasterV2GezamenlijkAandeel `xml:"geldt_voor,omitempty"`

	Gebaseerd_op *KadasterV2StukdeelArray `xml:"gebaseerd_op,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`

	Ten_name_van *KadasterV2Personen `xml:"ten_name_van,omitempty"`
}

type KadasterV2TenaamstellingArray struct {
	Item []*KadasterV2Tenaamstelling `xml:"item,omitempty"`
}

type KadasterV2AangebodenStuk struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aard string `xml:"aard,omitempty"`

	Deel_en_nummer *KadasterV2DeelEnNummer `xml:"deel_en_nummer,omitempty"`

	Verklaring_bewaarder string `xml:"verklaring_bewaarder,omitempty"`

	Toelichting_bewaarder string `xml:"toelichting_bewaarder,omitempty"`

	Tijdstip_aanbieding time.Time `xml:"tijdstip_aanbieding,omitempty"`

	Tijdstip_ondertekening time.Time `xml:"tijdstip_ondertekening,omitempty"`

	Status_stuk_or string `xml:"status_stuk_or,omitempty"`

	Indicatie_alle_objecten_betrokken bool `xml:"indicatie_alle_objecten_betrokken,omitempty"`

	Verzoeken *KadasterV2VerzoekArray `xml:"verzoeken,omitempty"`

	Stukdelen *KadasterV2StukdeelArray `xml:"stukdelen,omitempty"`
}

type KadasterV2AangebodenStukArray struct {
	Item []*KadasterV2AangebodenStuk `xml:"item,omitempty"`
}

type KadasterV2DeelEnNummer struct {
	Deel string `xml:"deel,omitempty"`

	Nummer string `xml:"nummer,omitempty"`

	Reeks string `xml:"reeks,omitempty"`

	Registercode string `xml:"registercode,omitempty"`

	Soort_register string `xml:"soort_register,omitempty"`
}

type KadasterV2ErfpachtCanon struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Soort string `xml:"soort,omitempty"`

	Einddatum_afkoop time.Time `xml:"einddatum_afkoop,omitempty"`

	Indicatie_oude_onroerende_zaak_betrokken bool `xml:"indicatie_oude_onroerende_zaak_betrokken,omitempty"`

	Jaarlijks_bedrag *KadasterV2JaarlijksBedrag `xml:"jaarlijks_bedrag,omitempty"`

	Betreft *KadasterV2ZakelijkRecht `xml:"betreft,omitempty"`

	Gebaseerd_op *KadasterV2Stukdeel `xml:"gebaseerd_op,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`
}

type KadasterV2ErfpachtCanonArray struct {
	Item []*KadasterV2ErfpachtCanon `xml:"item,omitempty"`
}

type KadasterV2JaarlijksBedrag struct {
	Bedrag *KadasterV2Bedrag `xml:"bedrag,omitempty"`

	Betreft_meer_onroerende_zaken bool `xml:"betreft_meer_onroerende_zaken,omitempty"`
}

type KadasterV2JaarlijksBedragArray struct {
	Item []*KadasterV2JaarlijksBedrag `xml:"item,omitempty"`
}

type KadasterV2Verzoek struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aard string `xml:"aard,omitempty"`

	Redenen_verzoek *KadasterV2RedenVerzoekArray `xml:"redenen_verzoek,omitempty"`
}

type KadasterV2VerzoekArray struct {
	Item []*KadasterV2Verzoek `xml:"item,omitempty"`
}

type KadasterV2Bijlage struct {
	Aard string `xml:"aard,omitempty"`

	Tijdstip_aanbieding_bijlage time.Time `xml:"tijdstip_aanbieding_bijlage,omitempty"`
}

type KadasterV2BijlageArray struct {
	Item []*KadasterV2Bijlage `xml:"item,omitempty"`
}

type KadasterV2GezamenlijkAandeel struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aandeel *KadasterV2Breuk `xml:"aandeel,omitempty"`
}

type KadasterV2GezamenlijkAandeelArray struct {
	Item []*KadasterV2GezamenlijkAandeel `xml:"item,omitempty"`
}

type KadasterV2Breuk struct {
	Teller int32 `xml:"teller,omitempty"`

	Noemer int32 `xml:"noemer,omitempty"`
}

type KadasterV2OnroerendeZaak struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Kadastrale_aanduiding *KadasterV2KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Adressen *KadasterV2AdresLocatieArray `xml:"adressen,omitempty"`

	Tekeningen *KadasterV2StukdeelArray `xml:"tekeningen,omitempty"`

	Uin string `xml:"uin,omitempty"`

	Aard string `xml:"aard,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Kadastrale_grootte *KadasterV2KadastraleGrootte `xml:"kadastrale_grootte,omitempty"`

	Plaatscoordinaten *KadasterV2Coordinaat `xml:"plaatscoordinaten,omitempty"`

	Belast_met *KadasterV2LandinrichtingsRenteArray `xml:"belast_met,omitempty"`

	Aard_cultuur_onbebouwd string `xml:"aard_cultuur_onbebouwd,omitempty"`

	Aard_cultuur_bebouwd string `xml:"aard_cultuur_bebouwd,omitempty"`

	Omschrijving_onderzoek_erfdienstbaarheden string `xml:"omschrijving_onderzoek_erfdienstbaarheden,omitempty"`

	Toestandsdatum_onderzoek_erfdienstbaarheden time.Time `xml:"toestandsdatum_onderzoek_erfdienstbaarheden,omitempty"`

	Koopsom *KadasterV2Koopsom `xml:"koopsom,omitempty"`

	Indicatie_voorlopige_aantekening bool `xml:"indicatie_voorlopige_aantekening,omitempty"`

	Indicatie_tweede_splitsing bool `xml:"indicatie_tweede_splitsing,omitempty"`

	Indicatie_gedeeltelijk_bezwaard_oud_object bool `xml:"indicatie_gedeeltelijk_bezwaard_oud_object,omitempty"`

	Indicatie_gesplitst bool `xml:"indicatie_gesplitst,omitempty"`

	Indicatie_voorgenomen_aantekening bool `xml:"indicatie_voorgenomen_aantekening,omitempty"`

	Toelichting_bewaarder string `xml:"toelichting_bewaarder,omitempty"`

	Historie *KadasterV2Historie `xml:"historie,omitempty"`

	Overgegaan_in *KadasterV2OnroerendeZaakFiliatieArray `xml:"overgegaan_in,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`

	Ontstaan_uit *KadasterV2OnroerendeZaakFiliatieArray `xml:"ontstaan_uit,omitempty"`

	Aantekeningen_wkpb *KadasterV2AantekeningWKPBArray `xml:"aantekeningen_wkpb,omitempty"`

	Aanvullingen *KadasterV2AanvullingArray `xml:"aanvullingen,omitempty"`
}

type KadasterV2OnroerendeZaakArray struct {
	Item []*KadasterV2OnroerendeZaak `xml:"item,omitempty"`
}

type KadasterV2OnroerendeZaakReference struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Kadastrale_aanduiding *KadasterV2KadastraleAanduiding `xml:"kadastrale_aanduiding,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type KadasterV2OnroerendeZaakReferenceArray struct {
	Item []*KadasterV2OnroerendeZaakReference `xml:"item,omitempty"`
}

type KadasterV2Aanvulling struct {
	Type_ string `xml:"type,omitempty"`

	Aangeboden_stuk *KadasterV2AangebodenStuk `xml:"aangeboden_stuk,omitempty"`
}

type KadasterV2AanvullingArray struct {
	Item []*KadasterV2Aanvulling `xml:"item,omitempty"`
}

type KadasterV2AdresLocatie struct {
	Locatie_binnenland *KadasterV2LocatieBinnenland `xml:"locatie_binnenland,omitempty"`

	Locatie_buitenland *KadasterV2LocatieBuitenland `xml:"locatie_buitenland,omitempty"`

	Locatie_postbus *KadasterV2LocatiePostbus `xml:"locatie_postbus,omitempty"`
}

type KadasterV2AdresLocatieArray struct {
	Item []*KadasterV2AdresLocatie `xml:"item,omitempty"`
}

type KadasterV2LocatieBinnenland struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Straatnaam string `xml:"straatnaam,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`

	Huisnummertoevoeging string `xml:"huisnummertoevoeging,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	In_onderzoek string `xml:"in_onderzoek,omitempty"`

	Woonplaats string `xml:"woonplaats,omitempty"`
}

type KadasterV2LocatieBinnenlandArray struct {
	Item []*KadasterV2LocatieBinnenland `xml:"item,omitempty"`
}

type KadasterV2LocatieBuitenland struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Adres string `xml:"adres,omitempty"`

	Adres_regels *StringArray `xml:"adres_regels,omitempty"`

	Woonplaats string `xml:"woonplaats,omitempty"`

	Regio string `xml:"regio,omitempty"`

	Land string `xml:"land,omitempty"`
}

type KadasterV2LocatieBuitenlandArray struct {
	Item []*KadasterV2LocatieBuitenland `xml:"item,omitempty"`
}

type KadasterV2LocatiePostbus struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Postbus_nummer string `xml:"postbus_nummer,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Woonplaats string `xml:"woonplaats,omitempty"`
}

type KadasterV2LocatiePostbusArray struct {
	Item []*KadasterV2LocatiePostbus `xml:"item,omitempty"`
}

type KadasterV2AantekeningWKPB struct {
	Omschrijving string `xml:"omschrijving,omitempty"`

	Gemeentelijke_registratie_id string `xml:"gemeentelijke_registratie_id,omitempty"`

	Nummer string `xml:"nummer,omitempty"`

	Datum_in_werking string `xml:"datum_in_werking,omitempty"`

	Omschrijving_publiekrechtelijke_beperking string `xml:"omschrijving_publiekrechtelijke_beperking,omitempty"`
}

type KadasterV2AantekeningWKPBArray struct {
	Item []*KadasterV2AantekeningWKPB `xml:"item,omitempty"`
}

type KadasterV2KadastraleAanduiding struct {
	Kadastrale_gemeente string `xml:"kadastrale_gemeente,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer int32 `xml:"perceelnummer,omitempty"`

	Appartementsrecht_volgnummer int32 `xml:"appartementsrecht_volgnummer,omitempty"`
}

type KadasterV2LandinrichtingsRente struct {
	Bedrag *KadasterV2Bedrag `xml:"bedrag,omitempty"`

	Eindjaar int32 `xml:"eindjaar,omitempty"`
}

type KadasterV2LandinrichtingsRenteArray struct {
	Item []*KadasterV2LandinrichtingsRente `xml:"item,omitempty"`
}

type KadasterV2Historie struct {
	Logisch_tijdstip_ontstaan time.Time `xml:"logisch_tijdstip_ontstaan,omitempty"`

	Logisch_tijdstip_vervallen time.Time `xml:"logisch_tijdstip_vervallen,omitempty"`

	Volgnummer int32 `xml:"volgnummer,omitempty"`

	Status_historie string `xml:"status_historie,omitempty"`

	Technisch_tijdstip_ontstaan time.Time `xml:"technisch_tijdstip_ontstaan,omitempty"`

	Technisch_tijdstip_vervallen time.Time `xml:"technisch_tijdstip_vervallen,omitempty"`
}

type KadasterV2OnroerendeZaakFiliatie struct {
	Aard string `xml:"aard,omitempty"`

	Overgangsgrootte float32 `xml:"overgangsgrootte,omitempty"`

	Betreft *KadasterV2OnroerendeZaakReferenceArray `xml:"betreft,omitempty"`
}

type KadasterV2OnroerendeZaakFiliatieArray struct {
	Item []*KadasterV2OnroerendeZaakFiliatie `xml:"item,omitempty"`
}

type KadasterV2Coordinaat struct {
	Identifier int32 `xml:"identifier,omitempty"`

	Name int32 `xml:"name,omitempty"`

	Description int32 `xml:"description,omitempty"`

	Rd_coordinaat *KadasterV2RDCoordinaat `xml:"rd_coordinaat,omitempty"`
}

type KadasterV2RDCoordinaat struct {
	X string `xml:"x,omitempty"`

	Y string `xml:"y,omitempty"`
}

type KadasterV2KoopsommenOverzichtResponse struct {
	Overzicht *KadasterV2KoopsommenOverzicht `xml:"overzicht,omitempty"`

	Document []byte `xml:"document,omitempty"`
}

type KadasterV2KoopsommenOverzicht struct {
	Productienummer string `xml:"productienummer,omitempty"`

	Adres_range string `xml:"adres_range,omitempty"`

	Woonplaats_naam string `xml:"woonplaats_naam,omitempty"`

	Actualiteittijdstip_koopsommen time.Time `xml:"actualiteittijdstip_koopsommen,omitempty"`

	Tijdstip_samenstelling_bericht time.Time `xml:"tijdstip_samenstelling_bericht,omitempty"`

	Gevraagd_gebied *KadasterV2GevraagdGebied `xml:"gevraagd_gebied,omitempty"`
}

type KadasterV2GevraagdGebied struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`

	Huisnummertoevoeging string `xml:"huisnummertoevoeging,omitempty"`

	Koopsom_gegevens *KadasterV2KoopsomGegevenArray `xml:"koopsom_gegevens,omitempty"`

	Onroerende_zaken *KadasterV2OnroerendeZaakArray `xml:"onroerende_zaken,omitempty"`
}

type KadasterV2KoopsomGegeven struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Huisnummer_volledig string `xml:"huisnummer_volledig,omitempty"`

	Omschrijving_cultuur string `xml:"omschrijving_cultuur,omitempty"`

	Kadastrale_grootte *KadasterV2KadastraleGrootte `xml:"kadastrale_grootte,omitempty"`

	Koopsom *KadasterV2Koopsom `xml:"koopsom,omitempty"`
}

type KadasterV2KoopsomGegevenArray struct {
	Item []*KadasterV2KoopsomGegeven `xml:"item,omitempty"`
}

type KadasterV2KadastraleGrootte struct {
	Waarde float32 `xml:"waarde,omitempty"`

	Soort_grootte string `xml:"soort_grootte,omitempty"`
}

type KadasterV2Koopsom struct {
	Bedrag *KadasterV2Bedrag `xml:"bedrag,omitempty"`

	Indicatie_meer_objecten bool `xml:"indicatie_meer_objecten,omitempty"`

	Koopjaar int32 `xml:"koopjaar,omitempty"`

	Koopdatum time.Time `xml:"koopdatum,omitempty"`
}

type KadasterV2Bedrag struct {
	Som float32 `xml:"som,omitempty"`

	Valuta string `xml:"valuta,omitempty"`
}

type KadasterV2GerelateerdeOpenbareRuimte struct {
	Openbare_ruimte_naam string `xml:"openbare_ruimte_naam,omitempty"`

	Openbare_ruimte_plaats string `xml:"openbare_ruimte_plaats,omitempty"`

	In_onderzoek string `xml:"in_onderzoek,omitempty"`
}

type KadasterV2NatuurlijkPersoon struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Indicatie_beschikkingsbevoegdheid bool `xml:"indicatie_beschikkingsbevoegdheid,omitempty"`

	Bsn string `xml:"bsn,omitempty"`

	Aanduiding_naamgebruik string `xml:"aanduiding_naamgebruik,omitempty"`

	Adellijke_titel_of_predikaat string `xml:"adellijke_titel_of_predikaat,omitempty"`

	Voorletters string `xml:"voorletters,omitempty"`

	Voornamen string `xml:"voornamen,omitempty"`

	Voorvoegsel string `xml:"voorvoegsel,omitempty"`

	Geslachtnaam string `xml:"geslachtnaam,omitempty"`

	Geslacht string `xml:"geslacht,omitempty"`

	Datum_geboorte time.Time `xml:"datum_geboorte,omitempty"`

	Indicatie_overleden bool `xml:"indicatie_overleden,omitempty"`

	Datum_overlijden time.Time `xml:"datum_overlijden,omitempty"`

	Land_waarnaar_vertrokken string `xml:"land_waarnaar_vertrokken,omitempty"`

	Bpr_gekoppeld bool `xml:"bpr_gekoppeld,omitempty"`

	Burgerlijke_staat_omschrijving string `xml:"burgerlijke_staat_omschrijving,omitempty"`

	Postlocatie *KadasterV2AdresLocatie `xml:"postlocatie,omitempty"`

	Woonlocatie *KadasterV2AdresLocatie `xml:"woonlocatie,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`

	Aanvullingen *KadasterV2AanvullingArray `xml:"aanvullingen,omitempty"`
}

type KadasterV2NatuurlijkPersoonArray struct {
	Item []*KadasterV2NatuurlijkPersoon `xml:"item,omitempty"`
}

type KadasterV2NietNatuurlijkPersoon struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Indicatie_beschikkingsbevoegdheid bool `xml:"indicatie_beschikkingsbevoegdheid,omitempty"`

	Statutaire_naam string `xml:"statutaire_naam,omitempty"`

	Statutaire_zetel string `xml:"statutaire_zetel,omitempty"`

	Rechtsvorm string `xml:"rechtsvorm,omitempty"`

	Dossier_nummer string `xml:"dossier_nummer,omitempty"`

	Rsin string `xml:"rsin,omitempty"`

	Postlocatie *KadasterV2AdresLocatie `xml:"postlocatie,omitempty"`

	Woonlocatie *KadasterV2AdresLocatie `xml:"woonlocatie,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`

	Aanvullingen *KadasterV2AanvullingArray `xml:"aanvullingen,omitempty"`
}

type KadasterV2NietNatuurlijkPersoonArray struct {
	Item []*KadasterV2NietNatuurlijkPersoon `xml:"item,omitempty"`
}

type KadasterV2ZakelijkRecht struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aard string `xml:"aard,omitempty"`

	Onroerende_zaak *KadasterV2OnroerendeZaak `xml:"onroerende_zaak,omitempty"`

	Belast_met *KadasterV2ZakelijkRechtArray `xml:"belast_met,omitempty"`

	Bestemt_tot *KadasterV2Mandeligheid `xml:"bestemt_tot,omitempty"`

	Ontstaan_uit *KadasterV2Splitsing `xml:"ontstaan_uit,omitempty"`

	Betrokken_bij *KadasterV2Splitsing `xml:"betrokken_bij,omitempty"`

	Beperkt_tot *KadasterV2TenaamstellingArray `xml:"beperkt_tot,omitempty"`

	Indicatie_splitsing_betrokken bool `xml:"indicatie_splitsing_betrokken,omitempty"`
}

type KadasterV2ZakelijkRechtArray struct {
	Item []*KadasterV2ZakelijkRecht `xml:"item,omitempty"`
}

type KadasterV2Personen struct {
	Natuurlijk_personen *KadasterV2NatuurlijkPersoonArray `xml:"natuurlijk_personen,omitempty"`

	Niet_natuurlijk_personen *KadasterV2NietNatuurlijkPersoonArray `xml:"niet_natuurlijk_personen,omitempty"`
}

type KadasterV2Aantekening struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aard string `xml:"aard,omitempty"`

	Omschrijving string `xml:"omschrijving,omitempty"`

	Einddatum time.Time `xml:"einddatum,omitempty"`

	Einddatum_recht time.Time `xml:"einddatum_recht,omitempty"`

	Herkaveling_bloknummer int32 `xml:"herkaveling_bloknummer,omitempty"`

	Betreft *KadasterV2OnroerendeZaak `xml:"betreft,omitempty"`

	Hypotheken *KadasterV2HypotheekArray `xml:"hypotheken,omitempty"`

	Beslagleggingen *KadasterV2BeslagleggingArray `xml:"beslagleggingen,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`

	Tenaamstellingen *KadasterV2TenaamstellingArray `xml:"tenaamstellingen,omitempty"`

	Gebaseerd_op *KadasterV2Stukdeel `xml:"gebaseerd_op,omitempty"`

	Betrokken_personen *KadasterV2Personen `xml:"betrokken_personen,omitempty"`
}

type KadasterV2AantekeningArray struct {
	Item []*KadasterV2Aantekening `xml:"item,omitempty"`
}

type KadasterV2Mandeligheid struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Hoort_bij *KadasterV2HoofdzaakArray `xml:"hoort_bij,omitempty"`

	Gebaseerd_op *KadasterV2StukdeelArray `xml:"gebaseerd_op,omitempty"`
}

type KadasterV2MandeligheidArray struct {
	Item []*KadasterV2Mandeligheid `xml:"item,omitempty"`
}

type KadasterV2Hoofdzaak struct {
	Aandeel_in_mandeligheid *KadasterV2Breuk `xml:"aandeel_in_mandeligheid,omitempty"`

	Betreft *KadasterV2OnroerendeZaak `xml:"betreft,omitempty"`
}

type KadasterV2HoofdzaakArray struct {
	Item []*KadasterV2Hoofdzaak `xml:"item,omitempty"`
}

type KadasterV2Splitsing struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Vereniging_van_eigenaren *KadasterV2NietNatuurlijkPersoon `xml:"vereniging_van_eigenaren,omitempty"`

	Gebaseerd_op *KadasterV2StukdeelArray `xml:"gebaseerd_op,omitempty"`
}

type KadasterV2SplitsingArray struct {
	Item []*KadasterV2Splitsing `xml:"item,omitempty"`
}

type KadasterV2Stukken struct {
	Stukken *KadasterV2StukArray `xml:"stukken,omitempty"`

	Aangeboden_stukken *KadasterV2AangebodenStukArray `xml:"aangeboden_stukken,omitempty"`
}

type KadasterV2RedenVerzoek struct {
	Reden string `xml:"reden,omitempty"`

	Reden_omschrijving string `xml:"reden_omschrijving,omitempty"`
}

type KadasterV2RedenVerzoekArray struct {
	Item []*KadasterV2RedenVerzoek `xml:"item,omitempty"`
}

type KadasterV2Stuk struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Indicatie_correctie bool `xml:"indicatie_correctie,omitempty"`

	Toelichting_bewaarder string `xml:"toelichting_bewaarder,omitempty"`

	Acg_identificatie *KadasterV2AcgIdentificatie `xml:"acg_identificatie,omitempty"`

	Stukdelen *KadasterV2StukdeelArray `xml:"stukdelen,omitempty"`

	Bijlages *KadasterV2BijlageArray `xml:"bijlages,omitempty"`
}

type KadasterV2StukArray struct {
	Item []*KadasterV2Stuk `xml:"item,omitempty"`
}

type KadasterV2AcgIdentificatie struct {
	Akr_registercode string `xml:"akr_registercode,omitempty"`

	Akr_stukdelen *StringArray `xml:"akr_stukdelen,omitempty"`

	Volgnummer_staat75 int32 `xml:"volgnummer_staat75,omitempty"`
}

type KadasterV2HypothecairBericht struct {
	Productienummer string `xml:"productienummer,omitempty"`

	Tijdstip_samenstelling_bericht time.Time `xml:"tijdstip_samenstelling_bericht,omitempty"`

	Actualiteit_tijdstip_hypothecair time.Time `xml:"actualiteit_tijdstip_hypothecair,omitempty"`

	Actualiteit_tijdstip_kadastraal time.Time `xml:"actualiteit_tijdstip_kadastraal,omitempty"`

	Signalering_tijdstip_beslagen time.Time `xml:"signalering_tijdstip_beslagen,omitempty"`

	Signalering_tijdstip_hypothecair time.Time `xml:"signalering_tijdstip_hypothecair,omitempty"`

	Signalering_tijdstip_kadastraal time.Time `xml:"signalering_tijdstip_kadastraal,omitempty"`

	Betreft *KadasterV2OnroerendeZaak `xml:"betreft,omitempty"`

	Hypotheken *KadasterV2HypotheekArray `xml:"hypotheken,omitempty"`

	Beslagleggingen *KadasterV2BeslagleggingArray `xml:"beslagleggingen,omitempty"`

	Zakelijk_rechten *KadasterV2ZakelijkRechtArray `xml:"zakelijk_rechten,omitempty"`

	Aantekeningen *KadasterV2AantekeningArray `xml:"aantekeningen,omitempty"`

	Stukken *KadasterV2Stukken `xml:"stukken,omitempty"`
}

type KadasterV2Stukdeel struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Aard string `xml:"aard,omitempty"`

	Omschrijving_gekozen_woonplaats string `xml:"omschrijving_gekozen_woonplaats,omitempty"`

	Omschrijving_kadastrale_objecten string `xml:"omschrijving_kadastrale_objecten,omitempty"`

	Bedrag_transacties_om_levering *KadasterV2Bedrag `xml:"bedrag_transacties_om_levering,omitempty"`

	Bedrag_vorderingsbeslag *KadasterV2Bedrag `xml:"bedrag_vorderingsbeslag,omitempty"`

	Bedrag_zekerheidsstelling_hypotheek *KadasterV2Bedrag `xml:"bedrag_zekerheidsstelling_hypotheek,omitempty"`

	Aanvulling_op *KadasterV2StukdeelArray `xml:"aanvulling_op,omitempty"`

	Vermeld_in *KadasterV2OnroerendeZaakReferenceArray `xml:"vermeld_in,omitempty"`
}

type KadasterV2StukdeelArray struct {
	Item []*KadasterV2Stukdeel `xml:"item,omitempty"`
}

type KadasterV2Hypotheek struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Omschrijving_betrokken_recht string `xml:"omschrijving_betrokken_recht,omitempty"`

	Aandeel_in_betrokken_recht *KadasterV2Breuk `xml:"aandeel_in_betrokken_recht,omitempty"`

	Onroerende_zaak *KadasterV2OnroerendeZaak `xml:"onroerende_zaak,omitempty"`

	Personen *KadasterV2Personen `xml:"personen,omitempty"`

	Gebaseerd_op *KadasterV2Stukdeel `xml:"gebaseerd_op,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`
}

type KadasterV2HypotheekArray struct {
	Item []*KadasterV2Hypotheek `xml:"item,omitempty"`
}

type KadasterV2Beslaglegging struct {
	Identificatie string `xml:"identificatie,omitempty"`

	Omschrijving_betrokken_recht string `xml:"omschrijving_betrokken_recht,omitempty"`

	Aard_beslag string `xml:"aard_beslag,omitempty"`

	Aandeel_in_betrokken_recht *KadasterV2Breuk `xml:"aandeel_in_betrokken_recht,omitempty"`

	Onroerende_zaak *KadasterV2OnroerendeZaak `xml:"onroerende_zaak,omitempty"`

	Personen *KadasterV2Personen `xml:"personen,omitempty"`

	Gebaseerd_op *KadasterV2Stukdeel `xml:"gebaseerd_op,omitempty"`

	Vermeld_in *KadasterV2StukdeelArray `xml:"vermeld_in,omitempty"`
}

type KadasterV2BeslagleggingArray struct {
	Item []*KadasterV2Beslaglegging `xml:"item,omitempty"`
}

type KvkDossier struct {
	Update_info *KvkUpdateReference `xml:"update_info,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Main_establishment_number string `xml:"main_establishment_number,omitempty"`

	Indication_main_establishment bool `xml:"indication_main_establishment,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Chamber_number string `xml:"chamber_number,omitempty"`

	Legal_form_code string `xml:"legal_form_code,omitempty"`

	Legal_form_text string `xml:"legal_form_text,omitempty"`

	Indication_organisation_code string `xml:"indication_organisation_code,omitempty"`

	Legal_name string `xml:"legal_name,omitempty"`

	Trade_name_45 string `xml:"trade_name_45,omitempty"`

	Trade_name_full string `xml:"trade_name_full,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Establishment_postcode string `xml:"establishment_postcode,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Establishment_house_number int32 `xml:"establishment_house_number,omitempty"`

	Establishment_house_number_addition string `xml:"establishment_house_number_addition,omitempty"`

	Correspondence_postcode string `xml:"correspondence_postcode,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Correspondence_house_number int32 `xml:"correspondence_house_number,omitempty"`

	Correspondence_house_number_addition string `xml:"correspondence_house_number_addition,omitempty"`

	Correspondence_country string `xml:"correspondence_country,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Mobile_number string `xml:"mobile_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Contact_title1 string `xml:"contact_title1,omitempty"`

	Contact_title2 string `xml:"contact_title2,omitempty"`

	Contact_initials string `xml:"contact_initials,omitempty"`

	Contact_prefix string `xml:"contact_prefix,omitempty"`

	Contact_surname string `xml:"contact_surname,omitempty"`

	Contact_gender string `xml:"contact_gender,omitempty"`

	Primary_sbi_code string `xml:"primary_sbi_code,omitempty"`

	Secondary_sbi_code1 string `xml:"secondary_sbi_code1,omitempty"`

	Secondary_sbi_code2 string `xml:"secondary_sbi_code2,omitempty"`

	Primary_sbi_code_text string `xml:"primary_sbi_code_text,omitempty"`

	Secondary_sbi_code1_text string `xml:"secondary_sbi_code1_text,omitempty"`

	Secondary_sbi_code2_text string `xml:"secondary_sbi_code2_text,omitempty"`

	Personnel int32 `xml:"personnel,omitempty"`

	Class_personnel int32 `xml:"class_personnel,omitempty"`

	Personnel_fulltime int32 `xml:"personnel_fulltime,omitempty"`

	Class_personnel_fulltime int32 `xml:"class_personnel_fulltime,omitempty"`

	Personnel_reference_date *KvkDate `xml:"personnel_reference_date,omitempty"`

	Indication_import bool `xml:"indication_import,omitempty"`

	Indication_export bool `xml:"indication_export,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`

	Indication_non_mailing bool `xml:"indication_non_mailing,omitempty"`

	Indication_bankruptcy bool `xml:"indication_bankruptcy,omitempty"`

	Indication_dip bool `xml:"indication_dip,omitempty"`

	Authorized_share_capital int64 `xml:"authorized_share_capital,omitempty"`

	Authorized_share_capital_currency string `xml:"authorized_share_capital_currency,omitempty"`

	Paid_up_share_capital int64 `xml:"paid_up_share_capital,omitempty"`

	Paid_up_share_capital_currency string `xml:"paid_up_share_capital_currency,omitempty"`

	Issued_share_capital int64 `xml:"issued_share_capital,omitempty"`

	Issued_share_capital_currency string `xml:"issued_share_capital_currency,omitempty"`

	Founding_date *KvkDate `xml:"founding_date,omitempty"`

	Continuation_date *KvkDate `xml:"continuation_date,omitempty"`

	Establishment_date *KvkDate `xml:"establishment_date,omitempty"`
}

type KvkExtractDocument struct {
	Document []byte `xml:"document,omitempty"`

	Document_date time.Time `xml:"document_date,omitempty"`
}

type KvkReference struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	Establishment_city string `xml:"establishment_city,omitempty"`

	Establishment_street string `xml:"establishment_street,omitempty"`

	Correspondence_city string `xml:"correspondence_city,omitempty"`

	Correspondence_street string `xml:"correspondence_street,omitempty"`

	Indication_economically_active bool `xml:"indication_economically_active,omitempty"`
}

type KvkReferenceArray struct {
	Item []*KvkReference `xml:"item,omitempty"`
}

type KvkReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *KvkReferenceArray `xml:"results,omitempty"`
}

type KvkUpdateReference struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Date_last_update time.Time `xml:"date_last_update,omitempty"`
}

type KvkUpdateReferenceArray struct {
	Item []*KvkUpdateReference `xml:"item,omitempty"`
}

type KvkUpdateReferencePagedResult struct {
	Paging *ResultInfo `xml:"paging,omitempty"`

	Results *KvkUpdateReferenceArray `xml:"results,omitempty"`
}

type KvkDate struct {
	Year int32 `xml:"Year,omitempty"`

	Month int32 `xml:"Month,omitempty"`

	Day int32 `xml:"Day,omitempty"`
}

type RDCoordinatesArray struct {
	Item []*RDCoordinates `xml:"item,omitempty"`
}

type LatLonCoordinatesArray struct {
	Item []*LatLonCoordinates `xml:"item,omitempty"`
}

type MetaServiceReference struct {
	Name string `xml:"name,omitempty"`

	Category string `xml:"category,omitempty"`

	Description string `xml:"description,omitempty"`

	Documentation_url string `xml:"documentation_url,omitempty"`

	Status string `xml:"status,omitempty"`
}

type MetaServiceReferenceArray struct {
	Item []*MetaServiceReference `xml:"item,omitempty"`
}

type MetaServiceDefinition struct {
	Name string `xml:"name,omitempty"`

	Category string `xml:"category,omitempty"`

	Description string `xml:"description,omitempty"`

	Documentation_url string `xml:"documentation_url,omitempty"`

	Methods *MetaMethodReferenceArray `xml:"methods,omitempty"`

	Protocols *MetaProtocolDefinitionArray `xml:"protocols,omitempty"`

	Status string `xml:"status,omitempty"`
}

type MetaServiceDefinitionArray struct {
	Item []*MetaServiceDefinition `xml:"item,omitempty"`
}

type MetaMethodReference struct {
	Name string `xml:"name,omitempty"`

	Deprecated bool `xml:"deprecated,omitempty"`

	Latest bool `xml:"latest,omitempty"`

	Description string `xml:"description,omitempty"`

	Documentation_url string `xml:"documentation_url,omitempty"`
}

type MetaMethodReferenceArray struct {
	Item []*MetaMethodReference `xml:"item,omitempty"`
}

type MetaMethodDefinition struct {
	Name string `xml:"name,omitempty"`

	Service_name string `xml:"service_name,omitempty"`

	Deprecated bool `xml:"deprecated,omitempty"`

	Latest bool `xml:"latest,omitempty"`

	Description string `xml:"description,omitempty"`

	Documentation string `xml:"documentation,omitempty"`

	Documentation_url string `xml:"documentation_url,omitempty"`

	Input_elements *MetaElementDefinitionArray `xml:"input_elements,omitempty"`

	Output_elements *MetaElementDefinitionArray `xml:"output_elements,omitempty"`

	Related_methods *MetaMethodReferenceArray `xml:"related_methods,omitempty"`
}

type MetaMethodDefinitionArray struct {
	Item []*MetaMethodDefinition `xml:"item,omitempty"`
}

type MetaElementDefinition struct {
	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Description string `xml:"description,omitempty"`

	Documentation_url string `xml:"documentation_url,omitempty"`

	Required bool `xml:"required,omitempty"`

	Validations *MetaValidationDefinitionArray `xml:"validations,omitempty"`

	Default_value string `xml:"default_value,omitempty"`
}

type MetaElementDefinitionArray struct {
	Item []*MetaElementDefinition `xml:"item,omitempty"`
}

type MetaValidationDefinition struct {
	Type_ string `xml:"type,omitempty"`

	Value string `xml:"value,omitempty"`
}

type MetaValidationDefinitionArray struct {
	Item []*MetaValidationDefinition `xml:"item,omitempty"`
}

type MetaProtocolDefinition struct {
	Documentation_url string `xml:"documentation_url,omitempty"`

	Location string `xml:"location,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type MetaProtocolDefinitionArray struct {
	Item []*MetaProtocolDefinition `xml:"item,omitempty"`
}

type NbwoWaarde struct {
	Waarde int32 `xml:"waarde,omitempty"`

	Betrouwbaarheid string `xml:"betrouwbaarheid,omitempty"`

	Datum_waardebepaling time.Time `xml:"datum_waardebepaling,omitempty"`

	Cultuurcode string `xml:"cultuurcode,omitempty"`

	Woningtype string `xml:"woningtype,omitempty"`

	Prijspeil_datum time.Time `xml:"prijspeil_datum,omitempty"`

	Straat string `xml:"straat,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Plaats string `xml:"plaats,omitempty"`

	Oppervlak int32 `xml:"oppervlak,omitempty"`

	Inhoud int32 `xml:"inhoud,omitempty"`

	Bouwjaar int32 `xml:"bouwjaar,omitempty"`

	Grootte int32 `xml:"grootte,omitempty"`

	Garage bool `xml:"garage,omitempty"`

	Tuin bool `xml:"tuin,omitempty"`

	Monument bool `xml:"monument,omitempty"`

	Schuur bool `xml:"schuur,omitempty"`

	Postcode_model_resultaat *NbwoPostcodeStraatModelResultaat `xml:"postcode_model_resultaat,omitempty"`

	Straat_model_resultaat *NbwoPostcodeStraatModelResultaat `xml:"straat_model_resultaat,omitempty"`

	Vorige_verkoop_model_resultaat *NbwoVorigeVerkoopModelResultaat `xml:"vorige_verkoop_model_resultaat,omitempty"`

	Kenmerken_model_resultaat *NbwoKenmerkenModelResultaat `xml:"kenmerken_model_resultaat,omitempty"`
}

type NbwoPostcodeStraatModelResultaat struct {
	Waarde int32 `xml:"waarde,omitempty"`

	Aantal int32 `xml:"aantal,omitempty"`

	Standaard_deviatie float64 `xml:"standaard_deviatie,omitempty"`

	Straat string `xml:"straat,omitempty"`

	Plaats string `xml:"plaats,omitempty"`

	Datum_van time.Time `xml:"datum_van,omitempty"`

	Waardeverdelingen *NbwoWaardeVerdelingArray `xml:"waardeverdelingen,omitempty"`
}

type NbwoWaardeVerdeling struct {
	Koopsom_van int32 `xml:"koopsom_van,omitempty"`

	Koopsom_tot int32 `xml:"koopsom_tot,omitempty"`

	Aantal int32 `xml:"aantal,omitempty"`
}

type NbwoWaardeVerdelingArray struct {
	Item []*NbwoWaardeVerdeling `xml:"item,omitempty"`
}

type NbwoVorigeVerkoopModelResultaat struct {
	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Vorige_verkoop_datum time.Time `xml:"vorige_verkoop_datum,omitempty"`

	Koopsom int32 `xml:"koopsom,omitempty"`

	Vorige_verkoop_waarde int32 `xml:"vorige_verkoop_waarde,omitempty"`
}

type NbwoKenmerkenModelResultaat struct {
	Waarde int32 `xml:"waarde,omitempty"`

	Betrouwbaarheid float64 `xml:"betrouwbaarheid,omitempty"`

	Aantal int32 `xml:"aantal,omitempty"`
}

type RiskPersonCompanyReport struct {
	Person *RiskPerson `xml:"person,omitempty"`

	Conclusion *RiskResult `xml:"conclusion,omitempty"`

	Report_data *RiskCompanyReport `xml:"report_data,omitempty"`
}

type RiskPerson struct {
	Gender string `xml:"gender,omitempty"`

	First_name string `xml:"first_name,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Birth_date time.Time `xml:"birth_date,omitempty"`

	Address *RiskAddress `xml:"address,omitempty"`
}

type RiskCompanyReport struct {
	Companies *RiskCompanyArray `xml:"companies,omitempty"`

	Insolvencies *RiskInsolvencyArray `xml:"insolvencies,omitempty"`

	Legal_restraint *RiskLegalRestraint `xml:"legal_restraint,omitempty"`
}

type RiskCompany struct {
	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Status string `xml:"status,omitempty"`

	Trade_names *StringArray `xml:"trade_names,omitempty"`

	Addresses *RiskCompanyAddressArray `xml:"addresses,omitempty"`

	Function_title string `xml:"function_title,omitempty"`

	Function_authorization string `xml:"function_authorization,omitempty"`

	Function_end_date time.Time `xml:"function_end_date,omitempty"`

	Insolvencies *RiskInsolvencyArray `xml:"insolvencies,omitempty"`

	Mutation_date time.Time `xml:"mutation_date,omitempty"`
}

type RiskCompanyArray struct {
	Item []*RiskCompany `xml:"item,omitempty"`
}

type RiskInsolvency struct {
	Registration *RiskInsolvencyRegistration `xml:"registration,omitempty"`

	Publications *RiskInsolvencyPublicationArray `xml:"publications,omitempty"`

	Curators *RiskInsolvencyCuratorArray `xml:"curators,omitempty"`

	Judge *RiskInsolvencyJudge `xml:"judge,omitempty"`
}

type RiskInsolvencyArray struct {
	Item []*RiskInsolvency `xml:"item,omitempty"`
}

type RiskInsolvencyRegistration struct {
	Case_number string `xml:"case_number,omitempty"`

	Previous_number string `xml:"previous_number,omitempty"`

	Court string `xml:"court,omitempty"`

	Jurisdiction string `xml:"jurisdiction,omitempty"`

	District string `xml:"district,omitempty"`

	Deptor string `xml:"deptor,omitempty"`

	Partner string `xml:"partner,omitempty"`
}

type RiskInsolvencyPublication struct {
	Number string `xml:"number,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Verdict_date time.Time `xml:"verdict_date,omitempty"`
}

type RiskInsolvencyPublicationArray struct {
	Item []*RiskInsolvencyPublication `xml:"item,omitempty"`
}

type RiskInsolvencyCurator struct {
	Titles string `xml:"titles,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Address *RiskAddress `xml:"address,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Fax_number string `xml:"fax_number,omitempty"`

	Email_address string `xml:"email_address,omitempty"`

	Mutation_date time.Time `xml:"mutation_date,omitempty"`
}

type RiskInsolvencyCuratorArray struct {
	Item []*RiskInsolvencyCurator `xml:"item,omitempty"`
}

type RiskInsolvencyJudge struct {
	Titles string `xml:"titles,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`
}

type RiskLegalRestraint struct {
	Verdict_number string `xml:"verdict_number,omitempty"`

	Person *RiskLegalRestraintPerson `xml:"person,omitempty"`

	Registry *RiskLegalRestraintRegistry `xml:"registry,omitempty"`

	Decisions *RiskLegalRestraintDecisionArray `xml:"decisions,omitempty"`
}

type RiskLegalRestraintPerson struct {
	First_name string `xml:"first_name,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Birth_date time.Time `xml:"birth_date,omitempty"`
}

type RiskLegalRestraintRegistry struct {
	Address *RiskAddress `xml:"address,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Email_address string `xml:"email_address,omitempty"`
}

type RiskLegalRestraintDecision struct {
	Number string `xml:"number,omitempty"`

	Cause string `xml:"cause,omitempty"`

	Date time.Time `xml:"date,omitempty"`

	Court string `xml:"court,omitempty"`

	Jurisdiction string `xml:"jurisdiction,omitempty"`

	District string `xml:"district,omitempty"`

	Administrator *RiskLegalRestraintAdministrator `xml:"administrator,omitempty"`
}

type RiskLegalRestraintDecisionArray struct {
	Item []*RiskLegalRestraintDecision `xml:"item,omitempty"`
}

type RiskLegalRestraintAdministrator struct {
	Role string `xml:"role,omitempty"`

	Name string `xml:"name,omitempty"`

	Address *RiskAddress `xml:"address,omitempty"`
}

type RiskAddress struct {
	Postcode string `xml:"postcode,omitempty"`

	Building string `xml:"building,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	City string `xml:"city,omitempty"`

	Municipality string `xml:"municipality,omitempty"`

	State string `xml:"state,omitempty"`

	Country string `xml:"country,omitempty"`
}

type RiskCompanyAddress struct {
	Type_ string `xml:"type,omitempty"`

	Address *RiskAddress `xml:"address,omitempty"`
}

type RiskCompanyAddressArray struct {
	Item []*RiskCompanyAddress `xml:"item,omitempty"`
}

type RiskResult struct {
	Person_score int32 `xml:"person_score,omitempty"`

	Address_score int32 `xml:"address_score,omitempty"`

	Postcode_score int32 `xml:"postcode_score,omitempty"`
}

type RoutePlannerRoute struct {
	Start *RouteLocation `xml:"start,omitempty"`

	Destination *RouteLocation `xml:"destination,omitempty"`

	Summary *RouteSummary `xml:"summary,omitempty"`

	Description *RoutePartArray `xml:"description,omitempty"`

	Map_ []byte `xml:"map,omitempty"`
}

type RouteLocation struct {
	Address *RouteAddress `xml:"address,omitempty"`

	Coordinates *RouteCoordinates `xml:"coordinates,omitempty"`
}

type RouteAddress struct {
	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Street string `xml:"street,omitempty"`

	City string `xml:"city,omitempty"`

	Country string `xml:"country,omitempty"`

	Address_key string `xml:"address_key,omitempty"`
}

type RouteCoordinates struct {
	Latlong *RouteCoordinateLatLong `xml:"latlong,omitempty"`

	Rd *RouteCoordinateRD `xml:"rd,omitempty"`
}

type RouteCoordinateLatLong struct {
	Latitude float64 `xml:"latitude,omitempty"`

	Longitude float64 `xml:"longitude,omitempty"`
}

type RouteCoordinateRD struct {
	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`
}

type RouteSummary struct {
	Time int32 `xml:"time,omitempty"`

	Distance int32 `xml:"distance,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type RoutePart struct {
	Description string `xml:"description,omitempty"`

	Time int32 `xml:"time,omitempty"`

	Distance int32 `xml:"distance,omitempty"`
}

type RoutePartArray struct {
	Item []*RoutePart `xml:"item,omitempty"`
}

type RouteInfo struct {
	Time int32 `xml:"time,omitempty"`

	Distance int32 `xml:"distance,omitempty"`
}

type RouteCoordinatesRD struct {
	Number int32 `xml:"number,omitempty"`

	Coordinates string `xml:"coordinates,omitempty"`
}

type RouteDescriptionCoordinatesRD struct {
	Routecoordinates *RouteCoordinatesRD `xml:"routecoordinates,omitempty"`

	Parts *RoutePartArray `xml:"parts,omitempty"`
}

type RouteCoordinatesLatLon struct {
	Number int32 `xml:"number,omitempty"`

	Coordinates string `xml:"coordinates,omitempty"`
}

type RouteDescriptionCoordinatesLatLon struct {
	Routecoordinates *RouteCoordinatesLatLon `xml:"routecoordinates,omitempty"`

	Parts *RoutePartArray `xml:"parts,omitempty"`
}

type SepaBankAccountData struct {
	Bban string `xml:"bban,omitempty"`

	Iban string `xml:"iban,omitempty"`

	Bic string `xml:"bic,omitempty"`
}

type SepaBankAccountIbanData struct {
	Iban string `xml:"iban,omitempty"`

	Iban_details *SepaIbanDetails `xml:"iban_details,omitempty"`

	Bank_details *SepaBankDetails `xml:"bank_details,omitempty"`

	Is_sepa_member bool `xml:"is_sepa_member,omitempty"`
}

type SepaIbanDetails struct {
	Check_digits string `xml:"check_digits,omitempty"`

	Country_code string `xml:"country_code,omitempty"`

	Bban string `xml:"bban,omitempty"`

	Bban_details *BbanDetails `xml:"bban_details,omitempty"`
}

type BbanDetails struct {
	Bban_bank_id string `xml:"bban_bank_id,omitempty"`

	Bban_branch_id string `xml:"bban_branch_id,omitempty"`

	Bban_account string `xml:"bban_account,omitempty"`

	Bban_account_prefix string `xml:"bban_account_prefix,omitempty"`

	Bban_account_type string `xml:"bban_account_type,omitempty"`

	Bban_checksum string `xml:"bban_checksum,omitempty"`

	Bban_currency string `xml:"bban_currency,omitempty"`

	Bban_holder_kennitala string `xml:"bban_holder_kennitala,omitempty"`

	Bban_owner_number string `xml:"bban_owner_number,omitempty"`

	Bban_sort_code string `xml:"bban_sort_code,omitempty"`

	Bban_zero string `xml:"bban_zero,omitempty"`
}

type SepaBankDetails struct {
	Bank_code string `xml:"bank_code,omitempty"`

	Bank_name string `xml:"bank_name,omitempty"`

	Bic string `xml:"bic,omitempty"`
}

type SepaInternationalBankAccountNumberFormatValidationReport struct {
	Iban string `xml:"iban,omitempty"`

	Valid bool `xml:"valid,omitempty"`
}

type SepaMatchAccountHolderResult struct {
	Iban string `xml:"iban,omitempty"`

	Iban_found bool `xml:"iban_found,omitempty"`

	Bank_account_details *SepaMatchBankAccountDetails `xml:"bank_account_details,omitempty"`

	Bank_account_holder *SepaMatchBankAccountHolder `xml:"bank_account_holder,omitempty"`

	Name_matching_result *SepaMatchNameMatchingResult `xml:"name_matching_result,omitempty"`
}

type SepaMatchBankAccountDetails struct {
	Foreign bool `xml:"foreign,omitempty"`

	Country_code string `xml:"country_code,omitempty"`

	Status string `xml:"status,omitempty"`
}

type SepaMatchBankAccountHolder struct {
	Type_ string `xml:"type,omitempty"`

	Number_of_account_holders int32 `xml:"number_of_account_holders,omitempty"`

	Joint_account bool `xml:"joint_account,omitempty"`

	Municipality string `xml:"municipality,omitempty"`
}

type SepaMatchNameMatchingResult struct {
	Type_ string `xml:"type,omitempty"`

	Name_suggestion string `xml:"name_suggestion,omitempty"`
}

type VatValidation struct {
	Vat_number string `xml:"vat_number,omitempty"`

	Valid bool `xml:"valid,omitempty"`
}

type VatProxyViesCheckVatResponse struct {
	Country_code string `xml:"country_code,omitempty"`

	Vat_number string `xml:"vat_number,omitempty"`

	Request_date time.Time `xml:"request_date,omitempty"`

	Valid bool `xml:"valid,omitempty"`

	Name string `xml:"name,omitempty"`

	Address string `xml:"address,omitempty"`
}

type AccountGetCreationTokenRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountGetCreationToken"`

	Return_url string `xml:"return_url,omitempty"`
}

type AccountGetCreationTokenResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountGetCreationTokenResponse"`

	Out *AccountCreationToken `xml:"out,omitempty"`
}

type AccountGetCreationStatusRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountGetCreationStatus"`

	Token string `xml:"token,omitempty"`
}

type AccountGetCreationStatusResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountGetCreationStatusResponse"`

	Accountid int32 `xml:"accountid,omitempty"`
}

type AccountGetOrderTokenRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountGetOrderToken"`

	Accountid int32 `xml:"accountid,omitempty"`

	Return_url string `xml:"return_url,omitempty"`
}

type AccountGetOrderTokenResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountGetOrderTokenResponse"`

	Out *AccountOrderToken `xml:"out,omitempty"`
}

type UserSessionRemoveRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userSessionRemove"`

	Userid int32 `xml:"userid,omitempty"`

	Reactid string `xml:"reactid,omitempty"`
}

type UserSessionRemoveResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userSessionRemoveResponse"`
}

type UserSessionListRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userSessionList"`

	Userid int32 `xml:"userid,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type UserSessionListResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userSessionListResponse"`

	Out *SessionPagedResult `xml:"out,omitempty"`
}

type UserViewBalanceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userViewBalance"`

	Userid int32 `xml:"userid,omitempty"`
}

type UserViewBalanceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userViewBalanceResponse"`

	Balance float32 `xml:"balance,omitempty"`
}

type UserEditBalanceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditBalance"`

	Userid int32 `xml:"userid,omitempty"`

	Balance float32 `xml:"balance,omitempty"`
}

type UserEditBalanceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditBalanceResponse"`
}

type AccountViewBalanceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountViewBalance"`

	Accountid int32 `xml:"accountid,omitempty"`
}

type AccountViewBalanceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountViewBalanceResponse"`

	Balance float32 `xml:"balance,omitempty"`
}

type UserViewV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userViewV2"`

	Userid int32 `xml:"userid,omitempty"`
}

type UserViewV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userViewV2Response"`

	Out *UserV2 `xml:"out,omitempty"`
}

type UserEditV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditV2"`

	Userid int32 `xml:"userid,omitempty"`

	Email string `xml:"email,omitempty"`

	Address string `xml:"address,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Password string `xml:"password,omitempty"`
}

type UserEditV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditV2Response"`
}

type UserEditExtendedV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditExtendedV2"`

	Userid int32 `xml:"userid,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	Password string `xml:"password,omitempty"`

	Email string `xml:"email,omitempty"`

	Companyname string `xml:"companyname,omitempty"`

	Address string `xml:"address,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Clientcode string `xml:"clientcode,omitempty"`

	Comments string `xml:"comments,omitempty"`

	Accountid int32 `xml:"accountid,omitempty"`

	Balancethreshold float32 `xml:"balancethreshold,omitempty"`

	Notificationrecipients string `xml:"notificationrecipients,omitempty"`
}

type UserEditExtendedV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditExtendedV2Response"`
}

type UserCreateV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userCreateV2"`

	Accountid int32 `xml:"accountid,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	Password string `xml:"password,omitempty"`

	Usergroups *IntArray `xml:"usergroups,omitempty"`

	Email string `xml:"email,omitempty"`

	Companyname string `xml:"companyname,omitempty"`

	Address string `xml:"address,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Clientcode string `xml:"clientcode,omitempty"`

	Comments string `xml:"comments,omitempty"`
}

type UserCreateV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userCreateV2Response"`

	Id int32 `xml:"id,omitempty"`

	Nickname string `xml:"nickname,omitempty"`

	Password string `xml:"password,omitempty"`
}

type CreateTestUserRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ createTestUser"`

	Application string `xml:"application,omitempty"`

	Email string `xml:"email,omitempty"`

	Companyname string `xml:"companyname,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Telephone string `xml:"telephone,omitempty"`
}

type CreateTestUserResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ createTestUserResponse"`

	User_credentials *UserCredentials `xml:"user_credentials,omitempty"`
}

type UserChangePasswordRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userChangePassword"`

	Userid int32 `xml:"userid,omitempty"`

	Old_password string `xml:"old_password,omitempty"`

	New_password string `xml:"new_password,omitempty"`
}

type UserChangePasswordResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userChangePasswordResponse"`
}

type UserRemoveRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userRemove"`

	Userid int32 `xml:"userid,omitempty"`
}

type UserRemoveResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userRemoveResponse"`
}

type UserNotifyRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userNotify"`

	Userid int32 `xml:"userid,omitempty"`

	Password string `xml:"password,omitempty"`
}

type UserNotifyResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userNotifyResponse"`

	Out string `xml:"out,omitempty"`
}

type UserListAssignableGroupsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userListAssignableGroups"`

	Userid int32 `xml:"userid,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type UserListAssignableGroupsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userListAssignableGroupsResponse"`

	Out *UserGroupPagedResult `xml:"out,omitempty"`
}

type UserAddGroupRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userAddGroup"`

	Userid int32 `xml:"userid,omitempty"`

	Usergroupid int32 `xml:"usergroupid,omitempty"`
}

type UserAddGroupResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userAddGroupResponse"`
}

type UserRemoveGroupRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userRemoveGroup"`

	Userid int32 `xml:"userid,omitempty"`

	Usergroupid int32 `xml:"usergroupid,omitempty"`
}

type UserRemoveGroupResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userRemoveGroupResponse"`
}

type AccountViewV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountViewV2"`

	Accountid int32 `xml:"accountid,omitempty"`
}

type AccountViewV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountViewV2Response"`

	Out *AccountV2 `xml:"out,omitempty"`
}

type AccountEditV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountEditV2"`

	Accountid int32 `xml:"accountid,omitempty"`

	Address string `xml:"address,omitempty"`

	Contactname string `xml:"contactname,omitempty"`

	Contactemail string `xml:"contactemail,omitempty"`

	Telephone string `xml:"telephone,omitempty"`

	Fax string `xml:"fax,omitempty"`

	Description string `xml:"description,omitempty"`

	Balancethreshold float32 `xml:"balancethreshold,omitempty"`
}

type AccountEditV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountEditV2Response"`
}

type AccountUserListV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountUserListV2"`

	Accountid int32 `xml:"accountid,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AccountUserListV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountUserListV2Response"`

	Out *UserV2PagedResult `xml:"out,omitempty"`
}

type AccountUserSearchV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountUserSearchV2"`

	Accountid int32 `xml:"accountid,omitempty"`

	Phrase string `xml:"phrase,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AccountUserSearchV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountUserSearchV2Response"`

	Out *UserV2PagedResult `xml:"out,omitempty"`
}

type AccountEditHostRestrictionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountEditHostRestrictions"`

	Accountid int32 `xml:"accountid,omitempty"`

	Restrictions string `xml:"restrictions,omitempty"`
}

type AccountEditHostRestrictionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountEditHostRestrictionsResponse"`
}

type AccountViewHostRestrictionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountViewHostRestrictions"`

	Accountid int32 `xml:"accountid,omitempty"`
}

type AccountViewHostRestrictionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ accountViewHostRestrictionsResponse"`

	Out string `xml:"out,omitempty"`
}

type UserEditHostRestrictionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditHostRestrictions"`

	Userid int32 `xml:"userid,omitempty"`

	Restrictions string `xml:"restrictions,omitempty"`
}

type UserEditHostRestrictionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userEditHostRestrictionsResponse"`
}

type UserViewHostRestrictionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userViewHostRestrictions"`

	Userid int32 `xml:"userid,omitempty"`
}

type UserViewHostRestrictionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ userViewHostRestrictionsResponse"`

	Out string `xml:"out,omitempty"`
}

type AddressReeksPostcodeSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksPostcodeSearch"`

	Address string `xml:"address,omitempty"`
}

type AddressReeksPostcodeSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksPostcodeSearchResponse"`

	Out *PCReeks `xml:"out,omitempty"`
}

type AddressReeksAddressSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksAddressSearch"`

	Address string `xml:"address,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressReeksAddressSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksAddressSearchResponse"`

	Out *RangeAddressPagedResult `xml:"out,omitempty"`
}

type AddressReeksFullParameterSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksFullParameterSearch"`

	Province string `xml:"province,omitempty"`

	District string `xml:"district,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	HouseNo int32 `xml:"houseNo,omitempty"`

	HouseNoAddition string `xml:"houseNoAddition,omitempty"`

	Nbcode string `xml:"nbcode,omitempty"`

	Lettercombination string `xml:"lettercombination,omitempty"`

	Addresstype string `xml:"addresstype,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressReeksFullParameterSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksFullParameterSearchResponse"`

	Out *PCReeksSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressReeksParameterSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksParameterSearch"`

	Province string `xml:"province,omitempty"`

	District string `xml:"district,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	HouseNo int32 `xml:"houseNo,omitempty"`

	HouseNoAddition string `xml:"houseNoAddition,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressReeksParameterSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksParameterSearchResponse"`

	Out *PCReeksSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressPerceelPhraseSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelPhraseSearch"`

	Address string `xml:"address,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressPerceelPhraseSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelPhraseSearchResponse"`

	Out *PerceelSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressPerceelFullParameterSearchV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelFullParameterSearchV2"`

	Province string `xml:"province,omitempty"`

	District string `xml:"district,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	HouseNo int32 `xml:"houseNo,omitempty"`

	HouseNoAddition string `xml:"houseNoAddition,omitempty"`

	Nbcode string `xml:"nbcode,omitempty"`

	Lettercombination string `xml:"lettercombination,omitempty"`

	Addresstype string `xml:"addresstype,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressPerceelFullParameterSearchV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelFullParameterSearchV2Response"`

	Out *PerceelSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressProvinceListNeighborhoodsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceListNeighborhoods"`

	Name string `xml:"name,omitempty"`

	Postbus bool `xml:"postbus,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressProvinceListNeighborhoodsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceListNeighborhoodsResponse"`

	Out *NeighborhoodPagedResult `xml:"out,omitempty"`
}

type AddressProvinceListDistrictsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceListDistricts"`

	Name string `xml:"name,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressProvinceListDistrictsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceListDistrictsResponse"`

	Out *DistrictPagedResult `xml:"out,omitempty"`
}

type AddressProvinceListRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceList"`

	Page int32 `xml:"page,omitempty"`
}

type AddressProvinceListResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceListResponse"`

	Out *ProvincePagedResult `xml:"out,omitempty"`
}

type AddressProvinceSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceSearch"`

	Name string `xml:"name,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressProvinceSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressProvinceSearchResponse"`

	Out *ProvincePagedResult `xml:"out,omitempty"`
}

type AddressDistrictSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressDistrictSearch"`

	Name string `xml:"name,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressDistrictSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressDistrictSearchResponse"`

	Out *DistrictPagedResult `xml:"out,omitempty"`
}

type AddressDistrictListCitiesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressDistrictListCities"`

	Name string `xml:"name,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressDistrictListCitiesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressDistrictListCitiesResponse"`

	Out *CityPagedResult `xml:"out,omitempty"`
}

type AddressDistrictListNeighborhoodsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressDistrictListNeighborhoods"`

	Name string `xml:"name,omitempty"`

	Postbus bool `xml:"postbus,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressDistrictListNeighborhoodsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressDistrictListNeighborhoodsResponse"`

	Out *NeighborhoodPagedResult `xml:"out,omitempty"`
}

type AddressCitySearchV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressCitySearchV2"`

	Name string `xml:"name,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressCitySearchV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressCitySearchV2Response"`

	Out *CityV2PagedResult `xml:"out,omitempty"`
}

type AddressCityListNeighborhoodsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressCityListNeighborhoods"`

	Name string `xml:"name,omitempty"`

	Postbus bool `xml:"postbus,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressCityListNeighborhoodsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressCityListNeighborhoodsResponse"`

	Out *NeighborhoodPagedResult `xml:"out,omitempty"`
}

type AddressPerceelFullParameterSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelFullParameterSearch"`

	Province string `xml:"province,omitempty"`

	District string `xml:"district,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	HouseNo int32 `xml:"houseNo,omitempty"`

	HouseNoAddition string `xml:"houseNoAddition,omitempty"`

	Nbcode string `xml:"nbcode,omitempty"`

	Lettercombination string `xml:"lettercombination,omitempty"`

	Addresstype string `xml:"addresstype,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressPerceelFullParameterSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelFullParameterSearchResponse"`

	Out *PerceelSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressPerceelParameterSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelParameterSearch"`

	Province string `xml:"province,omitempty"`

	District string `xml:"district,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	HouseNo int32 `xml:"houseNo,omitempty"`

	HouseNoAddition string `xml:"houseNoAddition,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressPerceelParameterSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressPerceelParameterSearchResponse"`

	Out *PerceelSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressReeksPhraseSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksPhraseSearch"`

	Address string `xml:"address,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressReeksPhraseSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressReeksPhraseSearchResponse"`

	Out *PCReeksSearchPartsPagedResult `xml:"out,omitempty"`
}

type AddressNeighborhoodNameRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressNeighborhoodName"`

	Nbcode string `xml:"nbcode,omitempty"`
}

type AddressNeighborhoodNameResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressNeighborhoodNameResponse"`

	Name string `xml:"name,omitempty"`
}

type AddressNeighborhoodPhraseSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressNeighborhoodPhraseSearch"`

	Name string `xml:"name,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AddressNeighborhoodPhraseSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ addressNeighborhoodPhraseSearchResponse"`

	Out *NeighborhoodNamePagedResult `xml:"out,omitempty"`
}

type AreaCodeLookupRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ areaCodeLookup"`

	Neighborhoodcode string `xml:"neighborhoodcode,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AreaCodeLookupResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ areaCodeLookupResponse"`

	Out *AreaCodePagedResult `xml:"out,omitempty"`
}

type AreaCodeToNeighborhoodcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ areaCodeToNeighborhoodcode"`

	Areacode string `xml:"areacode,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type AreaCodeToNeighborhoodcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ areaCodeToNeighborhoodcodeResponse"`

	Out *NeighborhoodPagedResult `xml:"out,omitempty"`
}

type AreaCodePostcodeLookupRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ areaCodePostcodeLookup"`

	Postcode string `xml:"postcode,omitempty"`
}

type AreaCodePostcodeLookupResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ areaCodePostcodeLookupResponse"`

	Out *AreaCodeArray `xml:"out,omitempty"`
}

type LoginRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ login"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`
}

type LoginResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ loginResponse"`

	Reactid string `xml:"reactid,omitempty"`
}

type LogoutRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ logout"`
}

type LogoutResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ logoutResponse"`
}

type BelgianBusinessGetDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ belgianBusinessGetDossier"`

	Vat_number string `xml:"vat_number,omitempty"`
}

type BelgianBusinessGetDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ belgianBusinessGetDossierResponse"`

	Out *BelgianBusinessDossier `xml:"out,omitempty"`
}

type BovagGetMemberByBovagIdRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ bovagGetMemberByBovagId"`

	Bovag_id string `xml:"bovag_id,omitempty"`
}

type BovagGetMemberByBovagIdResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ bovagGetMemberByBovagIdResponse"`

	Member *BovagMember `xml:"member,omitempty"`
}

type BovagGetMemberByDutchBusinessRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ bovagGetMemberByDutchBusiness"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type BovagGetMemberByDutchBusinessResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ bovagGetMemberByDutchBusinessResponse"`

	Member *BovagMember `xml:"member,omitempty"`
}

type BusinessGetEstablishmentNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetEstablishmentNumber"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`
}

type BusinessGetEstablishmentNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetEstablishmentNumberResponse"`

	Out string `xml:"out,omitempty"`
}

type BusinessGetBIKDescriptionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetBIKDescription"`

	Bikcode string `xml:"bikcode,omitempty"`
}

type BusinessGetBIKDescriptionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetBIKDescriptionResponse"`

	Out *BusinessBIKCodeInfo `xml:"out,omitempty"`
}

type BusinessGetSBIDescriptionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetSBIDescription"`

	Sbicode string `xml:"sbicode,omitempty"`
}

type BusinessGetSBIDescriptionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetSBIDescriptionResponse"`

	Out *BusinessSBICodeInfo `xml:"out,omitempty"`
}

type BusinessBIKToSBIRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessBIKToSBI"`

	Bikcode string `xml:"bikcode,omitempty"`
}

type BusinessBIKToSBIResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessBIKToSBIResponse"`

	Out *BusinessSBICodeArray `xml:"out,omitempty"`
}

type BusinessSBIToBIKRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSBIToBIK"`

	Sbicode string `xml:"sbicode,omitempty"`
}

type BusinessSBIToBIKResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSBIToBIKResponse"`

	Out *BusinessBIKCodeArray `xml:"out,omitempty"`
}

type BusinessGetDossierV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetDossierV3"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessGetDossierV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetDossierV3Response"`

	Out *BusinessDossierV3PagedResult `xml:"out,omitempty"`
}

type BusinessGetDossierExtendedRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetDossierExtended"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`
}

type BusinessGetDossierExtendedResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetDossierExtendedResponse"`

	Out *BusinessDossierExtended `xml:"out,omitempty"`
}

type BusinessSearchDossierNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchDossierNumber"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchDossierNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchDossierNumberResponse"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessSearchPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchPostcode"`

	Nbcode string `xml:"nbcode,omitempty"`

	Lettercomb string `xml:"lettercomb,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchPostcodeResponse"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessSearchAddressRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchAddress"`

	Streetname string `xml:"streetname,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Cityname string `xml:"cityname,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchAddressResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchAddressResponse"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessSearchNameRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchName"`

	Tradename string `xml:"tradename,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchNameResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchNameResponse"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessSearchParametersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchParameters"`

	Tradename string `xml:"tradename,omitempty"`

	Cityname string `xml:"cityname,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	Nbcode string `xml:"nbcode,omitempty"`

	Lettercomb string `xml:"lettercomb,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Telephoneno string `xml:"telephoneno,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchParametersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchParametersResponse"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessSearchParametersV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchParametersV3"`

	Tradename string `xml:"tradename,omitempty"`

	Cityname string `xml:"cityname,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Telephoneno string `xml:"telephoneno,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchParametersV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchParametersV3Response"`

	Out *BusinessReferenceV3PagedResult `xml:"out,omitempty"`
}

type BusinessSearchSelectionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchSelection"`

	City *StringArray `xml:"city,omitempty"`

	Postcode *StringArray `xml:"postcode,omitempty"`

	Sbi *StringArray `xml:"sbi,omitempty"`

	Primary_sbi_only bool `xml:"primary_sbi_only,omitempty"`

	Legal_forms *IntArray `xml:"legal_forms,omitempty"`

	Employees_min int32 `xml:"employees_min,omitempty"`

	Employees_max int32 `xml:"employees_max,omitempty"`

	Economically_active string `xml:"economically_active,omitempty"`

	Financial_status string `xml:"financial_status,omitempty"`

	Changed_since string `xml:"changed_since,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchSelectionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchSelectionResponse"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessSearchSelectionV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchSelectionV2"`

	City *StringArray `xml:"city,omitempty"`

	Postcode *StringArray `xml:"postcode,omitempty"`

	Sbi *StringArray `xml:"sbi,omitempty"`

	Primary_sbi_only bool `xml:"primary_sbi_only,omitempty"`

	Legal_forms *IntArray `xml:"legal_forms,omitempty"`

	Employees_min int32 `xml:"employees_min,omitempty"`

	Employees_max int32 `xml:"employees_max,omitempty"`

	Economically_active string `xml:"economically_active,omitempty"`

	Financial_status string `xml:"financial_status,omitempty"`

	Changed_since string `xml:"changed_since,omitempty"`

	New_since string `xml:"new_since,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchSelectionV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchSelectionV2Response"`

	Out *BusinessReferencePagedResult `xml:"out,omitempty"`
}

type BusinessGetDossierSBIRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetDossierSBI"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`
}

type BusinessGetDossierSBIResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessGetDossierSBIResponse"`

	Out *BusinessDossierSBI `xml:"out,omitempty"`
}

type BusinessUpdateCheckDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateCheckDossier"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`
}

type BusinessUpdateCheckDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateCheckDossierResponse"`

	Out *BusinessUpdateReference `xml:"out,omitempty"`
}

type BusinessUpdateGetChangedDossiersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateGetChangedDossiers"`

	Changed_since time.Time `xml:"changed_since,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessUpdateGetChangedDossiersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateGetChangedDossiersResponse"`

	Out *BusinessUpdateReferencePagedResult `xml:"out,omitempty"`
}

type BusinessUpdateGetDossiersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateGetDossiers"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessUpdateGetDossiersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateGetDossiersResponse"`

	Out *BusinessUpdateReferencePagedResult `xml:"out,omitempty"`
}

type BusinessUpdateAddDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateAddDossier"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`
}

type BusinessUpdateAddDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateAddDossierResponse"`
}

type BusinessUpdateRemoveDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateRemoveDossier"`

	Dossierno string `xml:"dossierno,omitempty"`

	Subdossierno string `xml:"subdossierno,omitempty"`
}

type BusinessUpdateRemoveDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessUpdateRemoveDossierResponse"`
}

type BusinessSearchParametersV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchParametersV2"`

	Tradename string `xml:"tradename,omitempty"`

	Cityname string `xml:"cityname,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Telephoneno string `xml:"telephoneno,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type BusinessSearchParametersV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ businessSearchParametersV2Response"`

	Out *BusinessReferenceV2PagedResult `xml:"out,omitempty"`
}

type CarRDWCarCheckCodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarCheckCode"`

	License_plate string `xml:"license_plate,omitempty"`

	Code string `xml:"code,omitempty"`
}

type CarRDWCarCheckCodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarCheckCodeResponse"`

	Out *CarCheckCode `xml:"out,omitempty"`
}

type CarRDWCarDataV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataV3"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarRDWCarDataV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataV3Response"`

	Out *CarDataV3Result `xml:"out,omitempty"`
}

type CarRDWCarDataBPV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataBPV2"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarRDWCarDataBPV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataBPV2Response"`

	Out *CarBPV2 `xml:"out,omitempty"`
}

type CarRDWCarDataExtendedRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataExtended"`

	License_plate string `xml:"license_plate,omitempty"`

	Code string `xml:"code,omitempty"`
}

type CarRDWCarDataExtendedResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataExtendedResponse"`

	Out *CarExtended `xml:"out,omitempty"`
}

type CarRDWCarDataPriceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataPrice"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarRDWCarDataPriceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataPriceResponse"`

	Out *CarRDWCarDataPrice `xml:"out,omitempty"`
}

type CarRDWCarDataOptionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataOptions"`

	Car_id string `xml:"car_id,omitempty"`
}

type CarRDWCarDataOptionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataOptionsResponse"`

	Out *CarOptions `xml:"out,omitempty"`
}

type CarVWEMeldcodeCheckRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEMeldcodeCheck"`

	License_plate string `xml:"license_plate,omitempty"`

	Code string `xml:"code,omitempty"`
}

type CarVWEMeldcodeCheckResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEMeldcodeCheckResponse"`

	Out *CarVWEMeldcodeCheck `xml:"out,omitempty"`
}

type CarVWEBasicTypeDataRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEBasicTypeData"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarVWEBasicTypeDataResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEBasicTypeDataResponse"`

	Out *CarVWEBasicTypeData `xml:"out,omitempty"`
}

type CarVWEVersionPriceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEVersionPrice"`

	License_plate string `xml:"license_plate,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`
}

type CarVWEVersionPriceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEVersionPriceResponse"`

	Out *CarVWEVersionPrice `xml:"out,omitempty"`
}

type CarVWEOptionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEOptions"`

	License_plate string `xml:"license_plate,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`
}

type CarVWEOptionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEOptionsResponse"`

	Out *CarVWEOptions `xml:"out,omitempty"`
}

type CarVWEListBrandsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEListBrands"`

	Production_year int32 `xml:"production_year,omitempty"`

	Kind_id int32 `xml:"kind_id,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type CarVWEListBrandsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEListBrandsResponse"`

	Out *CarVWEBrandPagedResult `xml:"out,omitempty"`
}

type CarVWEListModelsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEListModels"`

	Production_year int32 `xml:"production_year,omitempty"`

	Kind_id int32 `xml:"kind_id,omitempty"`

	Brand_id int32 `xml:"brand_id,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type CarVWEListModelsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEListModelsResponse"`

	Out *CarVWEModelPagedResult `xml:"out,omitempty"`
}

type CarVWEListVersionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEListVersions"`

	Production_year int32 `xml:"production_year,omitempty"`

	Kind_id int32 `xml:"kind_id,omitempty"`

	Brand_id int32 `xml:"brand_id,omitempty"`

	Model_id int32 `xml:"model_id,omitempty"`

	Fuel_type_id int32 `xml:"fuel_type_id,omitempty"`

	Body_style_id int32 `xml:"body_style_id,omitempty"`

	Doors int32 `xml:"doors,omitempty"`

	Gear_id int32 `xml:"gear_id,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type CarVWEListVersionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEListVersionsResponse"`

	Out *CarVWEVersionPagedResult `xml:"out,omitempty"`
}

type CarVWEVersionYearDataRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEVersionYearData"`

	Production_year int32 `xml:"production_year,omitempty"`

	Atl_code int32 `xml:"atl_code,omitempty"`
}

type CarVWEVersionYearDataResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEVersionYearDataResponse"`

	Out *CarVWEVersionYearData `xml:"out,omitempty"`
}

type CarVWEPhotosRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEPhotos"`

	Atl_code int32 `xml:"atl_code,omitempty"`
}

type CarVWEPhotosResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carVWEPhotosResponse"`

	Out *CarVWEPhotoArray `xml:"out,omitempty"`
}

type CarATDPriceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carATDPrice"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarATDPriceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carATDPriceResponse"`

	Out *CarATDPrices `xml:"out,omitempty"`
}

type CarRDWCarDataRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarData"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarRDWCarDataResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataResponse"`

	Out *Car `xml:"out,omitempty"`
}

type CarRDWCarDataV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataV2"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarRDWCarDataV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataV2Response"`

	Out *CarV2 `xml:"out,omitempty"`
}

type CarRDWCarDataBPRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataBP"`

	License_plate string `xml:"license_plate,omitempty"`
}

type CarRDWCarDataBPResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ carRDWCarDataBPResponse"`

	Out *CarBP `xml:"out,omitempty"`
}

type ComplianceSearchPersonsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceSearchPersons"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Date_of_birth time.Time `xml:"date_of_birth,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type ComplianceSearchPersonsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceSearchPersonsResponse"`

	Out *CompliancePersonSearchReferencePagedResult `xml:"out,omitempty"`
}

type ComplianceGetPersonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceGetPerson"`

	Compliance_person_id string `xml:"compliance_person_id,omitempty"`
}

type ComplianceGetPersonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceGetPersonResponse"`

	Out *CompliancePerson `xml:"out,omitempty"`
}

type ComplianceSearchBusinessesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceSearchBusinesses"`

	Business_name string `xml:"business_name,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	City string `xml:"city,omitempty"`

	County string `xml:"county,omitempty"`

	Country string `xml:"country,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type ComplianceSearchBusinessesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceSearchBusinessesResponse"`

	Out *ComplianceBusinessSearchReferencePagedResult `xml:"out,omitempty"`
}

type ComplianceGetBusinessRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceGetBusiness"`

	Compliance_business_id string `xml:"compliance_business_id,omitempty"`
}

type ComplianceGetBusinessResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ complianceGetBusinessResponse"`

	Out *ComplianceBusiness `xml:"out,omitempty"`
}

type CreditsafeGetReportFullRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeGetReportFull"`

	Company_id string `xml:"company_id,omitempty"`

	Language string `xml:"language,omitempty"`

	Document string `xml:"document,omitempty"`
}

type CreditsafeGetReportFullResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeGetReportFullResponse"`

	Out *CreditsafeCompanyReportFull `xml:"out,omitempty"`
}

type CreditsafeSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeSearch"`

	Country string `xml:"country,omitempty"`

	Id string `xml:"id,omitempty"`

	Registration_number string `xml:"registration_number,omitempty"`

	Status string `xml:"status,omitempty"`

	Office_type string `xml:"office_type,omitempty"`

	Name string `xml:"name,omitempty"`

	Name_match_type string `xml:"name_match_type,omitempty"`

	Address string `xml:"address,omitempty"`

	Address_match_type string `xml:"address_match_type,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number string `xml:"house_number,omitempty"`

	City string `xml:"city,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	Province string `xml:"province,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type CreditsafeSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeSearchResponse"`

	Out *CreditsafeCompanyPagedResult `xml:"out,omitempty"`
}

type CreditsafeSearchV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeSearchV2"`

	Countries *StringArray `xml:"countries,omitempty"`

	Id string `xml:"id,omitempty"`

	Status string `xml:"status,omitempty"`

	Registration_number string `xml:"registration_number,omitempty"`

	Registration_type string `xml:"registration_type,omitempty"`

	Vat_number string `xml:"vat_number,omitempty"`

	Province string `xml:"province,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Postal_code string `xml:"postal_code,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CreditsafeSearchV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeSearchV2Response"`

	Out *CreditsafeSearchResultV2 `xml:"out,omitempty"`
}

type CreditsafeGetReportFullV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeGetReportFullV2"`

	Company_id string `xml:"company_id,omitempty"`

	Language string `xml:"language,omitempty"`

	Reason string `xml:"reason,omitempty"`
}

type CreditsafeGetReportFullV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeGetReportFullV2Response"`

	Out *CreditsafeCompanyReportFullV2 `xml:"out,omitempty"`
}

type CreditsafeSearchCriteriaRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeSearchCriteria"`

	Countries *StringArray `xml:"countries,omitempty"`
}

type CreditsafeSearchCriteriaResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeSearchCriteriaResponse"`

	Out *CreditsafeSearchCriteria `xml:"out,omitempty"`
}

type CreditsafeMonitorAddCompanyRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeMonitorAddCompany"`

	Company_id string `xml:"company_id,omitempty"`
}

type CreditsafeMonitorAddCompanyResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeMonitorAddCompanyResponse"`
}

type CreditsafeMonitorRemoveCompanyRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeMonitorRemoveCompany"`

	Company_id string `xml:"company_id,omitempty"`
}

type CreditsafeMonitorRemoveCompanyResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeMonitorRemoveCompanyResponse"`
}

type CreditsafeMonitorGetUpdatedCompaniesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeMonitorGetUpdatedCompanies"`

	Changed_since time.Time `xml:"changed_since,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type CreditsafeMonitorGetUpdatedCompaniesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ creditsafeMonitorGetUpdatedCompaniesResponse"`

	Out *CreditsafeCompanyUpdatePagedResult `xml:"out,omitempty"`
}

type DnbSearchReferenceV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbSearchReferenceV2"`

	Name string `xml:"name,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	Houseno string `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Cityname string `xml:"cityname,omitempty"`

	Region string `xml:"region,omitempty"`

	Country string `xml:"country,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DnbSearchReferenceV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbSearchReferenceV2Response"`

	Out *DNBBusinessReferenceV2PagedResult `xml:"out,omitempty"`
}

type DnbGetReferenceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbGetReference"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbGetReferenceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbGetReferenceResponse"`

	Out *DNBBusinessReference `xml:"out,omitempty"`
}

type DnbWorldbaseMarketingRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbWorldbaseMarketing"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbWorldbaseMarketingResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbWorldbaseMarketingResponse"`

	Out *DNBMarketing `xml:"out,omitempty"`
}

type DnbWorldbaseMarketingPlusRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbWorldbaseMarketingPlus"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbWorldbaseMarketingPlusResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbWorldbaseMarketingPlusResponse"`

	Out *DNBMarketingPlusResult `xml:"out,omitempty"`
}

type DnbWorldbaseMarketingPlusLinkageRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbWorldbaseMarketingPlusLinkage"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbWorldbaseMarketingPlusLinkageResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbWorldbaseMarketingPlusLinkageResponse"`

	Out *DNBMarketingPlusLinkageResult `xml:"out,omitempty"`
}

type DnbQuickCheckRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbQuickCheck"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbQuickCheckResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbQuickCheckResponse"`

	Out *DNBQuickCheck `xml:"out,omitempty"`
}

type DnbBusinessVerificationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbBusinessVerification"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbBusinessVerificationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbBusinessVerificationResponse"`

	Out *DNBBusinessVerification `xml:"out,omitempty"`
}

type DnbEnterpriseManagementRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbEnterpriseManagement"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`
}

type DnbEnterpriseManagementResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbEnterpriseManagementResponse"`

	Out *DNBEnterpriseManagement `xml:"out,omitempty"`
}

type DnbSearchReferenceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbSearchReference"`

	Name string `xml:"name,omitempty"`

	Streetname string `xml:"streetname,omitempty"`

	Houseno string `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Cityname string `xml:"cityname,omitempty"`

	Country string `xml:"country,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DnbSearchReferenceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dnbSearchReferenceResponse"`

	Out *DNBBusinessReferencePagedResult `xml:"out,omitempty"`
}

type DriveInfoDistanceLookupRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ driveInfoDistanceLookup"`

	Nbcodefrom string `xml:"nbcodefrom,omitempty"`

	Nbcodeto string `xml:"nbcodeto,omitempty"`
}

type DriveInfoDistanceLookupResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ driveInfoDistanceLookupResponse"`

	Out *DriveInfo `xml:"out,omitempty"`
}

type DriveInfoTimeLookupRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ driveInfoTimeLookup"`

	Nbcodefrom string `xml:"nbcodefrom,omitempty"`

	Nbcodeto string `xml:"nbcodeto,omitempty"`
}

type DriveInfoTimeLookupResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ driveInfoTimeLookupResponse"`

	Out *DriveInfo `xml:"out,omitempty"`
}

type DutchAddressRangePostcodeSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchAddressRangePostcodeSearch"`

	Address string `xml:"address,omitempty"`
}

type DutchAddressRangePostcodeSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchAddressRangePostcodeSearchResponse"`

	Out *DutchAddressPostcodeRange `xml:"out,omitempty"`
}

type DutchBusinessGetDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type DutchBusinessGetDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierResponse"`

	Out *DutchBusinessDossier `xml:"out,omitempty"`
}

type DutchBusinessGetDossierV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierV2"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type DutchBusinessGetDossierV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierV2Response"`

	Out *DutchBusinessDossierV2 `xml:"out,omitempty"`
}

type DutchBusinessGetDossierV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierV3"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type DutchBusinessGetDossierV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierV3Response"`

	Out *DutchBusinessDossierV3 `xml:"out,omitempty"`
}

type DutchBusinessGetSBIRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetSBI"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Language string `xml:"language,omitempty"`
}

type DutchBusinessGetSBIResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetSBIResponse"`

	Out *DutchBusinessSBICollection `xml:"out,omitempty"`
}

type DutchBusinessGetVatNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetVatNumber"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetVatNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetVatNumberResponse"`

	Out *DutchBusinessVatNumber `xml:"out,omitempty"`
}

type DutchBusinessGetPositionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetPositions"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetPositionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetPositionsResponse"`

	Out *DutchBusinessPositions `xml:"out,omitempty"`
}

type DutchBusinessGetLegalEntityRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLegalEntity"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetLegalEntityResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLegalEntityResponse"`

	Out *DutchBusinessLegalEntityData `xml:"out,omitempty"`
}

type DutchBusinessGetOrganizationTreeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetOrganizationTree"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetOrganizationTreeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetOrganizationTreeResponse"`

	Out *DutchBusinessOrganizationTree `xml:"out,omitempty"`
}

type DutchBusinessGetNonMarketingIndicatorRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetNonMarketingIndicator"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type DutchBusinessGetNonMarketingIndicatorResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetNonMarketingIndicatorResponse"`

	Out *DutchBusinessNonMarketingIndicator `xml:"out,omitempty"`
}

type DutchBusinessSearchDossierNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchDossierNumber"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchDossierNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchDossierNumberResponse"`

	Out *DutchBusinessReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessGetConcernRelationsOverviewRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetConcernRelationsOverview"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetConcernRelationsOverviewResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetConcernRelationsOverviewResponse"`

	Out *DutchBusinessGetConcernRelationsOverviewResult `xml:"out,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetConcernRelationsDetails"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Include_source bool `xml:"include_source,omitempty"`
}

type DutchBusinessGetConcernRelationsDetailsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetConcernRelationsDetailsResponse"`

	Out *DutchBusinessGetConcernRelationsDetailsResult `xml:"out,omitempty"`
}

type DutchBusinessSearchParametersV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchParametersV2"`

	Trade_name string `xml:"trade_name,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Strict_search bool `xml:"strict_search,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchParametersV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchParametersV2Response"`

	Out *DutchBusinessReferenceV2PagedResult `xml:"out,omitempty"`
}

type DutchBusinessSearchRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearch"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Trade_name string `xml:"trade_name,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Strict_search bool `xml:"strict_search,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchResponse"`

	Out *DutchBusinessEstablishmentReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessSearchEstablishmentsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchEstablishments"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchEstablishmentsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchEstablishmentsResponse"`

	Out *DutchBusinessEstablishmentReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessSearchPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchPostcodeResponse"`

	Out *DutchBusinessReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessSearchSelectionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchSelection"`

	City *StringArray `xml:"city,omitempty"`

	Postcode *StringArray `xml:"postcode,omitempty"`

	Sbi *StringArray `xml:"sbi,omitempty"`

	Primary_sbi_only bool `xml:"primary_sbi_only,omitempty"`

	Legal_form *IntArray `xml:"legal_form,omitempty"`

	Employees_min int32 `xml:"employees_min,omitempty"`

	Employees_max int32 `xml:"employees_max,omitempty"`

	Economically_active string `xml:"economically_active,omitempty"`

	Financial_status string `xml:"financial_status,omitempty"`

	Changed_since string `xml:"changed_since,omitempty"`

	New_since string `xml:"new_since,omitempty"`

	Page int32 `xml:"page,omitempty"`

	Provinces *StringArray `xml:"provinces,omitempty"`

	Sbi_match_type string `xml:"sbi_match_type,omitempty"`
}

type DutchBusinessSearchSelectionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchSelectionResponse"`

	Out *DutchBusinessReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessSearchSelectionSimpleRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchSelectionSimple"`

	City *StringArray `xml:"city,omitempty"`

	Postcode *StringArray `xml:"postcode,omitempty"`

	Sbi *StringArray `xml:"sbi,omitempty"`

	Primary_sbi_only bool `xml:"primary_sbi_only,omitempty"`

	Legal_form *IntArray `xml:"legal_form,omitempty"`

	Employees_min int32 `xml:"employees_min,omitempty"`

	Employees_max int32 `xml:"employees_max,omitempty"`

	Economically_active string `xml:"economically_active,omitempty"`

	Financial_status string `xml:"financial_status,omitempty"`

	Changed_since string `xml:"changed_since,omitempty"`

	New_since string `xml:"new_since,omitempty"`

	Page int32 `xml:"page,omitempty"`

	Provinces *StringArray `xml:"provinces,omitempty"`
}

type DutchBusinessSearchSelectionSimpleResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchSelectionSimpleResponse"`

	Out *DutchBusinessReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessGetSBIDescriptionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetSBIDescription"`

	Sbi_code string `xml:"sbi_code,omitempty"`

	Language string `xml:"language,omitempty"`
}

type DutchBusinessGetSBIDescriptionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetSBIDescriptionResponse"`

	Out *DutchBusinessSBICodeInfo `xml:"out,omitempty"`
}

type DutchBusinessGetAnnualFinancialStatementRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetAnnualFinancialStatement"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Year int32 `xml:"year,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type DutchBusinessGetAnnualFinancialStatementResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetAnnualFinancialStatementResponse"`

	Out *DutchBusinessAnnualFinancialStatement `xml:"out,omitempty"`
}

type DutchBusinessGetAdditionalPositionsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetAdditionalPositions"`

	First_name string `xml:"first_name,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Birth_date time.Time `xml:"birth_date,omitempty"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetAdditionalPositionsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetAdditionalPositionsResponse"`

	Out *DutchBusinessGetAdditionalPositionsResult `xml:"out,omitempty"`
}

type DutchBusinessLookALikesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessLookALikes"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Geo_weight float64 `xml:"geo_weight,omitempty"`

	Nr_employees_weight float64 `xml:"nr_employees_weight,omitempty"`

	Sbi_weight float64 `xml:"sbi_weight,omitempty"`

	Web_weight float64 `xml:"web_weight,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessLookALikesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessLookALikesResponse"`

	Out *DutchBusinessLookALikeCompanyPagedResult `xml:"out,omitempty"`
}

type DutchBusinessGetExtractDocumentRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocument"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Allow_caching bool `xml:"allow_caching,omitempty"`
}

type DutchBusinessGetExtractDocumentResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentResponse"`

	Out *DutchBusinessExtractDocument `xml:"out,omitempty"`
}

type DutchBusinessGetExtractDocumentDataRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentData"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Allow_caching bool `xml:"allow_caching,omitempty"`
}

type DutchBusinessGetExtractDocumentDataResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentDataResponse"`

	Out *DutchBusinessExtractDocumentData `xml:"out,omitempty"`
}

type DutchBusinessGetExtractDocumentDataV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentDataV2"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetExtractDocumentDataV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentDataV2Response"`

	Out *DutchBusinessExtractDocumentDataV2 `xml:"out,omitempty"`
}

type DutchBusinessGetExtractDocumentDataV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentDataV3"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Include_source bool `xml:"include_source,omitempty"`
}

type DutchBusinessGetExtractDocumentDataV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractDocumentDataV3Response"`

	Out *DutchBusinessExtractDocumentDataV3 `xml:"out,omitempty"`
}

type DutchBusinessGetLegalExtractDocumentDataV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLegalExtractDocumentDataV2"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetLegalExtractDocumentDataV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLegalExtractDocumentDataV2Response"`

	Out *DutchBusinessExtractDocumentDataV2 `xml:"out,omitempty"`
}

type DutchBusinessGetLegalExtractDocumentDataV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLegalExtractDocumentDataV3"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Include_source bool `xml:"include_source,omitempty"`
}

type DutchBusinessGetLegalExtractDocumentDataV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLegalExtractDocumentDataV3Response"`

	Out *DutchBusinessExtractDocumentDataV3 `xml:"out,omitempty"`
}

type DutchBusinessGetExtractHistoryRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistory"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Period_start_date time.Time `xml:"period_start_date,omitempty"`

	Period_end_date time.Time `xml:"period_end_date,omitempty"`
}

type DutchBusinessGetExtractHistoryResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryResponse"`

	Out *DutchBusinessExtractHistory `xml:"out,omitempty"`
}

type DutchBusinessGetExtractHistoryChangedRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryChanged"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Period_start_date time.Time `xml:"period_start_date,omitempty"`

	Period_end_date time.Time `xml:"period_end_date,omitempty"`
}

type DutchBusinessGetExtractHistoryChangedResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryChangedResponse"`

	Out *DutchBusinessExtractHistory `xml:"out,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentData"`

	Extract_id string `xml:"extract_id,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataResponse"`

	Out *DutchBusinessExtractDocumentData `xml:"out,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataV2"`

	Extract_id string `xml:"extract_id,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataV2Response"`

	Out *DutchBusinessExtractDocumentDataV2 `xml:"out,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataV3"`

	Extract_id string `xml:"extract_id,omitempty"`

	Include_source bool `xml:"include_source,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataV3Response"`

	Out *DutchBusinessExtractDocumentDataV3 `xml:"out,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataV3ByDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataV3ByDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Include_source bool `xml:"include_source,omitempty"`

	Period_start time.Time `xml:"period_start,omitempty"`

	Period_end time.Time `xml:"period_end,omitempty"`
}

type DutchBusinessGetExtractHistoryDocumentDataV3ByDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetExtractHistoryDocumentDataV3ByDossierResponse"`

	Out *DutchBusinessExtractDocumentDataV3 `xml:"out,omitempty"`
}

type DutchBusinessGetDossierHistoryRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierHistory"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Period_start_date time.Time `xml:"period_start_date,omitempty"`

	Period_end_date time.Time `xml:"period_end_date,omitempty"`
}

type DutchBusinessGetDossierHistoryResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetDossierHistoryResponse"`

	Out *DutchBusinessDossierHistory `xml:"out,omitempty"`
}

type DutchBusinessUpdateGetDossiersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateGetDossiers"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessUpdateGetDossiersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateGetDossiersResponse"`

	Out *DutchBusinessUpdateReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessUpdateCheckDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateCheckDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`
}

type DutchBusinessUpdateCheckDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateCheckDossierResponse"`

	Out *DutchBusinessUpdateReference `xml:"out,omitempty"`
}

type DutchBusinessUpdateGetChangedDossiersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateGetChangedDossiers"`

	Changed_since time.Time `xml:"changed_since,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessUpdateGetChangedDossiersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateGetChangedDossiersResponse"`

	Out *DutchBusinessUpdateReferencePagedResult `xml:"out,omitempty"`
}

type DutchBusinessUpdateAddDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateAddDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type DutchBusinessUpdateAddDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateAddDossierResponse"`
}

type DutchBusinessUpdateRemoveDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateRemoveDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type DutchBusinessUpdateRemoveDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateRemoveDossierResponse"`
}

type DutchBusinessUpdateGetOverviewRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateGetOverview"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessUpdateGetOverviewResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUpdateGetOverviewResponse"`

	Out *DutchBusinessUpdateSubscriptionPagedResult `xml:"out,omitempty"`
}

type DutchBusinessSearchNewsByDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchNewsByDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Period_start time.Time `xml:"period_start,omitempty"`

	Period_end time.Time `xml:"period_end,omitempty"`

	Sort_order string `xml:"sort_order,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchNewsByDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchNewsByDossierResponse"`

	Out *DutchBusinessNewsItemPagedResult `xml:"out,omitempty"`
}

type DutchBusinessUBOStartInvestigationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOStartInvestigation"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Oldest_extract_date time.Time `xml:"oldest_extract_date,omitempty"`

	Use_updates bool `xml:"use_updates,omitempty"`
}

type DutchBusinessUBOStartInvestigationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOStartInvestigationResponse"`

	Out *DutchBusinessUBOInvestigationToken `xml:"out,omitempty"`
}

type DutchBusinessUBOCheckInvestigationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOCheckInvestigation"`

	Token string `xml:"token,omitempty"`
}

type DutchBusinessUBOCheckInvestigationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOCheckInvestigationResponse"`

	Out *DutchBusinessUBOInvestigationStatus `xml:"out,omitempty"`
}

type DutchBusinessUBOPickupInvestigationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOPickupInvestigation"`

	Token string `xml:"token,omitempty"`

	Include_source bool `xml:"include_source,omitempty"`
}

type DutchBusinessUBOPickupInvestigationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOPickupInvestigationResponse"`

	Out *DutchBusinessUBOInvestigationResult `xml:"out,omitempty"`
}

type DutchBusinessUBOCostsInvestigationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOCostsInvestigation"`

	Token string `xml:"token,omitempty"`
}

type DutchBusinessUBOCostsInvestigationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOCostsInvestigationResponse"`

	Out *DutchBusinessUBOCostsInvestigationResult `xml:"out,omitempty"`
}

type DutchBusinessUBOClassifyInvestigationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOClassifyInvestigation"`

	Token string `xml:"token,omitempty"`

	Scenario string `xml:"scenario,omitempty"`
}

type DutchBusinessUBOClassifyInvestigationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessUBOClassifyInvestigationResponse"`

	Out *DutchBusinessUBOClassificationResult `xml:"out,omitempty"`
}

type DutchBusinessGetLeiRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLei"`

	Dossier_number string `xml:"dossier_number,omitempty"`
}

type DutchBusinessGetLeiResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessGetLeiResponse"`

	Out *DutchBusinessGetLeiResult `xml:"out,omitempty"`
}

type DutchBusinessSearchParametersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchParameters"`

	Trade_name string `xml:"trade_name,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Strict_search bool `xml:"strict_search,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type DutchBusinessSearchParametersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchBusinessSearchParametersResponse"`

	Out *DutchBusinessReferencePagedResult `xml:"out,omitempty"`
}

type DutchVehicleGetVehicleV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetVehicleV2"`

	License_plate string `xml:"license_plate,omitempty"`
}

type DutchVehicleGetVehicleV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetVehicleV2Response"`

	Out *DutchVehicleV2 `xml:"out,omitempty"`
}

type DutchVehicleGetPurchaseReferenceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetPurchaseReference"`

	License_plate string `xml:"license_plate,omitempty"`
}

type DutchVehicleGetPurchaseReferenceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetPurchaseReferenceResponse"`

	Out *DutchVehiclePurchaseReference `xml:"out,omitempty"`
}

type DutchVehicleGetOwnerHistoryRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetOwnerHistory"`

	License_plate string `xml:"license_plate,omitempty"`
}

type DutchVehicleGetOwnerHistoryResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetOwnerHistoryResponse"`

	Out *DutchVehicleOwnerHistory `xml:"out,omitempty"`
}

type DutchVehicleGetMarketValueRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetMarketValue"`

	License_plate string `xml:"license_plate,omitempty"`
}

type DutchVehicleGetMarketValueResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetMarketValueResponse"`

	Out *DutchVehicleMarketValue `xml:"out,omitempty"`
}

type DutchVehicleGetVehicleRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetVehicle"`

	License_plate string `xml:"license_plate,omitempty"`
}

type DutchVehicleGetVehicleResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ dutchVehicleGetVehicleResponse"`

	Out *DutchVehicle `xml:"out,omitempty"`
}

type EdrGetScoreRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ edrGetScore"`

	Last_name string `xml:"last_name,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Surname_prefix string `xml:"surname_prefix,omitempty"`

	Gender string `xml:"gender,omitempty"`

	Birth_date string `xml:"birth_date,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number string `xml:"house_number,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`
}

type EdrGetScoreResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ edrGetScoreResponse"`

	Out *EDRScore `xml:"out,omitempty"`
}

type GeoLocationNeighborhoodCoordinatesRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationNeighborhoodCoordinatesRD"`

	Nbcode string `xml:"nbcode,omitempty"`
}

type GeoLocationNeighborhoodCoordinatesRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationNeighborhoodCoordinatesRDResponse"`

	Coordinates *RDCoordinates `xml:"coordinates,omitempty"`
}

type GeoLocationPostcodeCoordinatesRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationPostcodeCoordinatesRD"`

	Postcode string `xml:"postcode,omitempty"`
}

type GeoLocationPostcodeCoordinatesRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationPostcodeCoordinatesRDResponse"`

	Coordinates *RDCoordinates `xml:"coordinates,omitempty"`
}

type GeoLocationNeighborhoodCoordinatesLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationNeighborhoodCoordinatesLatLon"`

	Nbcode string `xml:"nbcode,omitempty"`
}

type GeoLocationNeighborhoodCoordinatesLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationNeighborhoodCoordinatesLatLonResponse"`

	Coordinates *LatLonCoordinates `xml:"coordinates,omitempty"`
}

type GeoLocationPostcodeCoordinatesLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationPostcodeCoordinatesLatLon"`

	Postcode string `xml:"postcode,omitempty"`
}

type GeoLocationPostcodeCoordinatesLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationPostcodeCoordinatesLatLonResponse"`

	Coordinates *LatLonCoordinates `xml:"coordinates,omitempty"`
}

type GeoLocationAddressCoordinatesLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationAddressCoordinatesLatLon"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`
}

type GeoLocationAddressCoordinatesLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationAddressCoordinatesLatLonResponse"`

	Coordinates *LatLonCoordinatesMatch `xml:"coordinates,omitempty"`
}

type GeoLocationAddressCoordinatesRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationAddressCoordinatesRD"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`
}

type GeoLocationAddressCoordinatesRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationAddressCoordinatesRDResponse"`

	Coordinates *RDCoordinatesMatch `xml:"coordinates,omitempty"`
}

type GeoLocationLatLonToPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationLatLonToPostcode"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type GeoLocationLatLonToPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationLatLonToPostcodeResponse"`

	Postcode string `xml:"postcode,omitempty"`
}

type GeoLocationLatLonToAddressV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationLatLonToAddressV2"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type GeoLocationLatLonToAddressV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationLatLonToAddressV2Response"`

	Address *GeoLocationAddressV2 `xml:"address,omitempty"`
}

type GeoLocationRDToPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationRDToPostcode"`

	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`
}

type GeoLocationRDToPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationRDToPostcodeResponse"`

	Postcode string `xml:"postcode,omitempty"`
}

type GeoLocationRDToAddressV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationRDToAddressV2"`

	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`
}

type GeoLocationRDToAddressV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationRDToAddressV2Response"`

	Address *GeoLocationAddressV2 `xml:"address,omitempty"`
}

type GeoLocationInternationalPostcodeCoordinatesLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalPostcodeCoordinatesLatLon"`

	Postcode string `xml:"postcode,omitempty"`

	Country string `xml:"country,omitempty"`
}

type GeoLocationInternationalPostcodeCoordinatesLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalPostcodeCoordinatesLatLonResponse"`

	Coordinates *LatLonCoordinates `xml:"coordinates,omitempty"`
}

type GeoLocationInternationalAddressCoordinatesLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalAddressCoordinatesLatLon"`

	Street string `xml:"street,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	City string `xml:"city,omitempty"`

	Province string `xml:"province,omitempty"`

	Country string `xml:"country,omitempty"`

	Language string `xml:"language,omitempty"`
}

type GeoLocationInternationalAddressCoordinatesLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalAddressCoordinatesLatLonResponse"`

	Coordinates *LatLonCoordinatesInternationalAddressArray `xml:"coordinates,omitempty"`
}

type GeoLocationInternationalAddressCoordinatesLatLonV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalAddressCoordinatesLatLonV2"`

	Country string `xml:"country,omitempty"`

	Postalcode string `xml:"postalcode,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Street string `xml:"street,omitempty"`

	City string `xml:"city,omitempty"`

	Province string `xml:"province,omitempty"`

	Matchrate float32 `xml:"matchrate,omitempty"`

	Language string `xml:"language,omitempty"`
}

type GeoLocationInternationalAddressCoordinatesLatLonV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalAddressCoordinatesLatLonV2Response"`

	Coordinates *LatLonCoordinatesInternationalAddressArray `xml:"coordinates,omitempty"`
}

type GeoLocationInternationalLatLonToAddressRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalLatLonToAddress"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type GeoLocationInternationalLatLonToAddressResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationInternationalLatLonToAddressResponse"`

	Address *GeoLocationInternationalAddress `xml:"address,omitempty"`
}

type GeoLocationNeighborhoodDistanceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationNeighborhoodDistance"`

	Nbcodefrom string `xml:"nbcodefrom,omitempty"`

	Nbcodeto string `xml:"nbcodeto,omitempty"`
}

type GeoLocationNeighborhoodDistanceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationNeighborhoodDistanceResponse"`

	Distance int32 `xml:"distance,omitempty"`
}

type GeoLocationPostcodeDistanceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationPostcodeDistance"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Postcodeto string `xml:"postcodeto,omitempty"`
}

type GeoLocationPostcodeDistanceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationPostcodeDistanceResponse"`

	Distance int32 `xml:"distance,omitempty"`
}

type GeoLocationHaversineDistanceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationHaversineDistance"`

	Latitude_coord_1 float32 `xml:"latitude_coord_1,omitempty"`

	Longitude_coord_1 float32 `xml:"longitude_coord_1,omitempty"`

	Latitude_coord_2 float32 `xml:"latitude_coord_2,omitempty"`

	Longitude_coord_2 float32 `xml:"longitude_coord_2,omitempty"`

	In_radians bool `xml:"in_radians,omitempty"`
}

type GeoLocationHaversineDistanceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationHaversineDistanceResponse"`

	Distance int32 `xml:"distance,omitempty"`
}

type GeoLocationDistanceSortedNeighborhoodCodesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedNeighborhoodCodes"`

	Nbcodefrom string `xml:"nbcodefrom,omitempty"`

	Nbcodes *StringArray `xml:"nbcodes,omitempty"`
}

type GeoLocationDistanceSortedNeighborhoodCodesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedNeighborhoodCodesResponse"`

	Nbcodes *SortedPostcodeArray `xml:"nbcodes,omitempty"`
}

type GeoLocationDistanceSortedNeighborhoodCodesRadiusRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedNeighborhoodCodesRadius"`

	Nbcodefrom string `xml:"nbcodefrom,omitempty"`

	Radius int32 `xml:"radius,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type GeoLocationDistanceSortedNeighborhoodCodesRadiusResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedNeighborhoodCodesRadiusResponse"`

	Out *SortedPostcodePagedResult `xml:"out,omitempty"`
}

type GeoLocationDistanceSortedPostcodesRadiusRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedPostcodesRadius"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Radius int32 `xml:"radius,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type GeoLocationDistanceSortedPostcodesRadiusResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedPostcodesRadiusResponse"`

	Out *SortedPostcodePagedResult `xml:"out,omitempty"`
}

type GeoLocationDistanceSortedPostcodesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedPostcodes"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Postcodes *StringArray `xml:"postcodes,omitempty"`
}

type GeoLocationDistanceSortedPostcodesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationDistanceSortedPostcodesResponse"`

	Postcodes *SortedPostcodeArray `xml:"postcodes,omitempty"`
}

type GeoLocationLatLonToRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationLatLonToRD"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`
}

type GeoLocationLatLonToRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationLatLonToRDResponse"`

	Out *RDCoordinates `xml:"out,omitempty"`
}

type GeoLocationRDToLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationRDToLatLon"`

	X int32 `xml:"x,omitempty"`

	Y int32 `xml:"y,omitempty"`
}

type GeoLocationRDToLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ geoLocationRDToLatLonResponse"`

	Out *LatLonCoordinates `xml:"out,omitempty"`
}

type GraydonCreditGetReportRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditGetReport"`

	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`

	Document string `xml:"document,omitempty"`
}

type GraydonCreditGetReportResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditGetReportResponse"`

	Out *GraydonCreditReport `xml:"out,omitempty"`
}

type GraydonCreditSearchIdentificationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditSearchIdentification"`

	Company_id string `xml:"company_id,omitempty"`

	Company_id_type string `xml:"company_id_type,omitempty"`

	Country_iso2 string `xml:"country_iso2,omitempty"`
}

type GraydonCreditSearchIdentificationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditSearchIdentificationResponse"`

	Out *GraydonReferenceArray `xml:"out,omitempty"`
}

type GraydonCreditSearchNameRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditSearchName"`

	Company_name string `xml:"company_name,omitempty"`

	Residence string `xml:"residence,omitempty"`

	Country_iso2 string `xml:"country_iso2,omitempty"`
}

type GraydonCreditSearchNameResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditSearchNameResponse"`

	Out *GraydonReferenceArray `xml:"out,omitempty"`
}

type GraydonCreditSearchPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditSearchPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	House_no int32 `xml:"house_no,omitempty"`

	Telephone_no string `xml:"telephone_no,omitempty"`

	Country_iso2 string `xml:"country_iso2,omitempty"`
}

type GraydonCreditSearchPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditSearchPostcodeResponse"`

	Out *GraydonReferenceArray `xml:"out,omitempty"`
}

type GraydonCreditQuickscanRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditQuickscan"`

	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`
}

type GraydonCreditQuickscanResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditQuickscanResponse"`

	Out *GraydonCreditReportQuickscan `xml:"out,omitempty"`
}

type GraydonCreditRatingsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditRatings"`

	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`
}

type GraydonCreditRatingsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditRatingsResponse"`

	Out *GraydonCreditReportRatings `xml:"out,omitempty"`
}

type GraydonCreditVatNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditVatNumber"`

	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`
}

type GraydonCreditVatNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditVatNumberResponse"`

	Out *GraydonCreditReportVatNumber `xml:"out,omitempty"`
}

type GraydonCreditCompanyLiaisonsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditCompanyLiaisons"`

	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`
}

type GraydonCreditCompanyLiaisonsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditCompanyLiaisonsResponse"`

	Out *GraydonCreditReportCompanyLiaisons `xml:"out,omitempty"`
}

type GraydonCreditManagementRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditManagement"`

	Graydon_company_id int32 `xml:"graydon_company_id,omitempty"`
}

type GraydonCreditManagementResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ graydonCreditManagementResponse"`

	Out *GraydonCreditReportManagement `xml:"out,omitempty"`
}

type InsolvencyGetCaseByPublicationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ insolvencyGetCaseByPublication"`

	Publication_id string `xml:"publication_id,omitempty"`
}

type InsolvencyGetCaseByPublicationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ insolvencyGetCaseByPublicationResponse"`

	Out *InsolvencyCase `xml:"out,omitempty"`
}

type InsolvencySearchPublicationsByCoCNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ insolvencySearchPublicationsByCoCNumber"`

	Coc_number string `xml:"coc_number,omitempty"`
}

type InsolvencySearchPublicationsByCoCNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ insolvencySearchPublicationsByCoCNumberResponse"`

	Out *InsolvencyPublicationList `xml:"out,omitempty"`
}

type InsolvencySearchPublicationsByPersonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ insolvencySearchPublicationsByPerson"`

	Last_name string `xml:"last_name,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Birth_date time.Time `xml:"birth_date,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`
}

type InsolvencySearchPublicationsByPersonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ insolvencySearchPublicationsByPersonResponse"`

	Out *InsolvencyPublicationList `xml:"out,omitempty"`
}

type InternationalAddressSearchV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ internationalAddressSearchV2"`

	Organization string `xml:"organization,omitempty"`

	Building string `xml:"building,omitempty"`

	Street string `xml:"street,omitempty"`

	Housenr string `xml:"housenr,omitempty"`

	Pobox string `xml:"pobox,omitempty"`

	Locality string `xml:"locality,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Province string `xml:"province,omitempty"`

	Country string `xml:"country,omitempty"`

	Language string `xml:"language,omitempty"`

	Country_format string `xml:"country_format,omitempty"`
}

type InternationalAddressSearchV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ internationalAddressSearchV2Response"`

	Out *InternationalAddressSearchV2Result `xml:"out,omitempty"`
}

type InternationalAddressSearchInteractiveRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ internationalAddressSearchInteractive"`

	Organization string `xml:"organization,omitempty"`

	Building string `xml:"building,omitempty"`

	Street string `xml:"street,omitempty"`

	Housenr string `xml:"housenr,omitempty"`

	Pobox string `xml:"pobox,omitempty"`

	Locality string `xml:"locality,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	Province string `xml:"province,omitempty"`

	Country string `xml:"country,omitempty"`

	Language string `xml:"language,omitempty"`

	Country_format string `xml:"country_format,omitempty"`
}

type InternationalAddressSearchInteractiveResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ internationalAddressSearchInteractiveResponse"`

	Out *InternationalAddressSearchV2Result `xml:"out,omitempty"`
}

type KadasterEigendomsBerichtDocumentPerceelRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtDocumentPerceel"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterEigendomsBerichtDocumentPerceelResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtDocumentPerceelResponse"`

	Berichtobjectresultaat *BerichtObjectDocumentResultaat `xml:"berichtobjectresultaat,omitempty"`
}

type KadasterEigendomsBerichtDocumentPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtDocumentPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterEigendomsBerichtDocumentPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtDocumentPostcodeResponse"`

	Berichtobjectresultaat *BerichtObjectDocumentResultaat `xml:"berichtobjectresultaat,omitempty"`
}

type KadasterEigendomsBerichtPerceelV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPerceelV2"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`
}

type KadasterEigendomsBerichtPerceelV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPerceelV2Response"`

	Berichtobjectresultaat *BerichtObjectResultaatV2 `xml:"berichtobjectresultaat,omitempty"`
}

type KadasterEigendomsBerichtPostcodeV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPostcodeV2"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`
}

type KadasterEigendomsBerichtPostcodeV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPostcodeV2Response"`

	Berichtobjectresultaat *BerichtObjectResultaatV2 `xml:"berichtobjectresultaat,omitempty"`
}

type KadasterKoopsommenOverzichtV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKoopsommenOverzichtV2"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterKoopsommenOverzichtV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKoopsommenOverzichtV2Response"`

	Koopsommenoverzicht *KoopsommenOverzichtV2 `xml:"koopsommenoverzicht,omitempty"`
}

type KadasterUittrekselKadastraleKaartPerceelV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPerceelV3"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterUittrekselKadastraleKaartPerceelV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPerceelV3Response"`

	Out *KadasterUittrekselKadastraleKaartResultaatV2 `xml:"out,omitempty"`
}

type KadasterUittrekselKadastraleKaartPostcodeV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPostcodeV3"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterUittrekselKadastraleKaartPostcodeV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPostcodeV3Response"`

	Out *KadasterUittrekselKadastraleKaartResultaatV2 `xml:"out,omitempty"`
}

type KadasterHypothecairBerichtPostcodeV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPostcodeV3"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterHypothecairBerichtPostcodeV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPostcodeV3Response"`

	Out *KadasterHypothecairBerichtResultaatV2 `xml:"out,omitempty"`
}

type KadasterHypothecairBerichtPerceelV3RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPerceelV3"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterHypothecairBerichtPerceelV3ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPerceelV3Response"`

	Out *KadasterHypothecairBerichtResultaatV2 `xml:"out,omitempty"`
}

type KadasterBronDocumentRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterBronDocument"`

	Aanduiding_soort_register string `xml:"aanduiding_soort_register,omitempty"`

	Deel string `xml:"deel,omitempty"`

	Nummer string `xml:"nummer,omitempty"`

	Reeks string `xml:"reeks,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterBronDocumentResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterBronDocumentResponse"`

	Bron_document *KadasterBronDocument `xml:"bron_document,omitempty"`
}

type KadasterValueListGetRangesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterValueListGetRanges"`
}

type KadasterValueListGetRangesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterValueListGetRangesResponse"`

	Out *KadasterValueList `xml:"out,omitempty"`
}

type KadasterValueListGetMunicipalitiesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterValueListGetMunicipalities"`
}

type KadasterValueListGetMunicipalitiesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterValueListGetMunicipalitiesResponse"`

	Out *KadasterValueList `xml:"out,omitempty"`
}

type KadasterAddressCoordinatesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterAddressCoordinates"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Houseno int32 `xml:"houseno,omitempty"`

	Housenoaddition string `xml:"housenoaddition,omitempty"`
}

type KadasterAddressCoordinatesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterAddressCoordinatesResponse"`

	Coordinates *KadasterCoordinates `xml:"coordinates,omitempty"`
}

type KadasterKadastraleKaartPerceelV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPerceelV2"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`

	Schaal int32 `xml:"schaal,omitempty"`
}

type KadasterKadastraleKaartPerceelV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPerceelV2Response"`

	Out *KadasterKadastraleKaartResultaatV2 `xml:"out,omitempty"`
}

type KadasterKadastraleKaartPostcodeV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPostcodeV2"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`

	Schaal int32 `xml:"schaal,omitempty"`
}

type KadasterKadastraleKaartPostcodeV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPostcodeV2Response"`

	Out *KadasterKadastraleKaartResultaatV2 `xml:"out,omitempty"`
}

type KadasterUittrekselKadastraleKaartPerceelRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPerceel"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterUittrekselKadastraleKaartPerceelResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPerceelResponse"`

	Uittreksel *UittrekselKadastraleKaart `xml:"uittreksel,omitempty"`
}

type KadasterUittrekselKadastraleKaartPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterUittrekselKadastraleKaartPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPostcodeResponse"`

	Uittreksel *UittrekselKadastraleKaart `xml:"uittreksel,omitempty"`
}

type KadasterHypothecairBerichtPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterHypothecairBerichtPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPostcodeResponse"`

	Out *KadasterHypothecairBerichtResultaat `xml:"out,omitempty"`
}

type KadasterHypothecairBerichtPerceelRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPerceel"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterHypothecairBerichtPerceelResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPerceelResponse"`

	Out *KadasterHypothecairBerichtResultaat `xml:"out,omitempty"`
}

type KadasterEigendomsBerichtPerceelRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPerceel"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`
}

type KadasterEigendomsBerichtPerceelResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPerceelResponse"`

	Berichtobjectresultaat *BerichtObjectResultaat `xml:"berichtobjectresultaat,omitempty"`
}

type KadasterEigendomsBerichtPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`
}

type KadasterEigendomsBerichtPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterEigendomsBerichtPostcodeResponse"`

	Berichtobjectresultaat *BerichtObjectResultaat `xml:"berichtobjectresultaat,omitempty"`
}

type KadasterUittrekselKadastraleKaartPerceelV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPerceelV2"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterUittrekselKadastraleKaartPerceelV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPerceelV2Response"`

	Out *KadasterUittrekselKadastraleKaartResultaat `xml:"out,omitempty"`
}

type KadasterUittrekselKadastraleKaartPostcodeV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPostcodeV2"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterUittrekselKadastraleKaartPostcodeV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterUittrekselKadastraleKaartPostcodeV2Response"`

	Out *KadasterUittrekselKadastraleKaartResultaat `xml:"out,omitempty"`
}

type KadasterKadastraleKaartPerceelRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPerceel"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`

	Schaal int32 `xml:"schaal,omitempty"`
}

type KadasterKadastraleKaartPerceelResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPerceelResponse"`

	Out *KadasterKadastraleKaartResultaat `xml:"out,omitempty"`
}

type KadasterKadastraleKaartPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`

	Schaal int32 `xml:"schaal,omitempty"`
}

type KadasterKadastraleKaartPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKadastraleKaartPostcodeResponse"`

	Out *KadasterKadastraleKaartResultaat `xml:"out,omitempty"`
}

type KadasterHypothecairBerichtPostcodeV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPostcodeV2"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterHypothecairBerichtPostcodeV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPostcodeV2Response"`

	Out *KadasterHypothecairBerichtResultaat `xml:"out,omitempty"`
}

type KadasterHypothecairBerichtPerceelV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPerceelV2"`

	Gemeentecode string `xml:"gemeentecode,omitempty"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer string `xml:"perceelnummer,omitempty"`

	Relatiecode string `xml:"relatiecode,omitempty"`

	Volgnummer string `xml:"volgnummer,omitempty"`

	Format string `xml:"format,omitempty"`
}

type KadasterHypothecairBerichtPerceelV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterHypothecairBerichtPerceelV2Response"`

	Out *KadasterHypothecairBerichtResultaat `xml:"out,omitempty"`
}

type KadasterKoopsommenOverzichtRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKoopsommenOverzicht"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`
}

type KadasterKoopsommenOverzichtResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterKoopsommenOverzichtResponse"`

	Koopsommenoverzicht *KoopsommenOverzicht `xml:"koopsommenoverzicht,omitempty"`
}

type KadasterV2GetKadastraalBerichtObjectByAdresRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKadastraalBerichtObjectByAdres"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Document string `xml:"document,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`
}

type KadasterV2GetKadastraalBerichtObjectByAdresResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKadastraalBerichtObjectByAdresResponse"`

	Resultaat *KadasterV2KadastraalBerichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKadastraalBerichtObjectByKadastraleAanduiding"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer int32 `xml:"perceelnummer,omitempty"`

	Appartementsnummer int32 `xml:"appartementsnummer,omitempty"`

	Document string `xml:"document,omitempty"`
}

type KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingResponse"`

	Resultaat *KadasterV2KadastraalBerichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetKadastraalBerichtObjectByObjectIdRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKadastraalBerichtObjectByObjectId"`

	Object_id string `xml:"object_id,omitempty"`

	Document string `xml:"document,omitempty"`
}

type KadasterV2GetKadastraalBerichtObjectByObjectIdResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKadastraalBerichtObjectByObjectIdResponse"`

	Resultaat *KadasterV2KadastraalBerichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetHypothecairBerichtObjectByAdresRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetHypothecairBerichtObjectByAdres"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Document string `xml:"document,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`
}

type KadasterV2GetHypothecairBerichtObjectByAdresResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetHypothecairBerichtObjectByAdresResponse"`

	Resultaat *KadasterV2HypothecairBerichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetHypothecairBerichtObjectByKadastraleAanduiding"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer int32 `xml:"perceelnummer,omitempty"`

	Appartementsnummer int32 `xml:"appartementsnummer,omitempty"`

	Document string `xml:"document,omitempty"`
}

type KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingResponse"`

	Resultaat *KadasterV2HypothecairBerichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetHypothecairBerichtObjectByObjectIdRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetHypothecairBerichtObjectByObjectId"`

	Object_id string `xml:"object_id,omitempty"`

	Document string `xml:"document,omitempty"`
}

type KadasterV2GetHypothecairBerichtObjectByObjectIdResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetHypothecairBerichtObjectByObjectIdResponse"`

	Resultaat *KadasterV2HypothecairBerichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetUittrekselKadastraleKaartByAdresRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetUittrekselKadastraleKaartByAdres"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`
}

type KadasterV2GetUittrekselKadastraleKaartByAdresResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetUittrekselKadastraleKaartByAdresResponse"`

	Resultaat *KadasterV2UittrekselKadastraleKaartResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetUittrekselKadastraleKaartByKadastraleAanduiding"`

	Gemeentenaam string `xml:"gemeentenaam,omitempty"`

	Sectie string `xml:"sectie,omitempty"`

	Perceelnummer int32 `xml:"perceelnummer,omitempty"`

	Appartementsnummer int32 `xml:"appartementsnummer,omitempty"`
}

type KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingResponse"`

	Resultaat *KadasterV2UittrekselKadastraleKaartResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetUittrekselKadastraleKaartByObjectIdRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetUittrekselKadastraleKaartByObjectId"`

	Object_id string `xml:"object_id,omitempty"`
}

type KadasterV2GetUittrekselKadastraleKaartByObjectIdResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetUittrekselKadastraleKaartByObjectIdResponse"`

	Resultaat *KadasterV2UittrekselKadastraleKaartResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetKoopsommenOverzichtRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKoopsommenOverzicht"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Document string `xml:"document,omitempty"`

	Huisletter string `xml:"huisletter,omitempty"`
}

type KadasterV2GetKoopsommenOverzichtResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetKoopsommenOverzichtResponse"`

	Resultaat *KadasterV2KoopsommenOverzichtResponse `xml:"resultaat,omitempty"`
}

type KadasterV2GetBronDocumentRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetBronDocument"`

	Register_code string `xml:"register_code,omitempty"`

	Deel string `xml:"deel,omitempty"`

	Nummer string `xml:"nummer,omitempty"`

	Reeks string `xml:"reeks,omitempty"`

	Soort_register string `xml:"soort_register,omitempty"`
}

type KadasterV2GetBronDocumentResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kadasterV2GetBronDocumentResponse"`

	Resultaat *KadasterV2BronDocumentResponse `xml:"resultaat,omitempty"`
}

type KvkGetDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkGetDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type KvkGetDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkGetDossierResponse"`

	Out *KvkDossier `xml:"out,omitempty"`
}

type KvkSearchDossierNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchDossierNumber"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Rsin_number string `xml:"rsin_number,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type KvkSearchDossierNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchDossierNumberResponse"`

	Out *KvkReferencePagedResult `xml:"out,omitempty"`
}

type KvkSearchParametersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchParameters"`

	Trade_name string `xml:"trade_name,omitempty"`

	City string `xml:"city,omitempty"`

	Street string `xml:"street,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Telephone_number string `xml:"telephone_number,omitempty"`

	Domain_name string `xml:"domain_name,omitempty"`

	Strict_search bool `xml:"strict_search,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type KvkSearchParametersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchParametersResponse"`

	Out *KvkReferencePagedResult `xml:"out,omitempty"`
}

type KvkSearchPostcodeRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchPostcode"`

	Postcode string `xml:"postcode,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type KvkSearchPostcodeResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchPostcodeResponse"`

	Out *KvkReferencePagedResult `xml:"out,omitempty"`
}

type KvkSearchSelectionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchSelection"`

	City *StringArray `xml:"city,omitempty"`

	Postcode *StringArray `xml:"postcode,omitempty"`

	Sbi *StringArray `xml:"sbi,omitempty"`

	Primary_sbi_only bool `xml:"primary_sbi_only,omitempty"`

	Legal_form *IntArray `xml:"legal_form,omitempty"`

	Employees_min int32 `xml:"employees_min,omitempty"`

	Employees_max int32 `xml:"employees_max,omitempty"`

	Economically_active string `xml:"economically_active,omitempty"`

	Financial_status string `xml:"financial_status,omitempty"`

	Changed_since string `xml:"changed_since,omitempty"`

	New_since string `xml:"new_since,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type KvkSearchSelectionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkSearchSelectionResponse"`

	Out *KvkReferencePagedResult `xml:"out,omitempty"`
}

type KvkGetExtractDocumentRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkGetExtractDocument"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Allow_caching bool `xml:"allow_caching,omitempty"`
}

type KvkGetExtractDocumentResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkGetExtractDocumentResponse"`

	Out *KvkExtractDocument `xml:"out,omitempty"`
}

type KvkUpdateCheckDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateCheckDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`
}

type KvkUpdateCheckDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateCheckDossierResponse"`

	Out *KvkUpdateReference `xml:"out,omitempty"`
}

type KvkUpdateGetChangedDossiersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateGetChangedDossiers"`

	Changed_since time.Time `xml:"changed_since,omitempty"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type KvkUpdateGetChangedDossiersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateGetChangedDossiersResponse"`

	Out *KvkUpdateReferencePagedResult `xml:"out,omitempty"`
}

type KvkUpdateGetDossiersRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateGetDossiers"`

	Update_types *StringArray `xml:"update_types,omitempty"`

	Page int32 `xml:"page,omitempty"`
}

type KvkUpdateGetDossiersResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateGetDossiersResponse"`

	Out *KvkUpdateReferencePagedResult `xml:"out,omitempty"`
}

type KvkUpdateAddDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateAddDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type KvkUpdateAddDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateAddDossierResponse"`
}

type KvkUpdateRemoveDossierRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateRemoveDossier"`

	Dossier_number string `xml:"dossier_number,omitempty"`

	Establishment_number string `xml:"establishment_number,omitempty"`
}

type KvkUpdateRemoveDossierResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ kvkUpdateRemoveDossierResponse"`
}

type MapViewPostcodeV2RequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewPostcodeV2"`

	Postcode string `xml:"postcode,omitempty"`

	Format string `xml:"format,omitempty"`

	Width int32 `xml:"width,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Zoom float32 `xml:"zoom,omitempty"`
}

type MapViewPostcodeV2ResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewPostcodeV2Response"`

	Image []byte `xml:"image,omitempty"`
}

type MapViewLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewLatLon"`

	Center_lat float32 `xml:"center_lat,omitempty"`

	Center_lon float32 `xml:"center_lon,omitempty"`

	Extra_latlon *LatLonCoordinatesArray `xml:"extra_latlon,omitempty"`

	Format string `xml:"format,omitempty"`

	Width int32 `xml:"width,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Zoom float32 `xml:"zoom,omitempty"`
}

type MapViewLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewLatLonResponse"`

	Image []byte `xml:"image,omitempty"`
}

type MapViewRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewRD"`

	Center_x int32 `xml:"center_x,omitempty"`

	Center_y int32 `xml:"center_y,omitempty"`

	Extra_xy *RDCoordinatesArray `xml:"extra_xy,omitempty"`

	Format string `xml:"format,omitempty"`

	Width int32 `xml:"width,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Zoom float32 `xml:"zoom,omitempty"`
}

type MapViewRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewRDResponse"`

	Image []byte `xml:"image,omitempty"`
}

type MapViewInternationalLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewInternationalLatLon"`

	Latitude float32 `xml:"latitude,omitempty"`

	Longitude float32 `xml:"longitude,omitempty"`

	Format string `xml:"format,omitempty"`

	Width int32 `xml:"width,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Zoom float32 `xml:"zoom,omitempty"`
}

type MapViewInternationalLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ mapViewInternationalLatLonResponse"`

	Image []byte `xml:"image,omitempty"`
}

type MetaGetServicesRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ metaGetServices"`
}

type MetaGetServicesResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ metaGetServicesResponse"`

	Out *MetaServiceReferenceArray `xml:"out,omitempty"`
}

type MetaGetServiceRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ metaGetService"`

	Service_name string `xml:"service_name,omitempty"`
}

type MetaGetServiceResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ metaGetServiceResponse"`

	Out *MetaServiceDefinition `xml:"out,omitempty"`
}

type MetaGetMethodRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ metaGetMethod"`

	Service_name string `xml:"service_name,omitempty"`

	Method_name string `xml:"method_name,omitempty"`
}

type MetaGetMethodResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ metaGetMethodResponse"`

	Out *MetaMethodDefinition `xml:"out,omitempty"`
}

type NbwoEstimateValueRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ nbwoEstimateValue"`

	Postcode string `xml:"postcode,omitempty"`

	Huisnummer int32 `xml:"huisnummer,omitempty"`

	Huisnummer_toevoeging string `xml:"huisnummer_toevoeging,omitempty"`

	Prijspeil_datum string `xml:"prijspeil_datum,omitempty"`

	Woningtype string `xml:"woningtype,omitempty"`

	Inhoud int32 `xml:"inhoud,omitempty"`

	Grootte int32 `xml:"grootte,omitempty"`
}

type NbwoEstimateValueResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ nbwoEstimateValueResponse"`

	Out *NbwoWaarde `xml:"out,omitempty"`
}

type RiskCheckPersonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ riskCheckPerson"`

	Gender string `xml:"gender,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Birth_date string `xml:"birth_date,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`

	Account_number string `xml:"account_number,omitempty"`

	Phone_number string `xml:"phone_number,omitempty"`

	Mobile_number string `xml:"mobile_number,omitempty"`

	Email string `xml:"email,omitempty"`

	Testing_date string `xml:"testing_date,omitempty"`
}

type RiskCheckPersonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ riskCheckPersonResponse"`

	Out *RiskResult `xml:"out,omitempty"`
}

type RiskCheckGetRiskPersonCompanyReportRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ riskCheckGetRiskPersonCompanyReport"`

	Gender string `xml:"gender,omitempty"`

	Initials string `xml:"initials,omitempty"`

	Prefix string `xml:"prefix,omitempty"`

	Last_name string `xml:"last_name,omitempty"`

	Birth_date string `xml:"birth_date,omitempty"`

	Street string `xml:"street,omitempty"`

	House_number int32 `xml:"house_number,omitempty"`

	House_number_addition string `xml:"house_number_addition,omitempty"`

	Postcode string `xml:"postcode,omitempty"`

	City string `xml:"city,omitempty"`
}

type RiskCheckGetRiskPersonCompanyReportResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ riskCheckGetRiskPersonCompanyReportResponse"`

	Out *RiskPersonCompanyReport `xml:"out,omitempty"`
}

type RoutePlannerGetRouteRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerGetRoute"`

	Start_postcode string `xml:"start_postcode,omitempty"`

	Start_house_number int32 `xml:"start_house_number,omitempty"`

	Start_house_number_addition string `xml:"start_house_number_addition,omitempty"`

	Start_street string `xml:"start_street,omitempty"`

	Start_city string `xml:"start_city,omitempty"`

	Start_country string `xml:"start_country,omitempty"`

	Destination_postcode string `xml:"destination_postcode,omitempty"`

	Destination_house_number int32 `xml:"destination_house_number,omitempty"`

	Destination_house_number_addition string `xml:"destination_house_number_addition,omitempty"`

	Destination_street string `xml:"destination_street,omitempty"`

	Destination_city string `xml:"destination_city,omitempty"`

	Destination_country string `xml:"destination_country,omitempty"`

	Route_type string `xml:"route_type,omitempty"`

	Language string `xml:"language,omitempty"`
}

type RoutePlannerGetRouteResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerGetRouteResponse"`

	Route *RoutePlannerRoute `xml:"route,omitempty"`
}

type RoutePlannerInformationAddressRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerInformationAddress"`

	Routetype string `xml:"routetype,omitempty"`

	From_postalcode string `xml:"from_postalcode,omitempty"`

	From_houseno string `xml:"from_houseno,omitempty"`

	From_street string `xml:"from_street,omitempty"`

	From_city string `xml:"from_city,omitempty"`

	From_country string `xml:"from_country,omitempty"`

	To_postalcode string `xml:"to_postalcode,omitempty"`

	To_houseno string `xml:"to_houseno,omitempty"`

	To_street string `xml:"to_street,omitempty"`

	To_city string `xml:"to_city,omitempty"`

	To_country string `xml:"to_country,omitempty"`
}

type RoutePlannerInformationAddressResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerInformationAddressResponse"`

	Route *RouteInfo `xml:"route,omitempty"`
}

type RoutePlannerDescriptionAddressRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionAddress"`

	Routetype string `xml:"routetype,omitempty"`

	From_postalcode string `xml:"from_postalcode,omitempty"`

	From_houseno string `xml:"from_houseno,omitempty"`

	From_street string `xml:"from_street,omitempty"`

	From_city string `xml:"from_city,omitempty"`

	From_country string `xml:"from_country,omitempty"`

	To_postalcode string `xml:"to_postalcode,omitempty"`

	To_houseno string `xml:"to_houseno,omitempty"`

	To_street string `xml:"to_street,omitempty"`

	To_city string `xml:"to_city,omitempty"`

	To_country string `xml:"to_country,omitempty"`

	Language string `xml:"language,omitempty"`
}

type RoutePlannerDescriptionAddressResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionAddressResponse"`

	Route *RoutePartArray `xml:"route,omitempty"`
}

type RoutePlannerDescriptionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescription"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Postcodeto string `xml:"postcodeto,omitempty"`

	English bool `xml:"english,omitempty"`
}

type RoutePlannerDescriptionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionResponse"`

	Route *RoutePartArray `xml:"route,omitempty"`
}

type RoutePlannerDescriptionShortestRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionShortest"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Postcodeto string `xml:"postcodeto,omitempty"`

	English bool `xml:"english,omitempty"`
}

type RoutePlannerDescriptionShortestResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionShortestResponse"`

	Route *RoutePartArray `xml:"route,omitempty"`
}

type RoutePlannerDescriptionCoordinatesRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionCoordinatesRD"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Postcodeto string `xml:"postcodeto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`

	English bool `xml:"english,omitempty"`
}

type RoutePlannerDescriptionCoordinatesRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionCoordinatesRDResponse"`

	Route *RouteDescriptionCoordinatesRD `xml:"route,omitempty"`
}

type RoutePlannerInformationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerInformation"`

	Postcodefrom string `xml:"postcodefrom,omitempty"`

	Postcodeto string `xml:"postcodeto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`
}

type RoutePlannerInformationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerInformationResponse"`

	Route *RouteInfo `xml:"route,omitempty"`
}

type RoutePlannerRDDescriptionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerRDDescription"`

	Xfrom int32 `xml:"xfrom,omitempty"`

	Yfrom int32 `xml:"yfrom,omitempty"`

	Xto int32 `xml:"xto,omitempty"`

	Yto int32 `xml:"yto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`

	English bool `xml:"english,omitempty"`
}

type RoutePlannerRDDescriptionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerRDDescriptionResponse"`

	Route *RoutePartArray `xml:"route,omitempty"`
}

type RoutePlannerRDInformationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerRDInformation"`

	Xfrom int32 `xml:"xfrom,omitempty"`

	Yfrom int32 `xml:"yfrom,omitempty"`

	Xto int32 `xml:"xto,omitempty"`

	Yto int32 `xml:"yto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`
}

type RoutePlannerRDInformationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerRDInformationResponse"`

	Route *RouteInfo `xml:"route,omitempty"`
}

type RoutePlannerRDDescriptionCoordinatesRDRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerRDDescriptionCoordinatesRD"`

	Xfrom int32 `xml:"xfrom,omitempty"`

	Yfrom int32 `xml:"yfrom,omitempty"`

	Xto int32 `xml:"xto,omitempty"`

	Yto int32 `xml:"yto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`

	English bool `xml:"english,omitempty"`
}

type RoutePlannerRDDescriptionCoordinatesRDResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerRDDescriptionCoordinatesRDResponse"`

	Route *RouteDescriptionCoordinatesRD `xml:"route,omitempty"`
}

type RoutePlannerInformationDutchAddressRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerInformationDutchAddress"`

	Routetype string `xml:"routetype,omitempty"`

	From_postalcode string `xml:"from_postalcode,omitempty"`

	From_housno string `xml:"from_housno,omitempty"`

	From_street string `xml:"from_street,omitempty"`

	From_city string `xml:"from_city,omitempty"`

	To_postalcode string `xml:"to_postalcode,omitempty"`

	To_housno string `xml:"to_housno,omitempty"`

	To_street string `xml:"to_street,omitempty"`

	To_city string `xml:"to_city,omitempty"`
}

type RoutePlannerInformationDutchAddressResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerInformationDutchAddressResponse"`

	Route *RouteInfo `xml:"route,omitempty"`
}

type RoutePlannerDescriptionDutchAddressRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionDutchAddress"`

	Routetype string `xml:"routetype,omitempty"`

	From_postalcode string `xml:"from_postalcode,omitempty"`

	From_housno string `xml:"from_housno,omitempty"`

	From_street string `xml:"from_street,omitempty"`

	From_city string `xml:"from_city,omitempty"`

	To_postalcode string `xml:"to_postalcode,omitempty"`

	To_housno string `xml:"to_housno,omitempty"`

	To_street string `xml:"to_street,omitempty"`

	To_city string `xml:"to_city,omitempty"`

	Language string `xml:"language,omitempty"`
}

type RoutePlannerDescriptionDutchAddressResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerDescriptionDutchAddressResponse"`

	Route *RoutePartArray `xml:"route,omitempty"`
}

type RoutePlannerEUDescriptionRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUDescription"`

	Latitudefrom float32 `xml:"latitudefrom,omitempty"`

	Longitudefrom float32 `xml:"longitudefrom,omitempty"`

	Latitudeto float32 `xml:"latitudeto,omitempty"`

	Longitudeto float32 `xml:"longitudeto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`

	Language string `xml:"language,omitempty"`
}

type RoutePlannerEUDescriptionResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUDescriptionResponse"`

	Route *RoutePartArray `xml:"route,omitempty"`
}

type RoutePlannerEUInformationRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUInformation"`

	Latitudefrom float32 `xml:"latitudefrom,omitempty"`

	Longitudefrom float32 `xml:"longitudefrom,omitempty"`

	Latitudeto float32 `xml:"latitudeto,omitempty"`

	Longitudeto float32 `xml:"longitudeto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`
}

type RoutePlannerEUInformationResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUInformationResponse"`

	Route *RouteInfo `xml:"route,omitempty"`
}

type RoutePlannerEUDescriptionCoordinatesLatLonRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUDescriptionCoordinatesLatLon"`

	Latitudefrom float32 `xml:"latitudefrom,omitempty"`

	Longitudefrom float32 `xml:"longitudefrom,omitempty"`

	Latitudeto float32 `xml:"latitudeto,omitempty"`

	Longitudeto float32 `xml:"longitudeto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`

	Language string `xml:"language,omitempty"`
}

type RoutePlannerEUDescriptionCoordinatesLatLonResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUDescriptionCoordinatesLatLonResponse"`

	Route *RouteDescriptionCoordinatesLatLon `xml:"route,omitempty"`
}

type RoutePlannerEUMapRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUMap"`

	Latitudefrom float32 `xml:"latitudefrom,omitempty"`

	Longitudefrom float32 `xml:"longitudefrom,omitempty"`

	Latitudeto float32 `xml:"latitudeto,omitempty"`

	Longitudeto float32 `xml:"longitudeto,omitempty"`

	Routetype string `xml:"routetype,omitempty"`

	Language string `xml:"language,omitempty"`

	Format string `xml:"format,omitempty"`

	Width int32 `xml:"width,omitempty"`

	Height int32 `xml:"height,omitempty"`

	View string `xml:"view,omitempty"`
}

type RoutePlannerEUMapResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ routePlannerEUMapResponse"`

	Image []byte `xml:"image,omitempty"`
}

type SepaConvertBankAccountNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaConvertBankAccountNumber"`

	Account_number string `xml:"account_number,omitempty"`

	Nbi string `xml:"nbi,omitempty"`

	Country_code string `xml:"country_code,omitempty"`
}

type SepaConvertBankAccountNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaConvertBankAccountNumberResponse"`

	Response *SepaBankAccountIbanData `xml:"response,omitempty"`
}

type SepaValidateInternationalBankAccountNumberFormatRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaValidateInternationalBankAccountNumberFormat"`

	Iban string `xml:"iban,omitempty"`
}

type SepaValidateInternationalBankAccountNumberFormatResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaValidateInternationalBankAccountNumberFormatResponse"`

	Validation_report *SepaInternationalBankAccountNumberFormatValidationReport `xml:"validation_report,omitempty"`
}

type SepaIbanDetailsRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaIbanDetails"`

	Iban string `xml:"iban,omitempty"`
}

type SepaIbanDetailsResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaIbanDetailsResponse"`

	Iban_details *SepaBankAccountIbanData `xml:"iban_details,omitempty"`
}

type SepaMatchAccountHolderRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaMatchAccountHolder"`

	Iban string `xml:"iban,omitempty"`

	Name string `xml:"name,omitempty"`
}

type SepaMatchAccountHolderResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaMatchAccountHolderResponse"`

	Match_account_holder *SepaMatchAccountHolderResult `xml:"match_account_holder,omitempty"`
}

type SepaConvertBasicBankAccountNumberRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaConvertBasicBankAccountNumber"`

	Bban string `xml:"bban,omitempty"`

	Country_iso string `xml:"country_iso,omitempty"`
}

type SepaConvertBasicBankAccountNumberResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ sepaConvertBasicBankAccountNumberResponse"`

	Bankaccount *SepaBankAccountData `xml:"bankaccount,omitempty"`
}

type VatValidateRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ vatValidate"`

	Vat_number string `xml:"vat_number,omitempty"`
}

type VatValidateResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ vatValidateResponse"`

	Out *VatValidation `xml:"out,omitempty"`
}

type VatViesProxyCheckVatRequestType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ vatViesProxyCheckVat"`

	Vat_number string `xml:"vat_number,omitempty"`
}

type VatViesProxyCheckVatResponseType struct {
	XMLName xml.Name `xml:"http://www.webservices.nl/soap/ vatViesProxyCheckVatResponse"`

	Out *VatProxyViesCheckVatResponse `xml:"out,omitempty"`
}

type WebservicesNLPortType interface {

	/* Retrieve a token with which a new account may be created via the Webview */
	AccountGetCreationToken(request *AccountGetCreationTokenRequestType) (*AccountGetCreationTokenResponseType, error)

	/* Get the id of an account created with a token retrieved using accountGetCreationToken */
	AccountGetCreationStatus(request *AccountGetCreationStatusRequestType) (*AccountGetCreationStatusResponseType, error)

	/* Retrieve a token that can be used order account balance via the <Webview Interface> */
	AccountGetOrderToken(request *AccountGetOrderTokenRequestType) (*AccountGetOrderTokenResponseType, error)

	/* Removes all sessions of a user, or just a single session if reactid is not 0 */
	UserSessionRemove(request *UserSessionRemoveRequestType) (*UserSessionRemoveResponseType, error)

	/* Lists all the current sessions of a user */
	UserSessionList(request *UserSessionListRequestType) (*UserSessionListResponseType, error)

	/* Returns the users balance */
	UserViewBalance(request *UserViewBalanceRequestType) (*UserViewBalanceResponseType, error)

	/* Change the user's balance */
	UserEditBalance(request *UserEditBalanceRequestType) (*UserEditBalanceResponseType, error)

	/* Returns the accounts balance */
	AccountViewBalance(request *AccountViewBalanceRequestType) (*AccountViewBalanceResponseType, error)

	/* View users details */
	UserViewV2(request *UserViewV2RequestType) (*UserViewV2ResponseType, error)

	/* Edit a users details, the users current password is required */
	UserEditV2(request *UserEditV2RequestType) (*UserEditV2ResponseType, error)

	/* Edit details only an admin may change, such as a users nickname and password.notificationrecipients must have one of the following values: accountcontact, user,
	   accountcontact_and_user. Otherwise it will default to accountcontact. */
	UserEditExtendedV2(request *UserEditExtendedV2RequestType) (*UserEditExtendedV2ResponseType, error)

	/* Create and notify a new user */
	UserCreateV2(request *UserCreateV2RequestType) (*UserCreateV2ResponseType, error)

	/* Create a WebservicesNl test user */
	CreateTestUser(request *CreateTestUserRequestType) (*CreateTestUserResponseType, error)

	/* Change a users password */
	UserChangePassword(request *UserChangePasswordRequestType) (*UserChangePasswordResponseType, error)

	/* Deletes a user */
	UserRemove(request *UserRemoveRequestType) (*UserRemoveResponseType, error)

	/* Activates the user and sends user login information to users email address */
	UserNotify(request *UserNotifyRequestType) (*UserNotifyResponseType, error)

	/* Returns a list of all groups the user may assign to other users */
	UserListAssignableGroups(request *UserListAssignableGroupsRequestType) (*UserListAssignableGroupsResponseType, error)

	/* Add a user to a group */
	UserAddGroup(request *UserAddGroupRequestType) (*UserAddGroupResponseType, error)

	/* Remove a user from a group */
	UserRemoveGroup(request *UserRemoveGroupRequestType) (*UserRemoveGroupResponseType, error)

	/* View the account details */
	AccountViewV2(request *AccountViewV2RequestType) (*AccountViewV2ResponseType, error)

	/* Edit an accounts details */
	AccountEditV2(request *AccountEditV2RequestType) (*AccountEditV2ResponseType, error)

	/* Returns a list of all users in the account */
	AccountUserListV2(request *AccountUserListV2RequestType) (*AccountUserListV2ResponseType, error)

	/* Returns a list of users that are in the account and match the search phrase */
	AccountUserSearchV2(request *AccountUserSearchV2RequestType) (*AccountUserSearchV2ResponseType, error)

	/* Sets host restrictions for the account */
	AccountEditHostRestrictions(request *AccountEditHostRestrictionsRequestType) (*AccountEditHostRestrictionsResponseType, error)

	/* Gets host restrictions for the account */
	AccountViewHostRestrictions(request *AccountViewHostRestrictionsRequestType) (*AccountViewHostRestrictionsResponseType, error)

	/* Edit host restrictions for this user */
	UserEditHostRestrictions(request *UserEditHostRestrictionsRequestType) (*UserEditHostRestrictionsResponseType, error)

	/* View host restrictions for this user */
	UserViewHostRestrictions(request *UserViewHostRestrictionsRequestType) (*UserViewHostRestrictionsResponseType, error)

	/* Get an address by providing its complete postcode (including house number). This is a reeks level function.  */
	AddressReeksPostcodeSearch(request *AddressReeksPostcodeSearchRequestType) (*AddressReeksPostcodeSearchResponseType, error)

	/* Get an address by providing the address, postcode, city. This is a reeks level function.  */
	AddressReeksAddressSearch(request *AddressReeksAddressSearchRequestType) (*AddressReeksAddressSearchResponseType, error)

	/* Look for an address by province, municipality, city, street, house number and house number addition. Returns a list of matches. This is a reeks level function.  */
	AddressReeksFullParameterSearch(request *AddressReeksFullParameterSearchRequestType) (*AddressReeksFullParameterSearchResponseType, error)

	/* Look for an address by province, municipality, city, street, house number and house number addition. Returns a list of matches. This is a reeks level function. */
	AddressReeksParameterSearch(request *AddressReeksParameterSearchRequestType) (*AddressReeksParameterSearchResponseType, error)

	/* Look for an perceel address by simply providing a phrase such as "Somestreet 123". Returns a list of matches. This is a perceel level function (more accurate than reeks level). */
	AddressPerceelPhraseSearch(request *AddressPerceelPhraseSearchRequestType) (*AddressPerceelPhraseSearchResponseType, error)

	/* Look for an perceel address by province, municipality, city, street, house number and house number addition. Returns a list of matches This is a perceel level function (more accurate than reeks level).  */
	AddressPerceelFullParameterSearchV2(request *AddressPerceelFullParameterSearchV2RequestType) (*AddressPerceelFullParameterSearchV2ResponseType, error)

	/* Returns a list of all neighborhood codes in the province */
	AddressProvinceListNeighborhoods(request *AddressProvinceListNeighborhoodsRequestType) (*AddressProvinceListNeighborhoodsResponseType, error)

	/* Returns a list of municipalities which are in the specified province */
	AddressProvinceListDistricts(request *AddressProvinceListDistrictsRequestType) (*AddressProvinceListDistrictsResponseType, error)

	/* Returns a list of all provinces */
	AddressProvinceList(request *AddressProvinceListRequestType) (*AddressProvinceListResponseType, error)

	/* Returns a list of provinces which match the given name */
	AddressProvinceSearch(request *AddressProvinceSearchRequestType) (*AddressProvinceSearchResponseType, error)

	/* Returns a list of municipalities which match the given name */
	AddressDistrictSearch(request *AddressDistrictSearchRequestType) (*AddressDistrictSearchResponseType, error)

	/* Returns a list of cities which are in the specified municipality */
	AddressDistrictListCities(request *AddressDistrictListCitiesRequestType) (*AddressDistrictListCitiesResponseType, error)

	/* Returns a list of all neighborhood codes in the municipality */
	AddressDistrictListNeighborhoods(request *AddressDistrictListNeighborhoodsRequestType) (*AddressDistrictListNeighborhoodsResponseType, error)

	/* Search for all cities that match a phrase. Cities are also matched if input matches a commonly used alternative city name. Exact matches on the official name are listed first, the rest of the results is sorted alphabetically. */
	AddressCitySearchV2(request *AddressCitySearchV2RequestType) (*AddressCitySearchV2ResponseType, error)

	/* Returns a list of all neighborhood codes in the city */
	AddressCityListNeighborhoods(request *AddressCityListNeighborhoodsRequestType) (*AddressCityListNeighborhoodsResponseType, error)

	/* This is a deprecated function. Please use "addressPerceelFullParameterSearchV2" instead. */
	AddressPerceelFullParameterSearch(request *AddressPerceelFullParameterSearchRequestType) (*AddressPerceelFullParameterSearchResponseType, error)

	/* This is a deprecated function. Please use "addressPerceelFullParameterSearchV2" or "addressPerceelPhraseSearch" instead. */
	AddressPerceelParameterSearch(request *AddressPerceelParameterSearchRequestType) (*AddressPerceelParameterSearchResponseType, error)

	/* This function is deprecated and will not be available in the future. Use "addressReeksAddressSearch" or "addressReeksFullParameterSearch" instead. */
	AddressReeksPhraseSearch(request *AddressReeksPhraseSearchRequestType) (*AddressReeksPhraseSearchResponseType, error)

	/* This function is about to be removed! */
	AddressNeighborhoodName(request *AddressNeighborhoodNameRequestType) (*AddressNeighborhoodNameResponseType, error)

	/* This function is about to be removed! */
	AddressNeighborhoodPhraseSearch(request *AddressNeighborhoodPhraseSearchRequestType) (*AddressNeighborhoodPhraseSearchResponseType, error)

	/* Get the telephone area codes belonging to the neighborhood */
	AreaCodeLookup(request *AreaCodeLookupRequestType) (*AreaCodeLookupResponseType, error)

	/* Get the neighborhood codes of the neighborhoods that have the specified area code */
	AreaCodeToNeighborhoodcode(request *AreaCodeToNeighborhoodcodeRequestType) (*AreaCodeToNeighborhoodcodeResponseType, error)

	/* Get the telephone areacodes related to a given postcode */
	AreaCodePostcodeLookup(request *AreaCodePostcodeLookupRequestType) (*AreaCodePostcodeLookupResponseType, error)

	/* Logs the user in using username and password, returns a reactid for authentication in further requests */
	Login(request *LoginRequestType) (*LoginResponseType, error)

	/* Logs the user out, destroys the session associated with the reactid */
	Logout(request *LogoutRequestType) (*LogoutResponseType, error)

	/* Retrieve data on a belgian business establishment */
	BelgianBusinessGetDossier(request *BelgianBusinessGetDossierRequestType) (*BelgianBusinessGetDossierResponseType, error)

	/* Retrieve a Bovag member using a Bovag identifier. */
	BovagGetMemberByBovagId(request *BovagGetMemberByBovagIdRequestType) (*BovagGetMemberByBovagIdResponseType, error)

	/* Retrieve a Bovag member using a DutchBusiness identifiers. */
	BovagGetMemberByDutchBusiness(request *BovagGetMemberByDutchBusinessRequestType) (*BovagGetMemberByDutchBusinessResponseType, error)

	/* Retrieve the establishment number (Dutch 'Vestigingsnummer') for a business. */
	BusinessGetEstablishmentNumber(request *BusinessGetEstablishmentNumberRequestType) (*BusinessGetEstablishmentNumberResponseType, error)

	/* Look up a BIK ('Bedrijfsindeling kamers van koophandel') code, which is the type of code returned in a  <BusinessDossier> as the Primary or Secondary activity code. The description for all levels of the BIK code is returned, even if the full length BIK code does not exist. Descriptions of the section associated with the BIK code is also returned. */
	BusinessGetBIKDescription(request *BusinessGetBIKDescriptionRequestType) (*BusinessGetBIKDescriptionResponseType, error)

	/* Look up a SBI ('Standaard Bedrijfs Indeling 2008') code. */
	BusinessGetSBIDescription(request *BusinessGetSBIDescriptionRequestType) (*BusinessGetSBIDescriptionResponseType, error)

	BusinessBIKToSBI(request *BusinessBIKToSBIRequestType) (*BusinessBIKToSBIResponseType, error)

	BusinessSBIToBIK(request *BusinessSBIToBIKRequestType) (*BusinessSBIToBIKResponseType, error)

	/* Look up a business Dossier. The subdossierno parameter can be left empty to get all Sub Dossiers of one Dossier. */
	BusinessGetDossierV3(request *BusinessGetDossierV3RequestType) (*BusinessGetDossierV3ResponseType, error)

	/* Look up an extended business Dossier. */
	BusinessGetDossierExtended(request *BusinessGetDossierExtendedRequestType) (*BusinessGetDossierExtendedResponseType, error)

	/* Look up a business using its dossier number. */
	BusinessSearchDossierNumber(request *BusinessSearchDossierNumberRequestType) (*BusinessSearchDossierNumberResponseType, error)

	/* Look up a business using its postcode address. nbcode is required, other parameters may be left empty. */
	BusinessSearchPostcode(request *BusinessSearchPostcodeRequestType) (*BusinessSearchPostcodeResponseType, error)

	/* Lookup a business using its address. Specified parameters are compared against both establishment and correspondence addresses to find a match. Parameters may be left empty. */
	BusinessSearchAddress(request *BusinessSearchAddressRequestType) (*BusinessSearchAddressResponseType, error)

	/* Lookup a business using its name. */
	BusinessSearchName(request *BusinessSearchNameRequestType) (*BusinessSearchNameResponseType, error)

	/* Look up a business by a number of parameters. At least a tradename, neighborhood code, street, city or telephone number is required. Lookups based on postbus address are possible by specifying streetname "Postbus" and specifying the postbus number in the houseno parameter. Specified address parameters are compared against both establishment and correspondence addresses to find a match.  */
	BusinessSearchParameters(request *BusinessSearchParametersRequestType) (*BusinessSearchParametersResponseType, error)

	/* Look up a business by a number of parameters. At least a tradename, neighborhood code, street, city or telephone number is required. Lookups based on postbus address are possible by specifying streetname "Postbus" and specifying the postbus number in the houseno parameter. Specified address parameters are compared against both establishment and correspondence addresses to find a match.  */
	BusinessSearchParametersV3(request *BusinessSearchParametersV3RequestType) (*BusinessSearchParametersV3ResponseType, error)

	/* Look up a business using its location and SBI code. */
	BusinessSearchSelection(request *BusinessSearchSelectionRequestType) (*BusinessSearchSelectionResponseType, error)

	/* Look up a business using its location and SBI code. */
	BusinessSearchSelectionV2(request *BusinessSearchSelectionV2RequestType) (*BusinessSearchSelectionV2ResponseType, error)

	/* Returns the SBI codes and their descriptions for a dossier. */
	BusinessGetDossierSBI(request *BusinessGetDossierSBIRequestType) (*BusinessGetDossierSBIResponseType, error)

	BusinessUpdateCheckDossier(request *BusinessUpdateCheckDossierRequestType) (*BusinessUpdateCheckDossierResponseType, error)

	/* Retrieve dossier numbers for all dossiers changed since the given date. */
	BusinessUpdateGetChangedDossiers(request *BusinessUpdateGetChangedDossiersRequestType) (*BusinessUpdateGetChangedDossiersResponseType, error)

	/* Returns a list of all dossiers that have been updated since they were last retrieved by the user */
	BusinessUpdateGetDossiers(request *BusinessUpdateGetDossiersRequestType) (*BusinessUpdateGetDossiersResponseType, error)

	/* Add a dossier to the list of dossiers for which the user (the user whose credentials are used to make the call) wants to receive updates */
	BusinessUpdateAddDossier(request *BusinessUpdateAddDossierRequestType) (*BusinessUpdateAddDossierResponseType, error)

	/* Remove a dossier from the list of dossiers for which the user (the user whose credentials are used to make the call) wants to receive updates. */
	BusinessUpdateRemoveDossier(request *BusinessUpdateRemoveDossierRequestType) (*BusinessUpdateRemoveDossierResponseType, error)

	/* Look up a business by a number of parameters. At least a tradename, neighborhood code, street, city or telephone number is required. Lookups based on postbus address are possible by specifying streetname "Postbus" and specifying the postbus number in the houseno parameter. Specified address parameters are compared against both establishment and correspondence addresses to find a match.  */
	BusinessSearchParametersV2(request *BusinessSearchParametersV2RequestType) (*BusinessSearchParametersV2ResponseType, error)

	/* Check the validity of a license plate and check code ('meldcode') combination */
	CarRDWCarCheckCode(request *CarRDWCarCheckCodeRequestType) (*CarRDWCarCheckCodeResponseType, error)

	/* Retrieves data of a car with a Dutch license plate, including a list of types matched if more information is available. */
	CarRDWCarDataV3(request *CarRDWCarDataV3RequestType) (*CarRDWCarDataV3ResponseType, error)

	/* Find a specific Dutch car's data, including BPM and power and types */
	CarRDWCarDataBPV2(request *CarRDWCarDataBPV2RequestType) (*CarRDWCarDataBPV2ResponseType, error)

	/* Find a specific Dutch car's data, including European Approval Mark */
	CarRDWCarDataExtended(request *CarRDWCarDataExtendedRequestType) (*CarRDWCarDataExtendedResponseType, error)

	/* Retrieves data of a car with a Dutch license plate, including catalog price. */
	CarRDWCarDataPrice(request *CarRDWCarDataPriceRequestType) (*CarRDWCarDataPriceResponseType, error)

	/* Find a specific Dutch car's data, including information about extra options */
	CarRDWCarDataOptions(request *CarRDWCarDataOptionsRequestType) (*CarRDWCarDataOptionsResponseType, error)

	/* Check the validity of a license plate and check code ('meldcode') combination */
	CarVWEMeldcodeCheck(request *CarVWEMeldcodeCheckRequestType) (*CarVWEMeldcodeCheckResponseType, error)

	CarVWEBasicTypeData(request *CarVWEBasicTypeDataRequestType) (*CarVWEBasicTypeDataResponseType, error)

	CarVWEVersionPrice(request *CarVWEVersionPriceRequestType) (*CarVWEVersionPriceResponseType, error)

	/* Retrieve options of a specific model of a car */
	CarVWEOptions(request *CarVWEOptionsRequestType) (*CarVWEOptionsResponseType, error)

	CarVWEListBrands(request *CarVWEListBrandsRequestType) (*CarVWEListBrandsResponseType, error)

	CarVWEListModels(request *CarVWEListModelsRequestType) (*CarVWEListModelsResponseType, error)

	CarVWEListVersions(request *CarVWEListVersionsRequestType) (*CarVWEListVersionsResponseType, error)

	CarVWEVersionYearData(request *CarVWEVersionYearDataRequestType) (*CarVWEVersionYearDataResponseType, error)

	/* Retrieve photos of a specific model of a car */
	CarVWEPhotos(request *CarVWEPhotosRequestType) (*CarVWEPhotosResponseType, error)

	/* Retrieve Autodisk price information */
	CarATDPrice(request *CarATDPriceRequestType) (*CarATDPriceResponseType, error)

	/* Find a specific Dutch car's data */
	CarRDWCarData(request *CarRDWCarDataRequestType) (*CarRDWCarDataResponseType, error)

	/* Retrieves data of a car with a Dutch license plate, including a list of types matched if more information is available. */
	CarRDWCarDataV2(request *CarRDWCarDataV2RequestType) (*CarRDWCarDataV2ResponseType, error)

	/* Find a specific Dutch car's data, including BPM and power */
	CarRDWCarDataBP(request *CarRDWCarDataBPRequestType) (*CarRDWCarDataBPResponseType, error)

	/* Retrieve a list of person entities based on search criteria */
	ComplianceSearchPersons(request *ComplianceSearchPersonsRequestType) (*ComplianceSearchPersonsResponseType, error)

	/* Retrieve a person entity based on a provided id */
	ComplianceGetPerson(request *ComplianceGetPersonRequestType) (*ComplianceGetPersonResponseType, error)

	/* Retrieve a list of business entities based on search criteria */
	ComplianceSearchBusinesses(request *ComplianceSearchBusinessesRequestType) (*ComplianceSearchBusinessesResponseType, error)

	/* Retrieve a business based on a provided id */
	ComplianceGetBusiness(request *ComplianceGetBusinessRequestType) (*ComplianceGetBusinessResponseType, error)

	CreditsafeGetReportFull(request *CreditsafeGetReportFullRequestType) (*CreditsafeGetReportFullResponseType, error)

	CreditsafeSearch(request *CreditsafeSearchRequestType) (*CreditsafeSearchResponseType, error)

	CreditsafeSearchV2(request *CreditsafeSearchV2RequestType) (*CreditsafeSearchV2ResponseType, error)

	CreditsafeGetReportFullV2(request *CreditsafeGetReportFullV2RequestType) (*CreditsafeGetReportFullV2ResponseType, error)

	CreditsafeSearchCriteria(request *CreditsafeSearchCriteriaRequestType) (*CreditsafeSearchCriteriaResponseType, error)

	CreditsafeMonitorAddCompany(request *CreditsafeMonitorAddCompanyRequestType) (*CreditsafeMonitorAddCompanyResponseType, error)

	CreditsafeMonitorRemoveCompany(request *CreditsafeMonitorRemoveCompanyRequestType) (*CreditsafeMonitorRemoveCompanyResponseType, error)

	CreditsafeMonitorGetUpdatedCompanies(request *CreditsafeMonitorGetUpdatedCompaniesRequestType) (*CreditsafeMonitorGetUpdatedCompaniesResponseType, error)

	/* Look up a business and return basic information about the business. */
	DnbSearchReferenceV2(request *DnbSearchReferenceV2RequestType) (*DnbSearchReferenceV2ResponseType, error)

	/* Retrieve a DNBBusinessReference with name, address and D&B business key for a specific business. */
	DnbGetReference(request *DnbGetReferenceRequestType) (*DnbGetReferenceResponseType, error)

	/* Retrieve basic WorldBase business information */
	DnbWorldbaseMarketing(request *DnbWorldbaseMarketingRequestType) (*DnbWorldbaseMarketingResponseType, error)

	/* Retrieve detailed WorldBase business information */
	DnbWorldbaseMarketingPlus(request *DnbWorldbaseMarketingPlusRequestType) (*DnbWorldbaseMarketingPlusResponseType, error)

	/* Detailed WorldBase information, including information on a business' family tree */
	DnbWorldbaseMarketingPlusLinkage(request *DnbWorldbaseMarketingPlusLinkageRequestType) (*DnbWorldbaseMarketingPlusLinkageResponseType, error)

	/* Do a D&B Quick Check on a business. */
	DnbQuickCheck(request *DnbQuickCheckRequestType) (*DnbQuickCheckResponseType, error)

	/* Do a D&B Business Verification on a business. */
	DnbBusinessVerification(request *DnbBusinessVerificationRequestType) (*DnbBusinessVerificationResponseType, error)

	/* Retrieve D&B Enterprise Management for a business. */
	DnbEnterpriseManagement(request *DnbEnterpriseManagementRequestType) (*DnbEnterpriseManagementResponseType, error)

	/* Look up a business and return basic information about the business. */
	DnbSearchReference(request *DnbSearchReferenceRequestType) (*DnbSearchReferenceResponseType, error)

	/* Given two neighborhood codes, the drive distance in meters between the neighborhoodsis returned, for both the fastest and the shortest route */
	DriveInfoDistanceLookup(request *DriveInfoDistanceLookupRequestType) (*DriveInfoDistanceLookupResponseType, error)

	/* Given two neighborhood codes, the drivetime in minutes between the neighborhoods is returned, for both the fastest and the shortest route */
	DriveInfoTimeLookup(request *DriveInfoTimeLookupRequestType) (*DriveInfoTimeLookupResponseType, error)

	/* Validate and complete an address based on postcode and house number. */
	DutchAddressRangePostcodeSearch(request *DutchAddressRangePostcodeSearchRequestType) (*DutchAddressRangePostcodeSearchResponseType, error)

	/* Retrieve data on a business establishment */
	DutchBusinessGetDossier(request *DutchBusinessGetDossierRequestType) (*DutchBusinessGetDossierResponseType, error)

	/* Retrieve data on a business establishment */
	DutchBusinessGetDossierV2(request *DutchBusinessGetDossierV2RequestType) (*DutchBusinessGetDossierV2ResponseType, error)

	/* Retrieve data on a business establishment */
	DutchBusinessGetDossierV3(request *DutchBusinessGetDossierV3RequestType) (*DutchBusinessGetDossierV3ResponseType, error)

	/* Retrieve SBI Information on a company */
	DutchBusinessGetSBI(request *DutchBusinessGetSBIRequestType) (*DutchBusinessGetSBIResponseType, error)

	/* Retrieve data on a business establishment */
	DutchBusinessGetVatNumber(request *DutchBusinessGetVatNumberRequestType) (*DutchBusinessGetVatNumberResponseType, error)

	/* Retrieve position data on a business establishment */
	DutchBusinessGetPositions(request *DutchBusinessGetPositionsRequestType) (*DutchBusinessGetPositionsResponseType, error)

	/* Retrieve legal entity data on a business establishment */
	DutchBusinessGetLegalEntity(request *DutchBusinessGetLegalEntityRequestType) (*DutchBusinessGetLegalEntityResponseType, error)

	/* Retrieve the organization tree of a company */
	DutchBusinessGetOrganizationTree(request *DutchBusinessGetOrganizationTreeRequestType) (*DutchBusinessGetOrganizationTreeResponseType, error)

	/* Determine if a Company has non-marketing indicator */
	DutchBusinessGetNonMarketingIndicator(request *DutchBusinessGetNonMarketingIndicatorRequestType) (*DutchBusinessGetNonMarketingIndicatorResponseType, error)

	/* Search for business establishments using a known identifier */
	DutchBusinessSearchDossierNumber(request *DutchBusinessSearchDossierNumberRequestType) (*DutchBusinessSearchDossierNumberResponseType, error)

	/* Search for an overview of Corporate Group Relationships aka 'concern relaties' for specified dossier number */
	DutchBusinessGetConcernRelationsOverview(request *DutchBusinessGetConcernRelationsOverviewRequestType) (*DutchBusinessGetConcernRelationsOverviewResponseType, error)

	/* Get all known details of a Corporate Group Relationships (concernrelaties) for specified dossier number */
	DutchBusinessGetConcernRelationsDetails(request *DutchBusinessGetConcernRelationsDetailsRequestType) (*DutchBusinessGetConcernRelationsDetailsResponseType, error)

	/* Find business establishments using a variety of parameters */
	DutchBusinessSearchParametersV2(request *DutchBusinessSearchParametersV2RequestType) (*DutchBusinessSearchParametersV2ResponseType, error)

	/* Find business establishments using a variety of parameters */
	DutchBusinessSearch(request *DutchBusinessSearchRequestType) (*DutchBusinessSearchResponseType, error)

	/* Search for business establishments using a known identifier */
	DutchBusinessSearchEstablishments(request *DutchBusinessSearchEstablishmentsRequestType) (*DutchBusinessSearchEstablishmentsResponseType, error)

	/* Find business establishments based on postcode and house number */
	DutchBusinessSearchPostcode(request *DutchBusinessSearchPostcodeRequestType) (*DutchBusinessSearchPostcodeResponseType, error)

	/* Search for businesses matching all of the given criteria. */
	DutchBusinessSearchSelection(request *DutchBusinessSearchSelectionRequestType) (*DutchBusinessSearchSelectionResponseType, error)

	/* Simple search for businesses matching all of the given criteria. */
	DutchBusinessSearchSelectionSimple(request *DutchBusinessSearchSelectionSimpleRequestType) (*DutchBusinessSearchSelectionSimpleResponseType, error)

	/* Look up a SBI ('Standaard Bedrijfs Indeling 2008') code. */
	DutchBusinessGetSBIDescription(request *DutchBusinessGetSBIDescriptionRequestType) (*DutchBusinessGetSBIDescriptionResponseType, error)

	/* Get an annual financial statement for a dutch-business */
	DutchBusinessGetAnnualFinancialStatement(request *DutchBusinessGetAnnualFinancialStatementRequestType) (*DutchBusinessGetAnnualFinancialStatementResponseType, error)

	/* Retrieves the positions, current and historical, of a functionary. */
	DutchBusinessGetAdditionalPositions(request *DutchBusinessGetAdditionalPositionsRequestType) (*DutchBusinessGetAdditionalPositionsResponseType, error)

	/* Finds similar companies (dossier numbers) for a given company */
	DutchBusinessLookALikes(request *DutchBusinessLookALikesRequestType) (*DutchBusinessLookALikesResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractDocument(request *DutchBusinessGetExtractDocumentRequestType) (*DutchBusinessGetExtractDocumentResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractDocumentData(request *DutchBusinessGetExtractDocumentDataRequestType) (*DutchBusinessGetExtractDocumentDataResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractDocumentDataV2(request *DutchBusinessGetExtractDocumentDataV2RequestType) (*DutchBusinessGetExtractDocumentDataV2ResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractDocumentDataV3(request *DutchBusinessGetExtractDocumentDataV3RequestType) (*DutchBusinessGetExtractDocumentDataV3ResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetLegalExtractDocumentDataV2(request *DutchBusinessGetLegalExtractDocumentDataV2RequestType) (*DutchBusinessGetLegalExtractDocumentDataV2ResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetLegalExtractDocumentDataV3(request *DutchBusinessGetLegalExtractDocumentDataV3RequestType) (*DutchBusinessGetLegalExtractDocumentDataV3ResponseType, error)

	/* Retrieve a list of extract document references */
	DutchBusinessGetExtractHistory(request *DutchBusinessGetExtractHistoryRequestType) (*DutchBusinessGetExtractHistoryResponseType, error)

	/* Retrieve a list of extract document references */
	DutchBusinessGetExtractHistoryChanged(request *DutchBusinessGetExtractHistoryChangedRequestType) (*DutchBusinessGetExtractHistoryChangedResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractHistoryDocumentData(request *DutchBusinessGetExtractHistoryDocumentDataRequestType) (*DutchBusinessGetExtractHistoryDocumentDataResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractHistoryDocumentDataV2(request *DutchBusinessGetExtractHistoryDocumentDataV2RequestType) (*DutchBusinessGetExtractHistoryDocumentDataV2ResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractHistoryDocumentDataV3(request *DutchBusinessGetExtractHistoryDocumentDataV3RequestType) (*DutchBusinessGetExtractHistoryDocumentDataV3ResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	DutchBusinessGetExtractHistoryDocumentDataV3ByDossier(request *DutchBusinessGetExtractHistoryDocumentDataV3ByDossierRequestType) (*DutchBusinessGetExtractHistoryDocumentDataV3ByDossierResponseType, error)

	/* Get a list of logged updates for a specific business dossier */
	DutchBusinessGetDossierHistory(request *DutchBusinessGetDossierHistoryRequestType) (*DutchBusinessGetDossierHistoryResponseType, error)

	/* Returns a list of all dossiers (that have been added to the users follow list) thathave been updated since they were last retrieved by the user (the user whose credentials are used to make the call). */
	DutchBusinessUpdateGetDossiers(request *DutchBusinessUpdateGetDossiersRequestType) (*DutchBusinessUpdateGetDossiersResponseType, error)

	/* Retrieve information on the last change to a company. */
	DutchBusinessUpdateCheckDossier(request *DutchBusinessUpdateCheckDossierRequestType) (*DutchBusinessUpdateCheckDossierResponseType, error)

	/* Retrieve dossier numbers for all companies that changed since the given date. */
	DutchBusinessUpdateGetChangedDossiers(request *DutchBusinessUpdateGetChangedDossiersRequestType) (*DutchBusinessUpdateGetChangedDossiersResponseType, error)

	/* Add a company to the follow list for which the user (the user whose credentials are used to make the call) wants to receive updates. */
	DutchBusinessUpdateAddDossier(request *DutchBusinessUpdateAddDossierRequestType) (*DutchBusinessUpdateAddDossierResponseType, error)

	/* Remove a company from the follow list for which the user (the user whose credentials are used to make the call) wants to receive updates. */
	DutchBusinessUpdateRemoveDossier(request *DutchBusinessUpdateRemoveDossierRequestType) (*DutchBusinessUpdateRemoveDossierResponseType, error)

	/* Get a complete overview of all the companies that are on the follow list for the user (the user whose credentials are used to make the call). */
	DutchBusinessUpdateGetOverview(request *DutchBusinessUpdateGetOverviewRequestType) (*DutchBusinessUpdateGetOverviewResponseType, error)

	/* Get a paged result of all news items found. */
	DutchBusinessSearchNewsByDossier(request *DutchBusinessSearchNewsByDossierRequestType) (*DutchBusinessSearchNewsByDossierResponseType, error)

	/* Starts a UBO investigation for given dossierNumber and establishmentNumber */
	DutchBusinessUBOStartInvestigation(request *DutchBusinessUBOStartInvestigationRequestType) (*DutchBusinessUBOStartInvestigationResponseType, error)

	/* Checks the status of a UBO investigation */
	DutchBusinessUBOCheckInvestigation(request *DutchBusinessUBOCheckInvestigationRequestType) (*DutchBusinessUBOCheckInvestigationResponseType, error)

	/* Pickup the result of a UBO investigation */
	DutchBusinessUBOPickupInvestigation(request *DutchBusinessUBOPickupInvestigationRequestType) (*DutchBusinessUBOPickupInvestigationResponseType, error)

	/* Return the costs of a finished UBO investigation */
	DutchBusinessUBOCostsInvestigation(request *DutchBusinessUBOCostsInvestigationRequestType) (*DutchBusinessUBOCostsInvestigationResponseType, error)

	/* Labels the natural persons found in an UBO Investigation according to a set of rules. */
	DutchBusinessUBOClassifyInvestigation(request *DutchBusinessUBOClassifyInvestigationRequestType) (*DutchBusinessUBOClassifyInvestigationResponseType, error)

	/* Retrieve Lei information for a business using the Chamber of Commerce number */
	DutchBusinessGetLei(request *DutchBusinessGetLeiRequestType) (*DutchBusinessGetLeiResponseType, error)

	/* Find business establishments using a variety of parameters */
	DutchBusinessSearchParameters(request *DutchBusinessSearchParametersRequestType) (*DutchBusinessSearchParametersResponseType, error)

	/* Get the information of a dutch vehicle */
	DutchVehicleGetVehicleV2(request *DutchVehicleGetVehicleV2RequestType) (*DutchVehicleGetVehicleV2ResponseType, error)

	/* Get the price information of a dutch vehicle */
	DutchVehicleGetPurchaseReference(request *DutchVehicleGetPurchaseReferenceRequestType) (*DutchVehicleGetPurchaseReferenceResponseType, error)

	/* Get all the (previous) owners of a dutch vehicle */
	DutchVehicleGetOwnerHistory(request *DutchVehicleGetOwnerHistoryRequestType) (*DutchVehicleGetOwnerHistoryResponseType, error)

	/* Get all the current market value of a vehicle */
	DutchVehicleGetMarketValue(request *DutchVehicleGetMarketValueRequestType) (*DutchVehicleGetMarketValueResponseType, error)

	/* Get the information of a dutch vehicle */
	DutchVehicleGetVehicle(request *DutchVehicleGetVehicleRequestType) (*DutchVehicleGetVehicleResponseType, error)

	/* Provides a credit score for a person identified by a set of parameters. */
	EdrGetScore(request *EdrGetScoreRequestType) (*EdrGetScoreResponseType, error)

	/* Returns the coordinates in the RD system of the neighborhood, given the neighborhood code */
	GeoLocationNeighborhoodCoordinatesRD(request *GeoLocationNeighborhoodCoordinatesRDRequestType) (*GeoLocationNeighborhoodCoordinatesRDResponseType, error)

	/* Returns the coordinates of the given postcode in the RD system.  */
	GeoLocationPostcodeCoordinatesRD(request *GeoLocationPostcodeCoordinatesRDRequestType) (*GeoLocationPostcodeCoordinatesRDResponseType, error)

	/* Returns the coordinates in degrees of latitude/longitude of the neighborhood, given the neighborhood code */
	GeoLocationNeighborhoodCoordinatesLatLon(request *GeoLocationNeighborhoodCoordinatesLatLonRequestType) (*GeoLocationNeighborhoodCoordinatesLatLonResponseType, error)

	/* Returns the coordinates of the given postcode in degrees of latitude/longitude. */
	GeoLocationPostcodeCoordinatesLatLon(request *GeoLocationPostcodeCoordinatesLatLonRequestType) (*GeoLocationPostcodeCoordinatesLatLonResponseType, error)

	/* Returns the coordinates of the given address in degrees of latitude/longitude. Address may be specified by postcode + house number or by city + street + house number.  */
	GeoLocationAddressCoordinatesLatLon(request *GeoLocationAddressCoordinatesLatLonRequestType) (*GeoLocationAddressCoordinatesLatLonResponseType, error)

	/* Returns the coordinates of the given address in the RD system. Address may be specified by postcode + house number or by city + street + house number.  */
	GeoLocationAddressCoordinatesRD(request *GeoLocationAddressCoordinatesRDRequestType) (*GeoLocationAddressCoordinatesRDResponseType, error)

	/* Returns the postcode of the address closest to the specified latitude/longitude coordinate in the Netherlands. */
	GeoLocationLatLonToPostcode(request *GeoLocationLatLonToPostcodeRequestType) (*GeoLocationLatLonToPostcodeResponseType, error)

	/* Returns the address closest to the specified latitude/longitude coordinate in the Netherlands. The returned x, y, latitude and logitude values are the coordinates of the retrieved address. Distance is the distance between these coordinates and the input coordinate in meters */
	GeoLocationLatLonToAddressV2(request *GeoLocationLatLonToAddressV2RequestType) (*GeoLocationLatLonToAddressV2ResponseType, error)

	/* Returns the postcode of the address closest to the specified Rijksdriehoeksmeting coordinate in the Netherlands. */
	GeoLocationRDToPostcode(request *GeoLocationRDToPostcodeRequestType) (*GeoLocationRDToPostcodeResponseType, error)

	/* Returns the address closest to the specified Rijksdriehoeksmeting coordinate in the Netherlands. The returned x, y, latitude and logitude values are the coordinates of the retrieved address. Distance is the distance between these coordinates and the input coordinate in meters */
	GeoLocationRDToAddressV2(request *GeoLocationRDToAddressV2RequestType) (*GeoLocationRDToAddressV2ResponseType, error)

	/* Returns the coordinates of the given postcode in degrees of latitude/longitude. Most countries are supported in the function. Accuracy of the result may vary between countries.  */
	GeoLocationInternationalPostcodeCoordinatesLatLon(request *GeoLocationInternationalPostcodeCoordinatesLatLonRequestType) (*GeoLocationInternationalPostcodeCoordinatesLatLonResponseType, error)

	/* Returns the coordinates of the given address in degrees of latitude/longitude. Most countries are supported in the function. Accuracy of the result may vary between countries. Since the street and city have to contain the complete name and since this method acts with international data, we recommend to use <geoLocationInternationalPostcodeCoordinatesLatLon> if you know the postcode, since working with postcodes is less error prone. */
	GeoLocationInternationalAddressCoordinatesLatLon(request *GeoLocationInternationalAddressCoordinatesLatLonRequestType) (*GeoLocationInternationalAddressCoordinatesLatLonResponseType, error)

	/* Returns the coordinates of the given address in degrees of latitude/longitude. Most countries are supported in the function. Accuracy of the result may vary between countries. */
	GeoLocationInternationalAddressCoordinatesLatLonV2(request *GeoLocationInternationalAddressCoordinatesLatLonV2RequestType) (*GeoLocationInternationalAddressCoordinatesLatLonV2ResponseType, error)

	/* Returns the address closest to the specified latitude/longitude coordinate. The returned latitude and logitude values make up the coordinate of the retrieved address. Distance is the distance between this coordinate and the input coordinate in meters */
	GeoLocationInternationalLatLonToAddress(request *GeoLocationInternationalLatLonToAddressRequestType) (*GeoLocationInternationalLatLonToAddressResponseType, error)

	/* Returns the estimated distance in meters (in a direct line) between the two neighborhoods, given the neighborhood codes */
	GeoLocationNeighborhoodDistance(request *GeoLocationNeighborhoodDistanceRequestType) (*GeoLocationNeighborhoodDistanceResponseType, error)

	/* Returns the estimated distance in meters (in a direct line) between the two postcodes. */
	GeoLocationPostcodeDistance(request *GeoLocationPostcodeDistanceRequestType) (*GeoLocationPostcodeDistanceResponseType, error)

	/* Returns the distance in meters (in a direct line) between two latitude/longitude coordinates */
	GeoLocationHaversineDistance(request *GeoLocationHaversineDistanceRequestType) (*GeoLocationHaversineDistanceResponseType, error)

	/* Returns the neighborhood codes list sorted in order of increasing distance from a given neighborhood. */
	GeoLocationDistanceSortedNeighborhoodCodes(request *GeoLocationDistanceSortedNeighborhoodCodesRequestType) (*GeoLocationDistanceSortedNeighborhoodCodesResponseType, error)

	/* Returns the neighborhood codes list sorted in order of increasing distance from a given neighborhood, within a given radius (in meters). */
	GeoLocationDistanceSortedNeighborhoodCodesRadius(request *GeoLocationDistanceSortedNeighborhoodCodesRadiusRequestType) (*GeoLocationDistanceSortedNeighborhoodCodesRadiusResponseType, error)

	/* Returns a list of postcodes sorted in order of increasing distance from a given postcode, within a given radius (in meters). If the radius is larger than 1500 meters, the result will be based on neighborhood codes */
	GeoLocationDistanceSortedPostcodesRadius(request *GeoLocationDistanceSortedPostcodesRadiusRequestType) (*GeoLocationDistanceSortedPostcodesRadiusResponseType, error)

	/* Returns a given postcode list sorted in order of increasingdistance from a given postcode. */
	GeoLocationDistanceSortedPostcodes(request *GeoLocationDistanceSortedPostcodesRequestType) (*GeoLocationDistanceSortedPostcodesResponseType, error)

	/* Convert a latitude/longitude coordinate to a RD ('Rijksdriehoeksmeting') coordinate. */
	GeoLocationLatLonToRD(request *GeoLocationLatLonToRDRequestType) (*GeoLocationLatLonToRDResponseType, error)

	/* Convert a RD ('Rijksdriehoeksmeting') coordinate to a latitude/longitude coordinate. */
	GeoLocationRDToLatLon(request *GeoLocationRDToLatLonRequestType) (*GeoLocationRDToLatLonResponseType, error)

	/* Retrieve a Graydon credit report of a company registered in the Netherlands. */
	GraydonCreditGetReport(request *GraydonCreditGetReportRequestType) (*GraydonCreditGetReportResponseType, error)

	/* Search international Graydon credit report databases for a company using an identifier. */
	GraydonCreditSearchIdentification(request *GraydonCreditSearchIdentificationRequestType) (*GraydonCreditSearchIdentificationResponseType, error)

	/* Search international Graydon credit report database for a company by name. */
	GraydonCreditSearchName(request *GraydonCreditSearchNameRequestType) (*GraydonCreditSearchNameResponseType, error)

	/* Search international Graydon credit report database for a company using postcode. */
	GraydonCreditSearchPostcode(request *GraydonCreditSearchPostcodeRequestType) (*GraydonCreditSearchPostcodeResponseType, error)

	/* Retrieve a Graydon pd ratings and credit flag of a company registered in the Netherlands. */
	GraydonCreditQuickscan(request *GraydonCreditQuickscanRequestType) (*GraydonCreditQuickscanResponseType, error)

	/* Retrieve various Graydon credit ratings of a company registered in the Netherlands. */
	GraydonCreditRatings(request *GraydonCreditRatingsRequestType) (*GraydonCreditRatingsResponseType, error)

	/* Retrieve the BTW (VAT) number of a company registered in the Netherlands. */
	GraydonCreditVatNumber(request *GraydonCreditVatNumberRequestType) (*GraydonCreditVatNumberResponseType, error)

	/* Retrieve top-parent, parent and sibling companies of a company registered in the Netherlands. */
	GraydonCreditCompanyLiaisons(request *GraydonCreditCompanyLiaisonsRequestType) (*GraydonCreditCompanyLiaisonsResponseType, error)

	/* Retrieve company management credit scores of a company registered in the Netherlands. */
	GraydonCreditManagement(request *GraydonCreditManagementRequestType) (*GraydonCreditManagementResponseType, error)

	/* Fetch a insolvency case from a given publication. */
	InsolvencyGetCaseByPublication(request *InsolvencyGetCaseByPublicationRequestType) (*InsolvencyGetCaseByPublicationResponseType, error)

	/* Search for insolvency publication for a business */
	InsolvencySearchPublicationsByCoCNumber(request *InsolvencySearchPublicationsByCoCNumberRequestType) (*InsolvencySearchPublicationsByCoCNumberResponseType, error)

	/* Search for insolvency publication for a person */
	InsolvencySearchPublicationsByPerson(request *InsolvencySearchPublicationsByPersonRequestType) (*InsolvencySearchPublicationsByPersonResponseType, error)

	/* Returns up to 20 suggestions of valid addresses using the partial or full address data given. Returns more data than internationalAddressSearch. */
	InternationalAddressSearchV2(request *InternationalAddressSearchV2RequestType) (*InternationalAddressSearchV2ResponseType, error)

	/* Returns up to 20 suggestions of valid addresses using the partial or full address data given. Returns more data than internationalAddressSearch. */
	InternationalAddressSearchInteractive(request *InternationalAddressSearchInteractiveRequestType) (*InternationalAddressSearchInteractiveResponseType, error)

	/* Find a 'Eigendomsbericht' by parcel details. Returns the result in a file of the specified format. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterEigendomsBerichtDocumentPerceel(request *KadasterEigendomsBerichtDocumentPerceelRequestType) (*KadasterEigendomsBerichtDocumentPerceelResponseType, error)

	/* Find a 'Eigendomsbericht' by postcode and house number. Returns the result in a file of the specified format. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterEigendomsBerichtDocumentPostcode(request *KadasterEigendomsBerichtDocumentPostcodeRequestType) (*KadasterEigendomsBerichtDocumentPostcodeResponseType, error)

	/* Find a 'Eigendomsbericht' by parcel details. Returns the result as a <BerichtObjectResultaatV2>. In addition to the structured result, a file of the specified format is returned. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterEigendomsBerichtPerceelV2(request *KadasterEigendomsBerichtPerceelV2RequestType) (*KadasterEigendomsBerichtPerceelV2ResponseType, error)

	/* Find a 'Eigendomsbericht' by postcode and house number. Returns the result as a <BerichtObjectResultaatV2>. In addition to the structured result, a file of the specified format is returned. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterEigendomsBerichtPostcodeV2(request *KadasterEigendomsBerichtPostcodeV2RequestType) (*KadasterEigendomsBerichtPostcodeV2ResponseType, error)

	/* Returns a koopsommenoverzicht (in English: real estate transactions overview), which is a list of all transactions of the past five years in the given postcode range. */
	KadasterKoopsommenOverzichtV2(request *KadasterKoopsommenOverzichtV2RequestType) (*KadasterKoopsommenOverzichtV2ResponseType, error)

	/* Returns a 'Uittreksel Kadastrale Kaart' (a map showing the boundaries of a parcel and buildings on that parcel) in the specified format. */
	KadasterUittrekselKadastraleKaartPerceelV3(request *KadasterUittrekselKadastraleKaartPerceelV3RequestType) (*KadasterUittrekselKadastraleKaartPerceelV3ResponseType, error)

	/* Returns a 'Uittreksel Kadastrale Kaart' (a map showing the boundaries of a parcel and buildings on that parcel) in the specified format. */
	KadasterUittrekselKadastraleKaartPostcodeV3(request *KadasterUittrekselKadastraleKaartPostcodeV3RequestType) (*KadasterUittrekselKadastraleKaartPostcodeV3ResponseType, error)

	/* Find a 'Hypothecair bericht' by postcode and house number. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterHypothecairBerichtPostcodeV3(request *KadasterHypothecairBerichtPostcodeV3RequestType) (*KadasterHypothecairBerichtPostcodeV3ResponseType, error)

	/* Find a 'Hypothecair bericht' by parcel details.  */
	KadasterHypothecairBerichtPerceelV3(request *KadasterHypothecairBerichtPerceelV3RequestType) (*KadasterHypothecairBerichtPerceelV3ResponseType, error)

	/* Find a 'brondocument', a document which is the basis for an ascription. */
	KadasterBronDocument(request *KadasterBronDocumentRequestType) (*KadasterBronDocumentResponseType, error)

	KadasterValueListGetRanges(request *KadasterValueListGetRangesRequestType) (*KadasterValueListGetRangesResponseType, error)

	KadasterValueListGetMunicipalities(request *KadasterValueListGetMunicipalitiesRequestType) (*KadasterValueListGetMunicipalitiesResponseType, error)

	/* Returns the coordinates of the given address in both the RD system and the latitude/longitude system. The lat/lon coordinates are derived from the RD coordinates.The address may be specified by giving the postcode, house number & house number addition or by giving the cityname, streetname, house number & house number addition.  */
	KadasterAddressCoordinates(request *KadasterAddressCoordinatesRequestType) (*KadasterAddressCoordinatesResponseType, error)

	/* Returns a 'Kadata WMS De Kadastrale Kaart' map in the specified format. */
	KadasterKadastraleKaartPerceelV2(request *KadasterKadastraleKaartPerceelV2RequestType) (*KadasterKadastraleKaartPerceelV2ResponseType, error)

	/* Returns a 'Kadata WMS De Kadastrale Kaart' map in the specified format. */
	KadasterKadastraleKaartPostcodeV2(request *KadasterKadastraleKaartPostcodeV2RequestType) (*KadasterKadastraleKaartPostcodeV2ResponseType, error)

	/* This function is deprecated and will not be available in the future. Use "kadasterUittrekselKadastraleKaartPerceelV2" instead. */
	KadasterUittrekselKadastraleKaartPerceel(request *KadasterUittrekselKadastraleKaartPerceelRequestType) (*KadasterUittrekselKadastraleKaartPerceelResponseType, error)

	/* This function is deprecated and will not be available in the future. Use "kadasterUittrekselKadastraleKaartPostcodeV2" instead. */
	KadasterUittrekselKadastraleKaartPostcode(request *KadasterUittrekselKadastraleKaartPostcodeRequestType) (*KadasterUittrekselKadastraleKaartPostcodeResponseType, error)

	/* Find a 'Hypothecair bericht' by postcode and house number. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterHypothecairBerichtPostcode(request *KadasterHypothecairBerichtPostcodeRequestType) (*KadasterHypothecairBerichtPostcodeResponseType, error)

	/* Find a 'Hypothecair bericht' by parcel details.  */
	KadasterHypothecairBerichtPerceel(request *KadasterHypothecairBerichtPerceelRequestType) (*KadasterHypothecairBerichtPerceelResponseType, error)

	/* Find a 'Eigendomsbericht' by parcel details. Returns the result as a <BerichtObjectResultaat>. In addition to the structured result, a file of the specified format is returned. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterEigendomsBerichtPerceel(request *KadasterEigendomsBerichtPerceelRequestType) (*KadasterEigendomsBerichtPerceelResponseType, error)

	/* Find a 'Eigendomsbericht' by postcode and house number. Returns the result as a <BerichtObjectResultaat>. In addition to the structured result, a file of the specified format is returned. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterEigendomsBerichtPostcode(request *KadasterEigendomsBerichtPostcodeRequestType) (*KadasterEigendomsBerichtPostcodeResponseType, error)

	/* Returns a 'Uittreksel Kadastrale Kaart' (a map showing the boundaries of a parcel and buildings on that parcel) in the specified format. */
	KadasterUittrekselKadastraleKaartPerceelV2(request *KadasterUittrekselKadastraleKaartPerceelV2RequestType) (*KadasterUittrekselKadastraleKaartPerceelV2ResponseType, error)

	/* Returns a 'Uittreksel Kadastrale Kaart' (a map showing the boundaries of a parcel and buildings on that parcel) in the specified format. */
	KadasterUittrekselKadastraleKaartPostcodeV2(request *KadasterUittrekselKadastraleKaartPostcodeV2RequestType) (*KadasterUittrekselKadastraleKaartPostcodeV2ResponseType, error)

	/* Returns a 'Kadata WMS De Kadastrale Kaart' map in the specified format. */
	KadasterKadastraleKaartPerceel(request *KadasterKadastraleKaartPerceelRequestType) (*KadasterKadastraleKaartPerceelResponseType, error)

	/* Returns a 'Kadata WMS De Kadastrale Kaart' map in the specified format. */
	KadasterKadastraleKaartPostcode(request *KadasterKadastraleKaartPostcodeRequestType) (*KadasterKadastraleKaartPostcodeResponseType, error)

	/* Find a 'Hypothecair bericht' by postcode and house number. If the input matches more than one parcel, a list of the parcels found is returned instead. */
	KadasterHypothecairBerichtPostcodeV2(request *KadasterHypothecairBerichtPostcodeV2RequestType) (*KadasterHypothecairBerichtPostcodeV2ResponseType, error)

	/* Find a 'Hypothecair bericht' by parcel details.  */
	KadasterHypothecairBerichtPerceelV2(request *KadasterHypothecairBerichtPerceelV2RequestType) (*KadasterHypothecairBerichtPerceelV2ResponseType, error)

	/* Returns a koopsommenoverzicht (in English: real estate transactions overview), which is a list of all transactions of the past five years in the given postcode range. */
	KadasterKoopsommenOverzicht(request *KadasterKoopsommenOverzichtRequestType) (*KadasterKoopsommenOverzichtResponseType, error)

	/* Fetch a "KadastraalBerichtObject" using an address */
	KadasterV2GetKadastraalBerichtObjectByAdres(request *KadasterV2GetKadastraalBerichtObjectByAdresRequestType) (*KadasterV2GetKadastraalBerichtObjectByAdresResponseType, error)

	/* Fetch a "KadastraalBerichtObject" using a designation for the parcel */
	KadasterV2GetKadastraalBerichtObjectByKadastraleAanduiding(request *KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingRequestType) (*KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingResponseType, error)

	/* Fetch a "KadastraalBerichtObject" using its identifier */
	KadasterV2GetKadastraalBerichtObjectByObjectId(request *KadasterV2GetKadastraalBerichtObjectByObjectIdRequestType) (*KadasterV2GetKadastraalBerichtObjectByObjectIdResponseType, error)

	/* Fetch a "HypothecairBerichtObject" using an address */
	KadasterV2GetHypothecairBerichtObjectByAdres(request *KadasterV2GetHypothecairBerichtObjectByAdresRequestType) (*KadasterV2GetHypothecairBerichtObjectByAdresResponseType, error)

	/* Fetch a "HypothecairBerichtObject" using a designation for the parcel */
	KadasterV2GetHypothecairBerichtObjectByKadastraleAanduiding(request *KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingRequestType) (*KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingResponseType, error)

	/* Fetch a "HypothecairBerichtObject" using its identifier */
	KadasterV2GetHypothecairBerichtObjectByObjectId(request *KadasterV2GetHypothecairBerichtObjectByObjectIdRequestType) (*KadasterV2GetHypothecairBerichtObjectByObjectIdResponseType, error)

	/* Fetch an "UittrekselKadastraleKaart" using an address */
	KadasterV2GetUittrekselKadastraleKaartByAdres(request *KadasterV2GetUittrekselKadastraleKaartByAdresRequestType) (*KadasterV2GetUittrekselKadastraleKaartByAdresResponseType, error)

	/* Fetch an "UittrekselKadastraleKaart" using a designation for the parcel */
	KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduiding(request *KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingRequestType) (*KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingResponseType, error)

	/* Fetch an "UittrekselKadastraleKaart" using its identifier */
	KadasterV2GetUittrekselKadastraleKaartByObjectId(request *KadasterV2GetUittrekselKadastraleKaartByObjectIdRequestType) (*KadasterV2GetUittrekselKadastraleKaartByObjectIdResponseType, error)

	/* Fetch a list of "Koopsommen" using an address */
	KadasterV2GetKoopsommenOverzicht(request *KadasterV2GetKoopsommenOverzichtRequestType) (*KadasterV2GetKoopsommenOverzichtResponseType, error)

	/* Fetch a "KadasterV2BronDocumentResponse" */
	KadasterV2GetBronDocument(request *KadasterV2GetBronDocumentRequestType) (*KadasterV2GetBronDocumentResponseType, error)

	/* Retrieve data on a business establishment */
	KvkGetDossier(request *KvkGetDossierRequestType) (*KvkGetDossierResponseType, error)

	/* Search for business establishments using a known identifier */
	KvkSearchDossierNumber(request *KvkSearchDossierNumberRequestType) (*KvkSearchDossierNumberResponseType, error)

	/* Find business establishments using a variety of parameters */
	KvkSearchParameters(request *KvkSearchParametersRequestType) (*KvkSearchParametersResponseType, error)

	/* Find business establishments based on postcode and house number */
	KvkSearchPostcode(request *KvkSearchPostcodeRequestType) (*KvkSearchPostcodeResponseType, error)

	/* Search for businesses matching all of the given criteria. */
	KvkSearchSelection(request *KvkSearchSelectionRequestType) (*KvkSearchSelectionResponseType, error)

	/* Retrieve a Chamber of Commerce extract document (Dutch: Uittreksel Handelsregister) */
	KvkGetExtractDocument(request *KvkGetExtractDocumentRequestType) (*KvkGetExtractDocumentResponseType, error)

	/* Retrieve information on the last change of a business establishment */
	KvkUpdateCheckDossier(request *KvkUpdateCheckDossierRequestType) (*KvkUpdateCheckDossierResponseType, error)

	/* Retrieve all dossiers changed since the given date. */
	KvkUpdateGetChangedDossiers(request *KvkUpdateGetChangedDossiersRequestType) (*KvkUpdateGetChangedDossiersResponseType, error)

	/* Returns a list of all dossiers that have been updated since they were last retrieved by the user */
	KvkUpdateGetDossiers(request *KvkUpdateGetDossiersRequestType) (*KvkUpdateGetDossiersResponseType, error)

	/* Add a dossier to the list of dossiers for which the user (the user whose credentials are used to make the call) wants to receive updates */
	KvkUpdateAddDossier(request *KvkUpdateAddDossierRequestType) (*KvkUpdateAddDossierResponseType, error)

	/* Remove a dossier from the list of dossiers for which the user (the user whose credentials are used to make the call) wants to receive updates. */
	KvkUpdateRemoveDossier(request *KvkUpdateRemoveDossierRequestType) (*KvkUpdateRemoveDossierResponseType, error)

	/* Returns a map in JPG or PNG format showing the location of the given postcode */
	MapViewPostcodeV2(request *MapViewPostcodeV2RequestType) (*MapViewPostcodeV2ResponseType, error)

	/* Returns a map showing the location of the given coordinate (centerlat, centerlon). extra_latlon is an array of other locations to be shown. format can be 'jpg' or 'png' (default). The width and height parameters can be [1 - 2048]. Scale in meters per pixel [0.298 - 156412] */
	MapViewLatLon(request *MapViewLatLonRequestType) (*MapViewLatLonResponseType, error)

	/* Returns a map showing the location of the given Rijksdriehoeksmeting coordinate (centerx, centery). extra_xy is an array of other locations to be shown. format can be 'jpg' or 'png' (default). The width and height parameters can be [1-2048]. Scale in meters per pixel [0.298 - 156412] */
	MapViewRD(request *MapViewRDRequestType) (*MapViewRDResponseType, error)

	/* Returns a map showing the location of the given coordinate (latitude, longitude) */
	MapViewInternationalLatLon(request *MapViewInternationalLatLonRequestType) (*MapViewInternationalLatLonResponseType, error)

	/* Fetch a list of all WebservicesNL endpoints (interfaces). Provide interface parameter [optional] for specific binding */
	MetaGetServices(request *MetaGetServicesRequestType) (*MetaGetServicesResponseType, error)

	/* Returns a list of methods for provided webservice (Endpoint) */
	MetaGetService(request *MetaGetServiceRequestType) (*MetaGetServiceResponseType, error)

	/* Returns a detailed overview of the requested method */
	MetaGetMethod(request *MetaGetMethodRequestType) (*MetaGetMethodResponseType, error)

	/* Returns a value estimate for the real estate at the specified address */
	NbwoEstimateValue(request *NbwoEstimateValueRequestType) (*NbwoEstimateValueResponseType, error)

	/* Determines the creditworthiness of a person, address and postcode area.
	   The creditworthiness is expressed as a number between 0 and 11.
	   A credit score of 0 indicates that there is no information available for the given search.
	   A score between 1 and 11 indicates the creditworthiness, where a higher score means less creditworthiness,
	   thus higher risk. */
	RiskCheckPerson(request *RiskCheckPersonRequestType) (*RiskCheckPersonResponseType, error)

	/* Retrieve information about a persons creditworthiness based on his/her business ownerships and insolvency registrations. */
	RiskCheckGetRiskPersonCompanyReport(request *RiskCheckGetRiskPersonCompanyReportRequestType) (*RiskCheckGetRiskPersonCompanyReportResponseType, error)

	/* Get description and image for a route between two addresses */
	RoutePlannerGetRoute(request *RoutePlannerGetRouteRequestType) (*RoutePlannerGetRouteResponseType, error)

	/* Summary information about a route calculated between two addresses */
	RoutePlannerInformationAddress(request *RoutePlannerInformationAddressRequestType) (*RoutePlannerInformationAddressResponseType, error)

	/* Detailed information about a route calculated between two addresses */
	RoutePlannerDescriptionAddress(request *RoutePlannerDescriptionAddressRequestType) (*RoutePlannerDescriptionAddressResponseType, error)

	/* Returns a description of the fastest route between two postcodes. For every part of the route the drivetime in seconds and drivedistance in meters are given as well. If the english argument is true the description will be in english, otherwise it will be in dutch.  */
	RoutePlannerDescription(request *RoutePlannerDescriptionRequestType) (*RoutePlannerDescriptionResponseType, error)

	/* Returns a description of the shortest route between two postcodes. For every part of the route the drivetime in seconds and drivedistance in meters are given as well. If the english argument is true the description will be in english, otherwise it will be in dutch.  */
	RoutePlannerDescriptionShortest(request *RoutePlannerDescriptionShortestRequestType) (*RoutePlannerDescriptionShortestResponseType, error)

	/* Returns a description of the route between two dutch postcodes, including the RD coordinates along the route. For every part of the route the drivetime in seconds and drivedistance in meters are given as well. The routetype can be shortest or fastest. By default the fastest route will be calculated. If the english argument is true the description will be in english, otherwise it will be in dutch.  */
	RoutePlannerDescriptionCoordinatesRD(request *RoutePlannerDescriptionCoordinatesRDRequestType) (*RoutePlannerDescriptionCoordinatesRDResponseType, error)

	/* Returns information (drivetime in seconds and drivedistance in meters) of the route between two postcodes. The routetype can be shortest or fastest. By default the fastest route will be calculated.  */
	RoutePlannerInformation(request *RoutePlannerInformationRequestType) (*RoutePlannerInformationResponseType, error)

	/* Returns a description of the route between two coordinates in the RD (Rijksdriehoeksmeting) coordinate system. For every part of the route the drivetime in seconds and drivedistance in meters are given as well. The description is available in dutch and english, depending on the english parameter toggle. The fastest or shortest route can be calculated depending on theroutetype parameter.  */
	RoutePlannerRDDescription(request *RoutePlannerRDDescriptionRequestType) (*RoutePlannerRDDescriptionResponseType, error)

	/* Returns the route distance and drive time between two coordinates in the RD (Rijksdriehoeksmeting) coordinate system. The fastest or shorted route can be calculated depending on the routetypeparameter.  */
	RoutePlannerRDInformation(request *RoutePlannerRDInformationRequestType) (*RoutePlannerRDInformationResponseType, error)

	/* Returns a description of the route between two coordinates in the RD (Rijksdriehoeksmeting) coordinate system, including the RD coordinates along the route. For every part of the route the drivetime in seconds and drivedistance in meters are given as well. The fastest or shortest route can be calculated depending on the routetype parameter. The description is available in dutch and english, depending on the english parameter toggle.  */
	RoutePlannerRDDescriptionCoordinatesRD(request *RoutePlannerRDDescriptionCoordinatesRDRequestType) (*RoutePlannerRDDescriptionCoordinatesRDResponseType, error)

	/* Summary information about a route calculated between two dutch-addresses */
	RoutePlannerInformationDutchAddress(request *RoutePlannerInformationDutchAddressRequestType) (*RoutePlannerInformationDutchAddressResponseType, error)

	/* Detailed information about a route calculated between two dutch-addresses */
	RoutePlannerDescriptionDutchAddress(request *RoutePlannerDescriptionDutchAddressRequestType) (*RoutePlannerDescriptionDutchAddressResponseType, error)

	/* Returns a description of the route between two latitude/longitude coordinates in Europe. For every part of the route the drivetime and drivedistance are given as well. The description is available in several languages depending on the language parameter. The fastest or shortest route is be calculated depending on the routetype parameter.  */
	RoutePlannerEUDescription(request *RoutePlannerEUDescriptionRequestType) (*RoutePlannerEUDescriptionResponseType, error)

	/* Returns the route distance and drive time for the route between two latitude/longitude coordinates in Europe. The fastest or shortest route can be calculated depending on the routetype parameter.  */
	RoutePlannerEUInformation(request *RoutePlannerEUInformationRequestType) (*RoutePlannerEUInformationResponseType, error)

	/* Returns a description of the route between two latitude/longitude coordinates in Europe, including the latitude/longitude coordinates along the route. For every part of the route the drivetime in seconds and drivedistance in meters are given as well. The routetype can be shortest or fastest. By default information on the fastest route will be returned. The description is available in several languages depending on the language parameter.  */
	RoutePlannerEUDescriptionCoordinatesLatLon(request *RoutePlannerEUDescriptionCoordinatesLatLonRequestType) (*RoutePlannerEUDescriptionCoordinatesLatLonResponseType, error)

	/* Returns a map showing the route between two latitude/longitude coordinates in Europe.  */
	RoutePlannerEUMap(request *RoutePlannerEUMapRequestType) (*RoutePlannerEUMapResponseType, error)

	/* Convert a bank account number, bank code and country code to an IBAN number */
	SepaConvertBankAccountNumber(request *SepaConvertBankAccountNumberRequestType) (*SepaConvertBankAccountNumberResponseType, error)

	/* Validate format of an International Bank Account Number (IBAN) */
	SepaValidateInternationalBankAccountNumberFormat(request *SepaValidateInternationalBankAccountNumberFormatRequestType) (*SepaValidateInternationalBankAccountNumberFormatResponseType, error)

	/* Validate and display values of IBAN values */
	SepaIbanDetails(request *SepaIbanDetailsRequestType) (*SepaIbanDetailsResponseType, error)

	/* Verify IBAN-name combinations */
	SepaMatchAccountHolder(request *SepaMatchAccountHolderRequestType) (*SepaMatchAccountHolderResponseType, error)

	/* Convert a local Basic Bank Account Number (BBAN) number to its International Bank Account Number (IBAN) counterpart. */
	SepaConvertBasicBankAccountNumber(request *SepaConvertBasicBankAccountNumberRequestType) (*SepaConvertBasicBankAccountNumberResponseType, error)

	/* Check if a VAT number is valid */
	VatValidate(request *VatValidateRequestType) (*VatValidateResponseType, error)

	/* Check if a VAT number is valid */
	VatViesProxyCheckVat(request *VatViesProxyCheckVatRequestType) (*VatViesProxyCheckVatResponseType, error)
}

type webservicesNLPortType struct {
	client *soap.Client
}

func NewWebservicesNLPortType(client *soap.Client) WebservicesNLPortType {
	return &webservicesNLPortType{
		client: client,
	}
}

func (service *webservicesNLPortType) AccountGetCreationToken(request *AccountGetCreationTokenRequestType) (*AccountGetCreationTokenResponseType, error) {
	response := new(AccountGetCreationTokenResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountGetCreationToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountGetCreationStatus(request *AccountGetCreationStatusRequestType) (*AccountGetCreationStatusResponseType, error) {
	response := new(AccountGetCreationStatusResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountGetCreationStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountGetOrderToken(request *AccountGetOrderTokenRequestType) (*AccountGetOrderTokenResponseType, error) {
	response := new(AccountGetOrderTokenResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountGetOrderToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserSessionRemove(request *UserSessionRemoveRequestType) (*UserSessionRemoveResponseType, error) {
	response := new(UserSessionRemoveResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userSessionRemove", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserSessionList(request *UserSessionListRequestType) (*UserSessionListResponseType, error) {
	response := new(UserSessionListResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userSessionList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserViewBalance(request *UserViewBalanceRequestType) (*UserViewBalanceResponseType, error) {
	response := new(UserViewBalanceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userViewBalance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserEditBalance(request *UserEditBalanceRequestType) (*UserEditBalanceResponseType, error) {
	response := new(UserEditBalanceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userEditBalance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountViewBalance(request *AccountViewBalanceRequestType) (*AccountViewBalanceResponseType, error) {
	response := new(AccountViewBalanceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountViewBalance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserViewV2(request *UserViewV2RequestType) (*UserViewV2ResponseType, error) {
	response := new(UserViewV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userViewV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserEditV2(request *UserEditV2RequestType) (*UserEditV2ResponseType, error) {
	response := new(UserEditV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userEditV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserEditExtendedV2(request *UserEditExtendedV2RequestType) (*UserEditExtendedV2ResponseType, error) {
	response := new(UserEditExtendedV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userEditExtendedV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserCreateV2(request *UserCreateV2RequestType) (*UserCreateV2ResponseType, error) {
	response := new(UserCreateV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userCreateV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreateTestUser(request *CreateTestUserRequestType) (*CreateTestUserResponseType, error) {
	response := new(CreateTestUserResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/createTestUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserChangePassword(request *UserChangePasswordRequestType) (*UserChangePasswordResponseType, error) {
	response := new(UserChangePasswordResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userChangePassword", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserRemove(request *UserRemoveRequestType) (*UserRemoveResponseType, error) {
	response := new(UserRemoveResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userRemove", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserNotify(request *UserNotifyRequestType) (*UserNotifyResponseType, error) {
	response := new(UserNotifyResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userNotify", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserListAssignableGroups(request *UserListAssignableGroupsRequestType) (*UserListAssignableGroupsResponseType, error) {
	response := new(UserListAssignableGroupsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userListAssignableGroups", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserAddGroup(request *UserAddGroupRequestType) (*UserAddGroupResponseType, error) {
	response := new(UserAddGroupResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userAddGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserRemoveGroup(request *UserRemoveGroupRequestType) (*UserRemoveGroupResponseType, error) {
	response := new(UserRemoveGroupResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userRemoveGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountViewV2(request *AccountViewV2RequestType) (*AccountViewV2ResponseType, error) {
	response := new(AccountViewV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountViewV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountEditV2(request *AccountEditV2RequestType) (*AccountEditV2ResponseType, error) {
	response := new(AccountEditV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountEditV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountUserListV2(request *AccountUserListV2RequestType) (*AccountUserListV2ResponseType, error) {
	response := new(AccountUserListV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountUserListV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountUserSearchV2(request *AccountUserSearchV2RequestType) (*AccountUserSearchV2ResponseType, error) {
	response := new(AccountUserSearchV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountUserSearchV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountEditHostRestrictions(request *AccountEditHostRestrictionsRequestType) (*AccountEditHostRestrictionsResponseType, error) {
	response := new(AccountEditHostRestrictionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountEditHostRestrictions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AccountViewHostRestrictions(request *AccountViewHostRestrictionsRequestType) (*AccountViewHostRestrictionsResponseType, error) {
	response := new(AccountViewHostRestrictionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/accountViewHostRestrictions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserEditHostRestrictions(request *UserEditHostRestrictionsRequestType) (*UserEditHostRestrictionsResponseType, error) {
	response := new(UserEditHostRestrictionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userEditHostRestrictions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) UserViewHostRestrictions(request *UserViewHostRestrictionsRequestType) (*UserViewHostRestrictionsResponseType, error) {
	response := new(UserViewHostRestrictionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/userViewHostRestrictions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressReeksPostcodeSearch(request *AddressReeksPostcodeSearchRequestType) (*AddressReeksPostcodeSearchResponseType, error) {
	response := new(AddressReeksPostcodeSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressReeksPostcodeSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressReeksAddressSearch(request *AddressReeksAddressSearchRequestType) (*AddressReeksAddressSearchResponseType, error) {
	response := new(AddressReeksAddressSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressReeksAddressSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressReeksFullParameterSearch(request *AddressReeksFullParameterSearchRequestType) (*AddressReeksFullParameterSearchResponseType, error) {
	response := new(AddressReeksFullParameterSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressReeksFullParameterSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressReeksParameterSearch(request *AddressReeksParameterSearchRequestType) (*AddressReeksParameterSearchResponseType, error) {
	response := new(AddressReeksParameterSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressReeksParameterSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressPerceelPhraseSearch(request *AddressPerceelPhraseSearchRequestType) (*AddressPerceelPhraseSearchResponseType, error) {
	response := new(AddressPerceelPhraseSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressPerceelPhraseSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressPerceelFullParameterSearchV2(request *AddressPerceelFullParameterSearchV2RequestType) (*AddressPerceelFullParameterSearchV2ResponseType, error) {
	response := new(AddressPerceelFullParameterSearchV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressPerceelFullParameterSearchV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressProvinceListNeighborhoods(request *AddressProvinceListNeighborhoodsRequestType) (*AddressProvinceListNeighborhoodsResponseType, error) {
	response := new(AddressProvinceListNeighborhoodsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressProvinceListNeighborhoods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressProvinceListDistricts(request *AddressProvinceListDistrictsRequestType) (*AddressProvinceListDistrictsResponseType, error) {
	response := new(AddressProvinceListDistrictsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressProvinceListDistricts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressProvinceList(request *AddressProvinceListRequestType) (*AddressProvinceListResponseType, error) {
	response := new(AddressProvinceListResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressProvinceList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressProvinceSearch(request *AddressProvinceSearchRequestType) (*AddressProvinceSearchResponseType, error) {
	response := new(AddressProvinceSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressProvinceSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressDistrictSearch(request *AddressDistrictSearchRequestType) (*AddressDistrictSearchResponseType, error) {
	response := new(AddressDistrictSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressDistrictSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressDistrictListCities(request *AddressDistrictListCitiesRequestType) (*AddressDistrictListCitiesResponseType, error) {
	response := new(AddressDistrictListCitiesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressDistrictListCities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressDistrictListNeighborhoods(request *AddressDistrictListNeighborhoodsRequestType) (*AddressDistrictListNeighborhoodsResponseType, error) {
	response := new(AddressDistrictListNeighborhoodsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressDistrictListNeighborhoods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressCitySearchV2(request *AddressCitySearchV2RequestType) (*AddressCitySearchV2ResponseType, error) {
	response := new(AddressCitySearchV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressCitySearchV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressCityListNeighborhoods(request *AddressCityListNeighborhoodsRequestType) (*AddressCityListNeighborhoodsResponseType, error) {
	response := new(AddressCityListNeighborhoodsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressCityListNeighborhoods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressPerceelFullParameterSearch(request *AddressPerceelFullParameterSearchRequestType) (*AddressPerceelFullParameterSearchResponseType, error) {
	response := new(AddressPerceelFullParameterSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressPerceelFullParameterSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressPerceelParameterSearch(request *AddressPerceelParameterSearchRequestType) (*AddressPerceelParameterSearchResponseType, error) {
	response := new(AddressPerceelParameterSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressPerceelParameterSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressReeksPhraseSearch(request *AddressReeksPhraseSearchRequestType) (*AddressReeksPhraseSearchResponseType, error) {
	response := new(AddressReeksPhraseSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressReeksPhraseSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressNeighborhoodName(request *AddressNeighborhoodNameRequestType) (*AddressNeighborhoodNameResponseType, error) {
	response := new(AddressNeighborhoodNameResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressNeighborhoodName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AddressNeighborhoodPhraseSearch(request *AddressNeighborhoodPhraseSearchRequestType) (*AddressNeighborhoodPhraseSearchResponseType, error) {
	response := new(AddressNeighborhoodPhraseSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/addressNeighborhoodPhraseSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AreaCodeLookup(request *AreaCodeLookupRequestType) (*AreaCodeLookupResponseType, error) {
	response := new(AreaCodeLookupResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/areaCodeLookup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AreaCodeToNeighborhoodcode(request *AreaCodeToNeighborhoodcodeRequestType) (*AreaCodeToNeighborhoodcodeResponseType, error) {
	response := new(AreaCodeToNeighborhoodcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/areaCodeToNeighborhoodcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) AreaCodePostcodeLookup(request *AreaCodePostcodeLookupRequestType) (*AreaCodePostcodeLookupResponseType, error) {
	response := new(AreaCodePostcodeLookupResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/areaCodePostcodeLookup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) Login(request *LoginRequestType) (*LoginResponseType, error) {
	response := new(LoginResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/login", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) Logout(request *LogoutRequestType) (*LogoutResponseType, error) {
	response := new(LogoutResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/logout", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BelgianBusinessGetDossier(request *BelgianBusinessGetDossierRequestType) (*BelgianBusinessGetDossierResponseType, error) {
	response := new(BelgianBusinessGetDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/belgianBusinessGetDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BovagGetMemberByBovagId(request *BovagGetMemberByBovagIdRequestType) (*BovagGetMemberByBovagIdResponseType, error) {
	response := new(BovagGetMemberByBovagIdResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/bovagGetMemberByBovagId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BovagGetMemberByDutchBusiness(request *BovagGetMemberByDutchBusinessRequestType) (*BovagGetMemberByDutchBusinessResponseType, error) {
	response := new(BovagGetMemberByDutchBusinessResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/bovagGetMemberByDutchBusiness", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessGetEstablishmentNumber(request *BusinessGetEstablishmentNumberRequestType) (*BusinessGetEstablishmentNumberResponseType, error) {
	response := new(BusinessGetEstablishmentNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessGetEstablishmentNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessGetBIKDescription(request *BusinessGetBIKDescriptionRequestType) (*BusinessGetBIKDescriptionResponseType, error) {
	response := new(BusinessGetBIKDescriptionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessGetBIKDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessGetSBIDescription(request *BusinessGetSBIDescriptionRequestType) (*BusinessGetSBIDescriptionResponseType, error) {
	response := new(BusinessGetSBIDescriptionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessGetSBIDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessBIKToSBI(request *BusinessBIKToSBIRequestType) (*BusinessBIKToSBIResponseType, error) {
	response := new(BusinessBIKToSBIResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessBIKToSBI", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSBIToBIK(request *BusinessSBIToBIKRequestType) (*BusinessSBIToBIKResponseType, error) {
	response := new(BusinessSBIToBIKResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSBIToBIK", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessGetDossierV3(request *BusinessGetDossierV3RequestType) (*BusinessGetDossierV3ResponseType, error) {
	response := new(BusinessGetDossierV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessGetDossierV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessGetDossierExtended(request *BusinessGetDossierExtendedRequestType) (*BusinessGetDossierExtendedResponseType, error) {
	response := new(BusinessGetDossierExtendedResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessGetDossierExtended", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchDossierNumber(request *BusinessSearchDossierNumberRequestType) (*BusinessSearchDossierNumberResponseType, error) {
	response := new(BusinessSearchDossierNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchDossierNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchPostcode(request *BusinessSearchPostcodeRequestType) (*BusinessSearchPostcodeResponseType, error) {
	response := new(BusinessSearchPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchAddress(request *BusinessSearchAddressRequestType) (*BusinessSearchAddressResponseType, error) {
	response := new(BusinessSearchAddressResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchName(request *BusinessSearchNameRequestType) (*BusinessSearchNameResponseType, error) {
	response := new(BusinessSearchNameResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchParameters(request *BusinessSearchParametersRequestType) (*BusinessSearchParametersResponseType, error) {
	response := new(BusinessSearchParametersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchParameters", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchParametersV3(request *BusinessSearchParametersV3RequestType) (*BusinessSearchParametersV3ResponseType, error) {
	response := new(BusinessSearchParametersV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchParametersV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchSelection(request *BusinessSearchSelectionRequestType) (*BusinessSearchSelectionResponseType, error) {
	response := new(BusinessSearchSelectionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchSelection", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchSelectionV2(request *BusinessSearchSelectionV2RequestType) (*BusinessSearchSelectionV2ResponseType, error) {
	response := new(BusinessSearchSelectionV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchSelectionV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessGetDossierSBI(request *BusinessGetDossierSBIRequestType) (*BusinessGetDossierSBIResponseType, error) {
	response := new(BusinessGetDossierSBIResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessGetDossierSBI", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessUpdateCheckDossier(request *BusinessUpdateCheckDossierRequestType) (*BusinessUpdateCheckDossierResponseType, error) {
	response := new(BusinessUpdateCheckDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessUpdateCheckDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessUpdateGetChangedDossiers(request *BusinessUpdateGetChangedDossiersRequestType) (*BusinessUpdateGetChangedDossiersResponseType, error) {
	response := new(BusinessUpdateGetChangedDossiersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessUpdateGetChangedDossiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessUpdateGetDossiers(request *BusinessUpdateGetDossiersRequestType) (*BusinessUpdateGetDossiersResponseType, error) {
	response := new(BusinessUpdateGetDossiersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessUpdateGetDossiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessUpdateAddDossier(request *BusinessUpdateAddDossierRequestType) (*BusinessUpdateAddDossierResponseType, error) {
	response := new(BusinessUpdateAddDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessUpdateAddDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessUpdateRemoveDossier(request *BusinessUpdateRemoveDossierRequestType) (*BusinessUpdateRemoveDossierResponseType, error) {
	response := new(BusinessUpdateRemoveDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessUpdateRemoveDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) BusinessSearchParametersV2(request *BusinessSearchParametersV2RequestType) (*BusinessSearchParametersV2ResponseType, error) {
	response := new(BusinessSearchParametersV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/businessSearchParametersV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarCheckCode(request *CarRDWCarCheckCodeRequestType) (*CarRDWCarCheckCodeResponseType, error) {
	response := new(CarRDWCarCheckCodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarCheckCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataV3(request *CarRDWCarDataV3RequestType) (*CarRDWCarDataV3ResponseType, error) {
	response := new(CarRDWCarDataV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataBPV2(request *CarRDWCarDataBPV2RequestType) (*CarRDWCarDataBPV2ResponseType, error) {
	response := new(CarRDWCarDataBPV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataBPV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataExtended(request *CarRDWCarDataExtendedRequestType) (*CarRDWCarDataExtendedResponseType, error) {
	response := new(CarRDWCarDataExtendedResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataExtended", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataPrice(request *CarRDWCarDataPriceRequestType) (*CarRDWCarDataPriceResponseType, error) {
	response := new(CarRDWCarDataPriceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataPrice", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataOptions(request *CarRDWCarDataOptionsRequestType) (*CarRDWCarDataOptionsResponseType, error) {
	response := new(CarRDWCarDataOptionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEMeldcodeCheck(request *CarVWEMeldcodeCheckRequestType) (*CarVWEMeldcodeCheckResponseType, error) {
	response := new(CarVWEMeldcodeCheckResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEMeldcodeCheck", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEBasicTypeData(request *CarVWEBasicTypeDataRequestType) (*CarVWEBasicTypeDataResponseType, error) {
	response := new(CarVWEBasicTypeDataResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEBasicTypeData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEVersionPrice(request *CarVWEVersionPriceRequestType) (*CarVWEVersionPriceResponseType, error) {
	response := new(CarVWEVersionPriceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEVersionPrice", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEOptions(request *CarVWEOptionsRequestType) (*CarVWEOptionsResponseType, error) {
	response := new(CarVWEOptionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEListBrands(request *CarVWEListBrandsRequestType) (*CarVWEListBrandsResponseType, error) {
	response := new(CarVWEListBrandsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEListBrands", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEListModels(request *CarVWEListModelsRequestType) (*CarVWEListModelsResponseType, error) {
	response := new(CarVWEListModelsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEListModels", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEListVersions(request *CarVWEListVersionsRequestType) (*CarVWEListVersionsResponseType, error) {
	response := new(CarVWEListVersionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEListVersions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEVersionYearData(request *CarVWEVersionYearDataRequestType) (*CarVWEVersionYearDataResponseType, error) {
	response := new(CarVWEVersionYearDataResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEVersionYearData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarVWEPhotos(request *CarVWEPhotosRequestType) (*CarVWEPhotosResponseType, error) {
	response := new(CarVWEPhotosResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carVWEPhotos", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarATDPrice(request *CarATDPriceRequestType) (*CarATDPriceResponseType, error) {
	response := new(CarATDPriceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carATDPrice", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarData(request *CarRDWCarDataRequestType) (*CarRDWCarDataResponseType, error) {
	response := new(CarRDWCarDataResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataV2(request *CarRDWCarDataV2RequestType) (*CarRDWCarDataV2ResponseType, error) {
	response := new(CarRDWCarDataV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CarRDWCarDataBP(request *CarRDWCarDataBPRequestType) (*CarRDWCarDataBPResponseType, error) {
	response := new(CarRDWCarDataBPResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/carRDWCarDataBP", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) ComplianceSearchPersons(request *ComplianceSearchPersonsRequestType) (*ComplianceSearchPersonsResponseType, error) {
	response := new(ComplianceSearchPersonsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/complianceSearchPersons", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) ComplianceGetPerson(request *ComplianceGetPersonRequestType) (*ComplianceGetPersonResponseType, error) {
	response := new(ComplianceGetPersonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/complianceGetPerson", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) ComplianceSearchBusinesses(request *ComplianceSearchBusinessesRequestType) (*ComplianceSearchBusinessesResponseType, error) {
	response := new(ComplianceSearchBusinessesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/complianceSearchBusinesses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) ComplianceGetBusiness(request *ComplianceGetBusinessRequestType) (*ComplianceGetBusinessResponseType, error) {
	response := new(ComplianceGetBusinessResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/complianceGetBusiness", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeGetReportFull(request *CreditsafeGetReportFullRequestType) (*CreditsafeGetReportFullResponseType, error) {
	response := new(CreditsafeGetReportFullResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeGetReportFull", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeSearch(request *CreditsafeSearchRequestType) (*CreditsafeSearchResponseType, error) {
	response := new(CreditsafeSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeSearchV2(request *CreditsafeSearchV2RequestType) (*CreditsafeSearchV2ResponseType, error) {
	response := new(CreditsafeSearchV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeSearchV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeGetReportFullV2(request *CreditsafeGetReportFullV2RequestType) (*CreditsafeGetReportFullV2ResponseType, error) {
	response := new(CreditsafeGetReportFullV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeGetReportFullV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeSearchCriteria(request *CreditsafeSearchCriteriaRequestType) (*CreditsafeSearchCriteriaResponseType, error) {
	response := new(CreditsafeSearchCriteriaResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeSearchCriteria", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeMonitorAddCompany(request *CreditsafeMonitorAddCompanyRequestType) (*CreditsafeMonitorAddCompanyResponseType, error) {
	response := new(CreditsafeMonitorAddCompanyResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeMonitorAddCompany", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeMonitorRemoveCompany(request *CreditsafeMonitorRemoveCompanyRequestType) (*CreditsafeMonitorRemoveCompanyResponseType, error) {
	response := new(CreditsafeMonitorRemoveCompanyResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeMonitorRemoveCompany", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) CreditsafeMonitorGetUpdatedCompanies(request *CreditsafeMonitorGetUpdatedCompaniesRequestType) (*CreditsafeMonitorGetUpdatedCompaniesResponseType, error) {
	response := new(CreditsafeMonitorGetUpdatedCompaniesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/creditsafeMonitorGetUpdatedCompanies", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbSearchReferenceV2(request *DnbSearchReferenceV2RequestType) (*DnbSearchReferenceV2ResponseType, error) {
	response := new(DnbSearchReferenceV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbSearchReferenceV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbGetReference(request *DnbGetReferenceRequestType) (*DnbGetReferenceResponseType, error) {
	response := new(DnbGetReferenceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbGetReference", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbWorldbaseMarketing(request *DnbWorldbaseMarketingRequestType) (*DnbWorldbaseMarketingResponseType, error) {
	response := new(DnbWorldbaseMarketingResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbWorldbaseMarketing", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbWorldbaseMarketingPlus(request *DnbWorldbaseMarketingPlusRequestType) (*DnbWorldbaseMarketingPlusResponseType, error) {
	response := new(DnbWorldbaseMarketingPlusResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbWorldbaseMarketingPlus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbWorldbaseMarketingPlusLinkage(request *DnbWorldbaseMarketingPlusLinkageRequestType) (*DnbWorldbaseMarketingPlusLinkageResponseType, error) {
	response := new(DnbWorldbaseMarketingPlusLinkageResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbWorldbaseMarketingPlusLinkage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbQuickCheck(request *DnbQuickCheckRequestType) (*DnbQuickCheckResponseType, error) {
	response := new(DnbQuickCheckResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbQuickCheck", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbBusinessVerification(request *DnbBusinessVerificationRequestType) (*DnbBusinessVerificationResponseType, error) {
	response := new(DnbBusinessVerificationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbBusinessVerification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbEnterpriseManagement(request *DnbEnterpriseManagementRequestType) (*DnbEnterpriseManagementResponseType, error) {
	response := new(DnbEnterpriseManagementResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbEnterpriseManagement", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DnbSearchReference(request *DnbSearchReferenceRequestType) (*DnbSearchReferenceResponseType, error) {
	response := new(DnbSearchReferenceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dnbSearchReference", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DriveInfoDistanceLookup(request *DriveInfoDistanceLookupRequestType) (*DriveInfoDistanceLookupResponseType, error) {
	response := new(DriveInfoDistanceLookupResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/driveInfoDistanceLookup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DriveInfoTimeLookup(request *DriveInfoTimeLookupRequestType) (*DriveInfoTimeLookupResponseType, error) {
	response := new(DriveInfoTimeLookupResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/driveInfoTimeLookup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchAddressRangePostcodeSearch(request *DutchAddressRangePostcodeSearchRequestType) (*DutchAddressRangePostcodeSearchResponseType, error) {
	response := new(DutchAddressRangePostcodeSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchAddressRangePostcodeSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetDossier(request *DutchBusinessGetDossierRequestType) (*DutchBusinessGetDossierResponseType, error) {
	response := new(DutchBusinessGetDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetDossierV2(request *DutchBusinessGetDossierV2RequestType) (*DutchBusinessGetDossierV2ResponseType, error) {
	response := new(DutchBusinessGetDossierV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetDossierV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetDossierV3(request *DutchBusinessGetDossierV3RequestType) (*DutchBusinessGetDossierV3ResponseType, error) {
	response := new(DutchBusinessGetDossierV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetDossierV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetSBI(request *DutchBusinessGetSBIRequestType) (*DutchBusinessGetSBIResponseType, error) {
	response := new(DutchBusinessGetSBIResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetSBI", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetVatNumber(request *DutchBusinessGetVatNumberRequestType) (*DutchBusinessGetVatNumberResponseType, error) {
	response := new(DutchBusinessGetVatNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetVatNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetPositions(request *DutchBusinessGetPositionsRequestType) (*DutchBusinessGetPositionsResponseType, error) {
	response := new(DutchBusinessGetPositionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetPositions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetLegalEntity(request *DutchBusinessGetLegalEntityRequestType) (*DutchBusinessGetLegalEntityResponseType, error) {
	response := new(DutchBusinessGetLegalEntityResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetLegalEntity", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetOrganizationTree(request *DutchBusinessGetOrganizationTreeRequestType) (*DutchBusinessGetOrganizationTreeResponseType, error) {
	response := new(DutchBusinessGetOrganizationTreeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetOrganizationTree", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetNonMarketingIndicator(request *DutchBusinessGetNonMarketingIndicatorRequestType) (*DutchBusinessGetNonMarketingIndicatorResponseType, error) {
	response := new(DutchBusinessGetNonMarketingIndicatorResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetNonMarketingIndicator", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchDossierNumber(request *DutchBusinessSearchDossierNumberRequestType) (*DutchBusinessSearchDossierNumberResponseType, error) {
	response := new(DutchBusinessSearchDossierNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchDossierNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetConcernRelationsOverview(request *DutchBusinessGetConcernRelationsOverviewRequestType) (*DutchBusinessGetConcernRelationsOverviewResponseType, error) {
	response := new(DutchBusinessGetConcernRelationsOverviewResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetConcernRelationsOverview", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetConcernRelationsDetails(request *DutchBusinessGetConcernRelationsDetailsRequestType) (*DutchBusinessGetConcernRelationsDetailsResponseType, error) {
	response := new(DutchBusinessGetConcernRelationsDetailsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetConcernRelationsDetails", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchParametersV2(request *DutchBusinessSearchParametersV2RequestType) (*DutchBusinessSearchParametersV2ResponseType, error) {
	response := new(DutchBusinessSearchParametersV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchParametersV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearch(request *DutchBusinessSearchRequestType) (*DutchBusinessSearchResponseType, error) {
	response := new(DutchBusinessSearchResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchEstablishments(request *DutchBusinessSearchEstablishmentsRequestType) (*DutchBusinessSearchEstablishmentsResponseType, error) {
	response := new(DutchBusinessSearchEstablishmentsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchEstablishments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchPostcode(request *DutchBusinessSearchPostcodeRequestType) (*DutchBusinessSearchPostcodeResponseType, error) {
	response := new(DutchBusinessSearchPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchSelection(request *DutchBusinessSearchSelectionRequestType) (*DutchBusinessSearchSelectionResponseType, error) {
	response := new(DutchBusinessSearchSelectionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchSelection", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchSelectionSimple(request *DutchBusinessSearchSelectionSimpleRequestType) (*DutchBusinessSearchSelectionSimpleResponseType, error) {
	response := new(DutchBusinessSearchSelectionSimpleResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchSelectionSimple", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetSBIDescription(request *DutchBusinessGetSBIDescriptionRequestType) (*DutchBusinessGetSBIDescriptionResponseType, error) {
	response := new(DutchBusinessGetSBIDescriptionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetSBIDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetAnnualFinancialStatement(request *DutchBusinessGetAnnualFinancialStatementRequestType) (*DutchBusinessGetAnnualFinancialStatementResponseType, error) {
	response := new(DutchBusinessGetAnnualFinancialStatementResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetAnnualFinancialStatement", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetAdditionalPositions(request *DutchBusinessGetAdditionalPositionsRequestType) (*DutchBusinessGetAdditionalPositionsResponseType, error) {
	response := new(DutchBusinessGetAdditionalPositionsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetAdditionalPositions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessLookALikes(request *DutchBusinessLookALikesRequestType) (*DutchBusinessLookALikesResponseType, error) {
	response := new(DutchBusinessLookALikesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessLookALikes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractDocument(request *DutchBusinessGetExtractDocumentRequestType) (*DutchBusinessGetExtractDocumentResponseType, error) {
	response := new(DutchBusinessGetExtractDocumentResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractDocument", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractDocumentData(request *DutchBusinessGetExtractDocumentDataRequestType) (*DutchBusinessGetExtractDocumentDataResponseType, error) {
	response := new(DutchBusinessGetExtractDocumentDataResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractDocumentData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractDocumentDataV2(request *DutchBusinessGetExtractDocumentDataV2RequestType) (*DutchBusinessGetExtractDocumentDataV2ResponseType, error) {
	response := new(DutchBusinessGetExtractDocumentDataV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractDocumentDataV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractDocumentDataV3(request *DutchBusinessGetExtractDocumentDataV3RequestType) (*DutchBusinessGetExtractDocumentDataV3ResponseType, error) {
	response := new(DutchBusinessGetExtractDocumentDataV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractDocumentDataV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetLegalExtractDocumentDataV2(request *DutchBusinessGetLegalExtractDocumentDataV2RequestType) (*DutchBusinessGetLegalExtractDocumentDataV2ResponseType, error) {
	response := new(DutchBusinessGetLegalExtractDocumentDataV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetLegalExtractDocumentDataV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetLegalExtractDocumentDataV3(request *DutchBusinessGetLegalExtractDocumentDataV3RequestType) (*DutchBusinessGetLegalExtractDocumentDataV3ResponseType, error) {
	response := new(DutchBusinessGetLegalExtractDocumentDataV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetLegalExtractDocumentDataV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractHistory(request *DutchBusinessGetExtractHistoryRequestType) (*DutchBusinessGetExtractHistoryResponseType, error) {
	response := new(DutchBusinessGetExtractHistoryResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractHistory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractHistoryChanged(request *DutchBusinessGetExtractHistoryChangedRequestType) (*DutchBusinessGetExtractHistoryChangedResponseType, error) {
	response := new(DutchBusinessGetExtractHistoryChangedResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractHistoryChanged", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractHistoryDocumentData(request *DutchBusinessGetExtractHistoryDocumentDataRequestType) (*DutchBusinessGetExtractHistoryDocumentDataResponseType, error) {
	response := new(DutchBusinessGetExtractHistoryDocumentDataResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractHistoryDocumentData", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractHistoryDocumentDataV2(request *DutchBusinessGetExtractHistoryDocumentDataV2RequestType) (*DutchBusinessGetExtractHistoryDocumentDataV2ResponseType, error) {
	response := new(DutchBusinessGetExtractHistoryDocumentDataV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractHistoryDocumentDataV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractHistoryDocumentDataV3(request *DutchBusinessGetExtractHistoryDocumentDataV3RequestType) (*DutchBusinessGetExtractHistoryDocumentDataV3ResponseType, error) {
	response := new(DutchBusinessGetExtractHistoryDocumentDataV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractHistoryDocumentDataV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetExtractHistoryDocumentDataV3ByDossier(request *DutchBusinessGetExtractHistoryDocumentDataV3ByDossierRequestType) (*DutchBusinessGetExtractHistoryDocumentDataV3ByDossierResponseType, error) {
	response := new(DutchBusinessGetExtractHistoryDocumentDataV3ByDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetExtractHistoryDocumentDataV3ByDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetDossierHistory(request *DutchBusinessGetDossierHistoryRequestType) (*DutchBusinessGetDossierHistoryResponseType, error) {
	response := new(DutchBusinessGetDossierHistoryResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetDossierHistory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUpdateGetDossiers(request *DutchBusinessUpdateGetDossiersRequestType) (*DutchBusinessUpdateGetDossiersResponseType, error) {
	response := new(DutchBusinessUpdateGetDossiersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUpdateGetDossiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUpdateCheckDossier(request *DutchBusinessUpdateCheckDossierRequestType) (*DutchBusinessUpdateCheckDossierResponseType, error) {
	response := new(DutchBusinessUpdateCheckDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUpdateCheckDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUpdateGetChangedDossiers(request *DutchBusinessUpdateGetChangedDossiersRequestType) (*DutchBusinessUpdateGetChangedDossiersResponseType, error) {
	response := new(DutchBusinessUpdateGetChangedDossiersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUpdateGetChangedDossiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUpdateAddDossier(request *DutchBusinessUpdateAddDossierRequestType) (*DutchBusinessUpdateAddDossierResponseType, error) {
	response := new(DutchBusinessUpdateAddDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUpdateAddDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUpdateRemoveDossier(request *DutchBusinessUpdateRemoveDossierRequestType) (*DutchBusinessUpdateRemoveDossierResponseType, error) {
	response := new(DutchBusinessUpdateRemoveDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUpdateRemoveDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUpdateGetOverview(request *DutchBusinessUpdateGetOverviewRequestType) (*DutchBusinessUpdateGetOverviewResponseType, error) {
	response := new(DutchBusinessUpdateGetOverviewResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUpdateGetOverview", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchNewsByDossier(request *DutchBusinessSearchNewsByDossierRequestType) (*DutchBusinessSearchNewsByDossierResponseType, error) {
	response := new(DutchBusinessSearchNewsByDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchNewsByDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUBOStartInvestigation(request *DutchBusinessUBOStartInvestigationRequestType) (*DutchBusinessUBOStartInvestigationResponseType, error) {
	response := new(DutchBusinessUBOStartInvestigationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUBOStartInvestigation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUBOCheckInvestigation(request *DutchBusinessUBOCheckInvestigationRequestType) (*DutchBusinessUBOCheckInvestigationResponseType, error) {
	response := new(DutchBusinessUBOCheckInvestigationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUBOCheckInvestigation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUBOPickupInvestigation(request *DutchBusinessUBOPickupInvestigationRequestType) (*DutchBusinessUBOPickupInvestigationResponseType, error) {
	response := new(DutchBusinessUBOPickupInvestigationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUBOPickupInvestigation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUBOCostsInvestigation(request *DutchBusinessUBOCostsInvestigationRequestType) (*DutchBusinessUBOCostsInvestigationResponseType, error) {
	response := new(DutchBusinessUBOCostsInvestigationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUBOCostsInvestigation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessUBOClassifyInvestigation(request *DutchBusinessUBOClassifyInvestigationRequestType) (*DutchBusinessUBOClassifyInvestigationResponseType, error) {
	response := new(DutchBusinessUBOClassifyInvestigationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessUBOClassifyInvestigation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessGetLei(request *DutchBusinessGetLeiRequestType) (*DutchBusinessGetLeiResponseType, error) {
	response := new(DutchBusinessGetLeiResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessGetLei", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchBusinessSearchParameters(request *DutchBusinessSearchParametersRequestType) (*DutchBusinessSearchParametersResponseType, error) {
	response := new(DutchBusinessSearchParametersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchBusinessSearchParameters", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchVehicleGetVehicleV2(request *DutchVehicleGetVehicleV2RequestType) (*DutchVehicleGetVehicleV2ResponseType, error) {
	response := new(DutchVehicleGetVehicleV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchVehicleGetVehicleV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchVehicleGetPurchaseReference(request *DutchVehicleGetPurchaseReferenceRequestType) (*DutchVehicleGetPurchaseReferenceResponseType, error) {
	response := new(DutchVehicleGetPurchaseReferenceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchVehicleGetPurchaseReference", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchVehicleGetOwnerHistory(request *DutchVehicleGetOwnerHistoryRequestType) (*DutchVehicleGetOwnerHistoryResponseType, error) {
	response := new(DutchVehicleGetOwnerHistoryResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchVehicleGetOwnerHistory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchVehicleGetMarketValue(request *DutchVehicleGetMarketValueRequestType) (*DutchVehicleGetMarketValueResponseType, error) {
	response := new(DutchVehicleGetMarketValueResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchVehicleGetMarketValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) DutchVehicleGetVehicle(request *DutchVehicleGetVehicleRequestType) (*DutchVehicleGetVehicleResponseType, error) {
	response := new(DutchVehicleGetVehicleResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/dutchVehicleGetVehicle", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) EdrGetScore(request *EdrGetScoreRequestType) (*EdrGetScoreResponseType, error) {
	response := new(EdrGetScoreResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/edrGetScore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationNeighborhoodCoordinatesRD(request *GeoLocationNeighborhoodCoordinatesRDRequestType) (*GeoLocationNeighborhoodCoordinatesRDResponseType, error) {
	response := new(GeoLocationNeighborhoodCoordinatesRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationNeighborhoodCoordinatesRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationPostcodeCoordinatesRD(request *GeoLocationPostcodeCoordinatesRDRequestType) (*GeoLocationPostcodeCoordinatesRDResponseType, error) {
	response := new(GeoLocationPostcodeCoordinatesRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationPostcodeCoordinatesRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationNeighborhoodCoordinatesLatLon(request *GeoLocationNeighborhoodCoordinatesLatLonRequestType) (*GeoLocationNeighborhoodCoordinatesLatLonResponseType, error) {
	response := new(GeoLocationNeighborhoodCoordinatesLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationNeighborhoodCoordinatesLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationPostcodeCoordinatesLatLon(request *GeoLocationPostcodeCoordinatesLatLonRequestType) (*GeoLocationPostcodeCoordinatesLatLonResponseType, error) {
	response := new(GeoLocationPostcodeCoordinatesLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationPostcodeCoordinatesLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationAddressCoordinatesLatLon(request *GeoLocationAddressCoordinatesLatLonRequestType) (*GeoLocationAddressCoordinatesLatLonResponseType, error) {
	response := new(GeoLocationAddressCoordinatesLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationAddressCoordinatesLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationAddressCoordinatesRD(request *GeoLocationAddressCoordinatesRDRequestType) (*GeoLocationAddressCoordinatesRDResponseType, error) {
	response := new(GeoLocationAddressCoordinatesRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationAddressCoordinatesRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationLatLonToPostcode(request *GeoLocationLatLonToPostcodeRequestType) (*GeoLocationLatLonToPostcodeResponseType, error) {
	response := new(GeoLocationLatLonToPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationLatLonToPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationLatLonToAddressV2(request *GeoLocationLatLonToAddressV2RequestType) (*GeoLocationLatLonToAddressV2ResponseType, error) {
	response := new(GeoLocationLatLonToAddressV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationLatLonToAddressV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationRDToPostcode(request *GeoLocationRDToPostcodeRequestType) (*GeoLocationRDToPostcodeResponseType, error) {
	response := new(GeoLocationRDToPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationRDToPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationRDToAddressV2(request *GeoLocationRDToAddressV2RequestType) (*GeoLocationRDToAddressV2ResponseType, error) {
	response := new(GeoLocationRDToAddressV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationRDToAddressV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationInternationalPostcodeCoordinatesLatLon(request *GeoLocationInternationalPostcodeCoordinatesLatLonRequestType) (*GeoLocationInternationalPostcodeCoordinatesLatLonResponseType, error) {
	response := new(GeoLocationInternationalPostcodeCoordinatesLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationInternationalPostcodeCoordinatesLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationInternationalAddressCoordinatesLatLon(request *GeoLocationInternationalAddressCoordinatesLatLonRequestType) (*GeoLocationInternationalAddressCoordinatesLatLonResponseType, error) {
	response := new(GeoLocationInternationalAddressCoordinatesLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationInternationalAddressCoordinatesLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationInternationalAddressCoordinatesLatLonV2(request *GeoLocationInternationalAddressCoordinatesLatLonV2RequestType) (*GeoLocationInternationalAddressCoordinatesLatLonV2ResponseType, error) {
	response := new(GeoLocationInternationalAddressCoordinatesLatLonV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationInternationalAddressCoordinatesLatLonV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationInternationalLatLonToAddress(request *GeoLocationInternationalLatLonToAddressRequestType) (*GeoLocationInternationalLatLonToAddressResponseType, error) {
	response := new(GeoLocationInternationalLatLonToAddressResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationInternationalLatLonToAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationNeighborhoodDistance(request *GeoLocationNeighborhoodDistanceRequestType) (*GeoLocationNeighborhoodDistanceResponseType, error) {
	response := new(GeoLocationNeighborhoodDistanceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationNeighborhoodDistance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationPostcodeDistance(request *GeoLocationPostcodeDistanceRequestType) (*GeoLocationPostcodeDistanceResponseType, error) {
	response := new(GeoLocationPostcodeDistanceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationPostcodeDistance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationHaversineDistance(request *GeoLocationHaversineDistanceRequestType) (*GeoLocationHaversineDistanceResponseType, error) {
	response := new(GeoLocationHaversineDistanceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationHaversineDistance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationDistanceSortedNeighborhoodCodes(request *GeoLocationDistanceSortedNeighborhoodCodesRequestType) (*GeoLocationDistanceSortedNeighborhoodCodesResponseType, error) {
	response := new(GeoLocationDistanceSortedNeighborhoodCodesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationDistanceSortedNeighborhoodCodes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationDistanceSortedNeighborhoodCodesRadius(request *GeoLocationDistanceSortedNeighborhoodCodesRadiusRequestType) (*GeoLocationDistanceSortedNeighborhoodCodesRadiusResponseType, error) {
	response := new(GeoLocationDistanceSortedNeighborhoodCodesRadiusResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationDistanceSortedNeighborhoodCodesRadius", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationDistanceSortedPostcodesRadius(request *GeoLocationDistanceSortedPostcodesRadiusRequestType) (*GeoLocationDistanceSortedPostcodesRadiusResponseType, error) {
	response := new(GeoLocationDistanceSortedPostcodesRadiusResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationDistanceSortedPostcodesRadius", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationDistanceSortedPostcodes(request *GeoLocationDistanceSortedPostcodesRequestType) (*GeoLocationDistanceSortedPostcodesResponseType, error) {
	response := new(GeoLocationDistanceSortedPostcodesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationDistanceSortedPostcodes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationLatLonToRD(request *GeoLocationLatLonToRDRequestType) (*GeoLocationLatLonToRDResponseType, error) {
	response := new(GeoLocationLatLonToRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationLatLonToRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GeoLocationRDToLatLon(request *GeoLocationRDToLatLonRequestType) (*GeoLocationRDToLatLonResponseType, error) {
	response := new(GeoLocationRDToLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/geoLocationRDToLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditGetReport(request *GraydonCreditGetReportRequestType) (*GraydonCreditGetReportResponseType, error) {
	response := new(GraydonCreditGetReportResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditGetReport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditSearchIdentification(request *GraydonCreditSearchIdentificationRequestType) (*GraydonCreditSearchIdentificationResponseType, error) {
	response := new(GraydonCreditSearchIdentificationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditSearchIdentification", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditSearchName(request *GraydonCreditSearchNameRequestType) (*GraydonCreditSearchNameResponseType, error) {
	response := new(GraydonCreditSearchNameResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditSearchName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditSearchPostcode(request *GraydonCreditSearchPostcodeRequestType) (*GraydonCreditSearchPostcodeResponseType, error) {
	response := new(GraydonCreditSearchPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditSearchPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditQuickscan(request *GraydonCreditQuickscanRequestType) (*GraydonCreditQuickscanResponseType, error) {
	response := new(GraydonCreditQuickscanResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditQuickscan", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditRatings(request *GraydonCreditRatingsRequestType) (*GraydonCreditRatingsResponseType, error) {
	response := new(GraydonCreditRatingsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditRatings", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditVatNumber(request *GraydonCreditVatNumberRequestType) (*GraydonCreditVatNumberResponseType, error) {
	response := new(GraydonCreditVatNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditVatNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditCompanyLiaisons(request *GraydonCreditCompanyLiaisonsRequestType) (*GraydonCreditCompanyLiaisonsResponseType, error) {
	response := new(GraydonCreditCompanyLiaisonsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditCompanyLiaisons", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) GraydonCreditManagement(request *GraydonCreditManagementRequestType) (*GraydonCreditManagementResponseType, error) {
	response := new(GraydonCreditManagementResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/graydonCreditManagement", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) InsolvencyGetCaseByPublication(request *InsolvencyGetCaseByPublicationRequestType) (*InsolvencyGetCaseByPublicationResponseType, error) {
	response := new(InsolvencyGetCaseByPublicationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/insolvencyGetCaseByPublication", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) InsolvencySearchPublicationsByCoCNumber(request *InsolvencySearchPublicationsByCoCNumberRequestType) (*InsolvencySearchPublicationsByCoCNumberResponseType, error) {
	response := new(InsolvencySearchPublicationsByCoCNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/insolvencySearchPublicationsByCoCNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) InsolvencySearchPublicationsByPerson(request *InsolvencySearchPublicationsByPersonRequestType) (*InsolvencySearchPublicationsByPersonResponseType, error) {
	response := new(InsolvencySearchPublicationsByPersonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/insolvencySearchPublicationsByPerson", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) InternationalAddressSearchV2(request *InternationalAddressSearchV2RequestType) (*InternationalAddressSearchV2ResponseType, error) {
	response := new(InternationalAddressSearchV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/internationalAddressSearchV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) InternationalAddressSearchInteractive(request *InternationalAddressSearchInteractiveRequestType) (*InternationalAddressSearchInteractiveResponseType, error) {
	response := new(InternationalAddressSearchInteractiveResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/internationalAddressSearchInteractive", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterEigendomsBerichtDocumentPerceel(request *KadasterEigendomsBerichtDocumentPerceelRequestType) (*KadasterEigendomsBerichtDocumentPerceelResponseType, error) {
	response := new(KadasterEigendomsBerichtDocumentPerceelResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterEigendomsBerichtDocumentPerceel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterEigendomsBerichtDocumentPostcode(request *KadasterEigendomsBerichtDocumentPostcodeRequestType) (*KadasterEigendomsBerichtDocumentPostcodeResponseType, error) {
	response := new(KadasterEigendomsBerichtDocumentPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterEigendomsBerichtDocumentPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterEigendomsBerichtPerceelV2(request *KadasterEigendomsBerichtPerceelV2RequestType) (*KadasterEigendomsBerichtPerceelV2ResponseType, error) {
	response := new(KadasterEigendomsBerichtPerceelV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterEigendomsBerichtPerceelV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterEigendomsBerichtPostcodeV2(request *KadasterEigendomsBerichtPostcodeV2RequestType) (*KadasterEigendomsBerichtPostcodeV2ResponseType, error) {
	response := new(KadasterEigendomsBerichtPostcodeV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterEigendomsBerichtPostcodeV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterKoopsommenOverzichtV2(request *KadasterKoopsommenOverzichtV2RequestType) (*KadasterKoopsommenOverzichtV2ResponseType, error) {
	response := new(KadasterKoopsommenOverzichtV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterKoopsommenOverzichtV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterUittrekselKadastraleKaartPerceelV3(request *KadasterUittrekselKadastraleKaartPerceelV3RequestType) (*KadasterUittrekselKadastraleKaartPerceelV3ResponseType, error) {
	response := new(KadasterUittrekselKadastraleKaartPerceelV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterUittrekselKadastraleKaartPerceelV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterUittrekselKadastraleKaartPostcodeV3(request *KadasterUittrekselKadastraleKaartPostcodeV3RequestType) (*KadasterUittrekselKadastraleKaartPostcodeV3ResponseType, error) {
	response := new(KadasterUittrekselKadastraleKaartPostcodeV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterUittrekselKadastraleKaartPostcodeV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterHypothecairBerichtPostcodeV3(request *KadasterHypothecairBerichtPostcodeV3RequestType) (*KadasterHypothecairBerichtPostcodeV3ResponseType, error) {
	response := new(KadasterHypothecairBerichtPostcodeV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterHypothecairBerichtPostcodeV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterHypothecairBerichtPerceelV3(request *KadasterHypothecairBerichtPerceelV3RequestType) (*KadasterHypothecairBerichtPerceelV3ResponseType, error) {
	response := new(KadasterHypothecairBerichtPerceelV3ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterHypothecairBerichtPerceelV3", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterBronDocument(request *KadasterBronDocumentRequestType) (*KadasterBronDocumentResponseType, error) {
	response := new(KadasterBronDocumentResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterBronDocument", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterValueListGetRanges(request *KadasterValueListGetRangesRequestType) (*KadasterValueListGetRangesResponseType, error) {
	response := new(KadasterValueListGetRangesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterValueListGetRanges", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterValueListGetMunicipalities(request *KadasterValueListGetMunicipalitiesRequestType) (*KadasterValueListGetMunicipalitiesResponseType, error) {
	response := new(KadasterValueListGetMunicipalitiesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterValueListGetMunicipalities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterAddressCoordinates(request *KadasterAddressCoordinatesRequestType) (*KadasterAddressCoordinatesResponseType, error) {
	response := new(KadasterAddressCoordinatesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterAddressCoordinates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterKadastraleKaartPerceelV2(request *KadasterKadastraleKaartPerceelV2RequestType) (*KadasterKadastraleKaartPerceelV2ResponseType, error) {
	response := new(KadasterKadastraleKaartPerceelV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterKadastraleKaartPerceelV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterKadastraleKaartPostcodeV2(request *KadasterKadastraleKaartPostcodeV2RequestType) (*KadasterKadastraleKaartPostcodeV2ResponseType, error) {
	response := new(KadasterKadastraleKaartPostcodeV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterKadastraleKaartPostcodeV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterUittrekselKadastraleKaartPerceel(request *KadasterUittrekselKadastraleKaartPerceelRequestType) (*KadasterUittrekselKadastraleKaartPerceelResponseType, error) {
	response := new(KadasterUittrekselKadastraleKaartPerceelResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterUittrekselKadastraleKaartPerceel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterUittrekselKadastraleKaartPostcode(request *KadasterUittrekselKadastraleKaartPostcodeRequestType) (*KadasterUittrekselKadastraleKaartPostcodeResponseType, error) {
	response := new(KadasterUittrekselKadastraleKaartPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterUittrekselKadastraleKaartPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterHypothecairBerichtPostcode(request *KadasterHypothecairBerichtPostcodeRequestType) (*KadasterHypothecairBerichtPostcodeResponseType, error) {
	response := new(KadasterHypothecairBerichtPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterHypothecairBerichtPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterHypothecairBerichtPerceel(request *KadasterHypothecairBerichtPerceelRequestType) (*KadasterHypothecairBerichtPerceelResponseType, error) {
	response := new(KadasterHypothecairBerichtPerceelResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterHypothecairBerichtPerceel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterEigendomsBerichtPerceel(request *KadasterEigendomsBerichtPerceelRequestType) (*KadasterEigendomsBerichtPerceelResponseType, error) {
	response := new(KadasterEigendomsBerichtPerceelResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterEigendomsBerichtPerceel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterEigendomsBerichtPostcode(request *KadasterEigendomsBerichtPostcodeRequestType) (*KadasterEigendomsBerichtPostcodeResponseType, error) {
	response := new(KadasterEigendomsBerichtPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterEigendomsBerichtPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterUittrekselKadastraleKaartPerceelV2(request *KadasterUittrekselKadastraleKaartPerceelV2RequestType) (*KadasterUittrekselKadastraleKaartPerceelV2ResponseType, error) {
	response := new(KadasterUittrekselKadastraleKaartPerceelV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterUittrekselKadastraleKaartPerceelV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterUittrekselKadastraleKaartPostcodeV2(request *KadasterUittrekselKadastraleKaartPostcodeV2RequestType) (*KadasterUittrekselKadastraleKaartPostcodeV2ResponseType, error) {
	response := new(KadasterUittrekselKadastraleKaartPostcodeV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterUittrekselKadastraleKaartPostcodeV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterKadastraleKaartPerceel(request *KadasterKadastraleKaartPerceelRequestType) (*KadasterKadastraleKaartPerceelResponseType, error) {
	response := new(KadasterKadastraleKaartPerceelResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterKadastraleKaartPerceel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterKadastraleKaartPostcode(request *KadasterKadastraleKaartPostcodeRequestType) (*KadasterKadastraleKaartPostcodeResponseType, error) {
	response := new(KadasterKadastraleKaartPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterKadastraleKaartPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterHypothecairBerichtPostcodeV2(request *KadasterHypothecairBerichtPostcodeV2RequestType) (*KadasterHypothecairBerichtPostcodeV2ResponseType, error) {
	response := new(KadasterHypothecairBerichtPostcodeV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterHypothecairBerichtPostcodeV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterHypothecairBerichtPerceelV2(request *KadasterHypothecairBerichtPerceelV2RequestType) (*KadasterHypothecairBerichtPerceelV2ResponseType, error) {
	response := new(KadasterHypothecairBerichtPerceelV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterHypothecairBerichtPerceelV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterKoopsommenOverzicht(request *KadasterKoopsommenOverzichtRequestType) (*KadasterKoopsommenOverzichtResponseType, error) {
	response := new(KadasterKoopsommenOverzichtResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterKoopsommenOverzicht", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetKadastraalBerichtObjectByAdres(request *KadasterV2GetKadastraalBerichtObjectByAdresRequestType) (*KadasterV2GetKadastraalBerichtObjectByAdresResponseType, error) {
	response := new(KadasterV2GetKadastraalBerichtObjectByAdresResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetKadastraalBerichtObjectByAdres", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetKadastraalBerichtObjectByKadastraleAanduiding(request *KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingRequestType) (*KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingResponseType, error) {
	response := new(KadasterV2GetKadastraalBerichtObjectByKadastraleAanduidingResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetKadastraalBerichtObjectByKadastraleAanduiding", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetKadastraalBerichtObjectByObjectId(request *KadasterV2GetKadastraalBerichtObjectByObjectIdRequestType) (*KadasterV2GetKadastraalBerichtObjectByObjectIdResponseType, error) {
	response := new(KadasterV2GetKadastraalBerichtObjectByObjectIdResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetKadastraalBerichtObjectByObjectId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetHypothecairBerichtObjectByAdres(request *KadasterV2GetHypothecairBerichtObjectByAdresRequestType) (*KadasterV2GetHypothecairBerichtObjectByAdresResponseType, error) {
	response := new(KadasterV2GetHypothecairBerichtObjectByAdresResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetHypothecairBerichtObjectByAdres", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetHypothecairBerichtObjectByKadastraleAanduiding(request *KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingRequestType) (*KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingResponseType, error) {
	response := new(KadasterV2GetHypothecairBerichtObjectByKadastraleAanduidingResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetHypothecairBerichtObjectByKadastraleAanduiding", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetHypothecairBerichtObjectByObjectId(request *KadasterV2GetHypothecairBerichtObjectByObjectIdRequestType) (*KadasterV2GetHypothecairBerichtObjectByObjectIdResponseType, error) {
	response := new(KadasterV2GetHypothecairBerichtObjectByObjectIdResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetHypothecairBerichtObjectByObjectId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetUittrekselKadastraleKaartByAdres(request *KadasterV2GetUittrekselKadastraleKaartByAdresRequestType) (*KadasterV2GetUittrekselKadastraleKaartByAdresResponseType, error) {
	response := new(KadasterV2GetUittrekselKadastraleKaartByAdresResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetUittrekselKadastraleKaartByAdres", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduiding(request *KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingRequestType) (*KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingResponseType, error) {
	response := new(KadasterV2GetUittrekselKadastraleKaartByKadastraleAanduidingResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetUittrekselKadastraleKaartByKadastraleAanduiding", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetUittrekselKadastraleKaartByObjectId(request *KadasterV2GetUittrekselKadastraleKaartByObjectIdRequestType) (*KadasterV2GetUittrekselKadastraleKaartByObjectIdResponseType, error) {
	response := new(KadasterV2GetUittrekselKadastraleKaartByObjectIdResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetUittrekselKadastraleKaartByObjectId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetKoopsommenOverzicht(request *KadasterV2GetKoopsommenOverzichtRequestType) (*KadasterV2GetKoopsommenOverzichtResponseType, error) {
	response := new(KadasterV2GetKoopsommenOverzichtResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetKoopsommenOverzicht", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KadasterV2GetBronDocument(request *KadasterV2GetBronDocumentRequestType) (*KadasterV2GetBronDocumentResponseType, error) {
	response := new(KadasterV2GetBronDocumentResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kadasterV2GetBronDocument", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkGetDossier(request *KvkGetDossierRequestType) (*KvkGetDossierResponseType, error) {
	response := new(KvkGetDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkGetDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkSearchDossierNumber(request *KvkSearchDossierNumberRequestType) (*KvkSearchDossierNumberResponseType, error) {
	response := new(KvkSearchDossierNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkSearchDossierNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkSearchParameters(request *KvkSearchParametersRequestType) (*KvkSearchParametersResponseType, error) {
	response := new(KvkSearchParametersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkSearchParameters", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkSearchPostcode(request *KvkSearchPostcodeRequestType) (*KvkSearchPostcodeResponseType, error) {
	response := new(KvkSearchPostcodeResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkSearchPostcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkSearchSelection(request *KvkSearchSelectionRequestType) (*KvkSearchSelectionResponseType, error) {
	response := new(KvkSearchSelectionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkSearchSelection", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkGetExtractDocument(request *KvkGetExtractDocumentRequestType) (*KvkGetExtractDocumentResponseType, error) {
	response := new(KvkGetExtractDocumentResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkGetExtractDocument", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkUpdateCheckDossier(request *KvkUpdateCheckDossierRequestType) (*KvkUpdateCheckDossierResponseType, error) {
	response := new(KvkUpdateCheckDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkUpdateCheckDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkUpdateGetChangedDossiers(request *KvkUpdateGetChangedDossiersRequestType) (*KvkUpdateGetChangedDossiersResponseType, error) {
	response := new(KvkUpdateGetChangedDossiersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkUpdateGetChangedDossiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkUpdateGetDossiers(request *KvkUpdateGetDossiersRequestType) (*KvkUpdateGetDossiersResponseType, error) {
	response := new(KvkUpdateGetDossiersResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkUpdateGetDossiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkUpdateAddDossier(request *KvkUpdateAddDossierRequestType) (*KvkUpdateAddDossierResponseType, error) {
	response := new(KvkUpdateAddDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkUpdateAddDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) KvkUpdateRemoveDossier(request *KvkUpdateRemoveDossierRequestType) (*KvkUpdateRemoveDossierResponseType, error) {
	response := new(KvkUpdateRemoveDossierResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/kvkUpdateRemoveDossier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MapViewPostcodeV2(request *MapViewPostcodeV2RequestType) (*MapViewPostcodeV2ResponseType, error) {
	response := new(MapViewPostcodeV2ResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/mapViewPostcodeV2", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MapViewLatLon(request *MapViewLatLonRequestType) (*MapViewLatLonResponseType, error) {
	response := new(MapViewLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/mapViewLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MapViewRD(request *MapViewRDRequestType) (*MapViewRDResponseType, error) {
	response := new(MapViewRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/mapViewRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MapViewInternationalLatLon(request *MapViewInternationalLatLonRequestType) (*MapViewInternationalLatLonResponseType, error) {
	response := new(MapViewInternationalLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/mapViewInternationalLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MetaGetServices(request *MetaGetServicesRequestType) (*MetaGetServicesResponseType, error) {
	response := new(MetaGetServicesResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/metaGetServices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MetaGetService(request *MetaGetServiceRequestType) (*MetaGetServiceResponseType, error) {
	response := new(MetaGetServiceResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/metaGetService", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) MetaGetMethod(request *MetaGetMethodRequestType) (*MetaGetMethodResponseType, error) {
	response := new(MetaGetMethodResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/metaGetMethod", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) NbwoEstimateValue(request *NbwoEstimateValueRequestType) (*NbwoEstimateValueResponseType, error) {
	response := new(NbwoEstimateValueResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/nbwoEstimateValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RiskCheckPerson(request *RiskCheckPersonRequestType) (*RiskCheckPersonResponseType, error) {
	response := new(RiskCheckPersonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/riskCheckPerson", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RiskCheckGetRiskPersonCompanyReport(request *RiskCheckGetRiskPersonCompanyReportRequestType) (*RiskCheckGetRiskPersonCompanyReportResponseType, error) {
	response := new(RiskCheckGetRiskPersonCompanyReportResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/riskCheckGetRiskPersonCompanyReport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerGetRoute(request *RoutePlannerGetRouteRequestType) (*RoutePlannerGetRouteResponseType, error) {
	response := new(RoutePlannerGetRouteResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerGetRoute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerInformationAddress(request *RoutePlannerInformationAddressRequestType) (*RoutePlannerInformationAddressResponseType, error) {
	response := new(RoutePlannerInformationAddressResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerInformationAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerDescriptionAddress(request *RoutePlannerDescriptionAddressRequestType) (*RoutePlannerDescriptionAddressResponseType, error) {
	response := new(RoutePlannerDescriptionAddressResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerDescriptionAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerDescription(request *RoutePlannerDescriptionRequestType) (*RoutePlannerDescriptionResponseType, error) {
	response := new(RoutePlannerDescriptionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerDescriptionShortest(request *RoutePlannerDescriptionShortestRequestType) (*RoutePlannerDescriptionShortestResponseType, error) {
	response := new(RoutePlannerDescriptionShortestResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerDescriptionShortest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerDescriptionCoordinatesRD(request *RoutePlannerDescriptionCoordinatesRDRequestType) (*RoutePlannerDescriptionCoordinatesRDResponseType, error) {
	response := new(RoutePlannerDescriptionCoordinatesRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerDescriptionCoordinatesRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerInformation(request *RoutePlannerInformationRequestType) (*RoutePlannerInformationResponseType, error) {
	response := new(RoutePlannerInformationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerRDDescription(request *RoutePlannerRDDescriptionRequestType) (*RoutePlannerRDDescriptionResponseType, error) {
	response := new(RoutePlannerRDDescriptionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerRDDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerRDInformation(request *RoutePlannerRDInformationRequestType) (*RoutePlannerRDInformationResponseType, error) {
	response := new(RoutePlannerRDInformationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerRDInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerRDDescriptionCoordinatesRD(request *RoutePlannerRDDescriptionCoordinatesRDRequestType) (*RoutePlannerRDDescriptionCoordinatesRDResponseType, error) {
	response := new(RoutePlannerRDDescriptionCoordinatesRDResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerRDDescriptionCoordinatesRD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerInformationDutchAddress(request *RoutePlannerInformationDutchAddressRequestType) (*RoutePlannerInformationDutchAddressResponseType, error) {
	response := new(RoutePlannerInformationDutchAddressResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerInformationDutchAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerDescriptionDutchAddress(request *RoutePlannerDescriptionDutchAddressRequestType) (*RoutePlannerDescriptionDutchAddressResponseType, error) {
	response := new(RoutePlannerDescriptionDutchAddressResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerDescriptionDutchAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerEUDescription(request *RoutePlannerEUDescriptionRequestType) (*RoutePlannerEUDescriptionResponseType, error) {
	response := new(RoutePlannerEUDescriptionResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerEUDescription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerEUInformation(request *RoutePlannerEUInformationRequestType) (*RoutePlannerEUInformationResponseType, error) {
	response := new(RoutePlannerEUInformationResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerEUInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerEUDescriptionCoordinatesLatLon(request *RoutePlannerEUDescriptionCoordinatesLatLonRequestType) (*RoutePlannerEUDescriptionCoordinatesLatLonResponseType, error) {
	response := new(RoutePlannerEUDescriptionCoordinatesLatLonResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerEUDescriptionCoordinatesLatLon", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) RoutePlannerEUMap(request *RoutePlannerEUMapRequestType) (*RoutePlannerEUMapResponseType, error) {
	response := new(RoutePlannerEUMapResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/routePlannerEUMap", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) SepaConvertBankAccountNumber(request *SepaConvertBankAccountNumberRequestType) (*SepaConvertBankAccountNumberResponseType, error) {
	response := new(SepaConvertBankAccountNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/sepaConvertBankAccountNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) SepaValidateInternationalBankAccountNumberFormat(request *SepaValidateInternationalBankAccountNumberFormatRequestType) (*SepaValidateInternationalBankAccountNumberFormatResponseType, error) {
	response := new(SepaValidateInternationalBankAccountNumberFormatResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/sepaValidateInternationalBankAccountNumberFormat", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) SepaIbanDetails(request *SepaIbanDetailsRequestType) (*SepaIbanDetailsResponseType, error) {
	response := new(SepaIbanDetailsResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/sepaIbanDetails", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) SepaMatchAccountHolder(request *SepaMatchAccountHolderRequestType) (*SepaMatchAccountHolderResponseType, error) {
	response := new(SepaMatchAccountHolderResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/sepaMatchAccountHolder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) SepaConvertBasicBankAccountNumber(request *SepaConvertBasicBankAccountNumberRequestType) (*SepaConvertBasicBankAccountNumberResponseType, error) {
	response := new(SepaConvertBasicBankAccountNumberResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/sepaConvertBasicBankAccountNumber", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) VatValidate(request *VatValidateRequestType) (*VatValidateResponseType, error) {
	response := new(VatValidateResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/vatValidate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *webservicesNLPortType) VatViesProxyCheckVat(request *VatViesProxyCheckVatRequestType) (*VatViesProxyCheckVatResponseType, error) {
	response := new(VatViesProxyCheckVatResponseType)
	err := service.client.Call("https://api.webservices.nl/soap_doclit.php/vatViesProxyCheckVat", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
