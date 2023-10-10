package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func createMapFromRawData(data [][]string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	var fieldNames []string
	for i, line := range data {
		if i == 0 {
			fieldNames = line
			continue
		}
		attachment1Dot1Transformers := make(map[string]interface{})
		for j, field := range line {
			fieldName := fieldNames[j]
			attachment1Dot1Transformers[fieldName] = field
		}
		result = append(result, attachment1Dot1Transformers)
	}
	return result
}

func findMapAttachment1Point(obj *map[string]interface{}, data []map[string]interface{}) error {
	find := false
	mapAttachment1Point := make([]map[string]interface{}, 0)
	for _, record := range data {
		if record["contract_number"] == (*obj)["contract_number"] {
			mapAttachment1Point = append(mapAttachment1Point, record)
			find = true
		}
	}
	if !find {
		log.Printf("[WARNING]: findMapAttachment1Point(): no found any record in Attachment1Dot1Transformers for contract =  %v \n\t", (*obj)["contract_number"])
	}
	(*obj)["counters"] = mapAttachment1Point
	return nil
}

func findMapAttachment1Dot1Transformers(obj *map[string]interface{}, data []map[string]interface{}) error {
	find := false
	mapAttachment1Dot1Transformers := make([]map[string]interface{}, 0)
	for _, record := range data {
		if record["contract_number"] == (*obj)["contract_number"] && record["device_number"] == (*obj)["device_number"] {
			mapAttachment1Dot1Transformers = append(mapAttachment1Dot1Transformers, record)
			find = true
		}
	}
	if !find {
		log.Printf("[WARNING]: findMapAttachment1Dot1Transformers(): no found any record in Attachment1Dot1Transformers for contract = %v and name, device = %v\n\t", (*obj)["contract_number"], (*obj)["device_number"])
	}
	(*obj)["transformers"] = mapAttachment1Dot1Transformers
	return nil
}

func getRawDataFromCSV(file string) ([][]string, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("getRawDataFromCSV(): %v\n\t", err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("getRawDataFromCSV(): %v\n\t", err)
	}

	return data, nil
}

func parse() ([]map[string]interface{}, error) {

	dataContracts, err := getRawDataFromCSV("contracts.csv")
	if err != nil {
		log.Fatal(err)
	}
	dataAttachment1, err := getRawDataFromCSV("counters.csv")
	if err != nil {
		log.Fatal(err)
	}
	dataAttachment1Dot1, err := getRawDataFromCSV("transformers.csv")
	if err != nil {
		log.Fatal(err)
	}

	mapDataContracts := createMapFromRawData(dataContracts)
	mapDataAttachment1 := createMapFromRawData(dataAttachment1)
	mapDataAttachment1Dot1 := createMapFromRawData(dataAttachment1Dot1)

	for i, record := range mapDataAttachment1 {
		err = findMapAttachment1Dot1Transformers(&record, mapDataAttachment1Dot1)
		if err != nil {
			log.Println(err)
		}
		mapDataAttachment1[i] = record
	}

	for i, record := range mapDataContracts {
		err = findMapAttachment1Point(&record, mapDataAttachment1)
		if err != nil {
			log.Println(err)
		}
		mapDataContracts[i] = record
	}

	return mapDataContracts, nil
}

func download(link string, format string, filename string) {

	log.Println(link, format, filename)

	pass := base64.StdEncoding.EncodeToString([]byte("SiyarovRS@atomsbt.ru:1xz6991Z"))

	postURL := fmt.Sprintf("%v?DocumentFormat=%v", link, format)

	r, err := http.NewRequest("GET", postURL, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Authorization", fmt.Sprintf("Basic %v", pass))

	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Println(string(resp))
		// log.Println(len(resp))
		log.Fatal(res.StatusCode)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(resp)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}

func unpackField(items []map[string]interface{}, fieldName string) []string {
	result := make([]string, 0)
	for _, record := range items {
		result = append(result, record[fieldName].(string))
	}
	return result
}

func unpackFieldIfFieldKeyExist(items []map[string]interface{}, fieldName string, fieldKey string) []string {
	result := make([]string, 0)
	for _, record := range items {
		if record[fieldKey].(string) != "" {
			result = append(result, record[fieldName].(string))
		}
	}
	return result
}

func unpackIndex(items []map[string]interface{}) []string {
	result := make([]string, 0)
	for i, _ := range items {
		result = append(result, fmt.Sprintf("%v", i+1))
	}
	return result
}

func valOrDef(value interface{}, def string) string {
	valueStr, ok := value.(string)
	if !ok {
		return def
	}
	if valueStr == "" {
		return def
	}
	return valueStr
}

func makeDocument(body []byte, templateID string) (string, string, error) {
	pass := base64.StdEncoding.EncodeToString([]byte("SiyarovRS@atomsbt.ru:1xz6991Z"))

	// HTTP endpoint
	postURL := "https://atomsbt.doc.one/api/v3/documents?TemplateID=" + templateID

	r, err := http.NewRequest("POST", postURL, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Authorization", fmt.Sprintf("Basic %v", pass))
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Println(string(resp))
		log.Fatal(res.StatusCode)
	}

	rData := make(map[string]interface{})
	err = json.Unmarshal(resp, &rData)
	if err != nil {
		log.Fatal(err)
	}

	link, ok := rData["Link"].(string)
	if !ok {
		log.Fatal(fmt.Errorf("invalid data conversion"))
	}

	format, ok := rData["DownloadFormats"].(string)
	if !ok {
		log.Fatal(fmt.Errorf("invalid data conversion"))
	}

	return link, format, nil
}

func makeContract(record map[string]interface{}) (string, string, error) {
	court := "согласно действующему законодательству"
	if record["city"].(string) == "Донецк" {
		court = "Донецкой Народной Республики"
	} else if record["city"].(string) == "Луганск" {
		court = "Луганской Народной Республики"
	}

	body, err := json.Marshal(map[string]interface{}{
		"Data": map[string][]string{
			"arbitration_court":           {court},
			"city":                        {record["city"].(string)},
			"consignee_address":           {record["consignee_address"].(string)},
			"consignee_birth":             {record["consignee_birth"].(string)},
			"consignee_birth_place":       {record["consignee_birth_place"].(string)},
			"consignee_cor_address":       {record["consignee_cor_address"].(string)},
			"consignee_email":             {record["consignee_email"].(string)},
			"consignee_inn":               {record["consignee_inn"].(string)},
			"consignee_live_address":      {record["consignee_live_address"].(string)},
			"consignee_mobile":            {record["consignee_mobile"].(string)},
			"consignee_name":              {record["consignee_name"].(string)},
			"consignee_other":             {record["consignee_other"].(string)},
			"consignee_passport":          {record["consignee_passport"].(string)},
			"consignee_passport_date":     {record["consignee_passport_date"].(string)},
			"consignee_phone":             {record["consignee_phone"].(string)},
			"consignee_place_address":     {record["consignee_place_address"].(string)},
			"consignee_place_email":       {record["consignee_place_email"].(string)},
			"consignee_place_phone":       {record["consignee_place_phone"].(string)},
			"consignee_reg_address":       {record["consignee_reg_address"].(string)},
			"contract_cost_number":        {record["contract_cost_number"].(string)},
			"contract_cost_string":        {record["contract_cost_string"].(string)},
			"contract_date":               {"«01» сентября 2023 г"}, // статично
			"contract_date_from":          {"«01» октября 2023 г"},  // статично
			"contract_date_to":            {"«__» ____________ г."}, // статично
			"contract_nds_number":         {record["contract_nds_number"].(string)},
			"contract_nds_string":         {record["contract_nds_string"].(string)},
			"contract_number":             {record["contract_number"].(string)},
			"contract_out_number":         {record["contract_out_number"].(string)},
			"contract_out_string":         {record["contract_out_string"].(string)},
			"contract_plan":               {record["contract_plan"].(string)},
			"organization_address":        {record["organization_address"].(string)},
			"organization_address_before": {record["organization_address_before"].(string)},
			"organization_bank":           {valOrDef(record["organization_bank"].(string), "_")},
			"organization_bank_bik":       {valOrDef(record["organization_bank_bik"].(string), "_")},
			"organization_document":       {record["organization_document"].(string)},
			"organization_email":          {record["organization_email"].(string)},
			"organization_fio":            {record["organization_fio"].(string)},
			"organization_fullname":       {record["organization_fullname"].(string)},
			"organization_inn":            {record["organization_inn"].(string)},
			"organization_invoice":        {valOrDef(record["organization_invoice"].(string), "_")},
			"organization_kpp":            {valOrDef(record["organization_kpp"].(string), "_")},
			"organization_ks":             {valOrDef(record["organization_ks"].(string), "_")},
			"organization_mobile":         {record["organization_mobile"].(string)},
			"organization_name":           {record["organization_name"].(string)},
			"organization_ogrn":           {record["organization_ogrn"].(string)},
			"organization_okpo":           {valOrDef(record["organization_okpo"].(string), "_")},
			"organization_okved":          {record["organization_okved"].(string)},
			"organization_phone":          {record["organization_phone"].(string)},
			"organization_post":           {record["organization_post"].(string)},
			"organization_post_address":   {record["organization_post_address"].(string)},
			"organization_rs":             {valOrDef(record["organization_rs"].(string), "_")},
			"provider_ca_address":         {"115432, г. Москва, вн. тер. г. муниципальный округ Даниловский, проезд Проектируемый 4062-й, д. 6 стр. 25, помещ. 1н/6"},
			"provider_ca_document":        {valOrDef(record["provider_ca_document"], "__________________________________________________")},
			"provider_ca_fio":             {valOrDef(record["provider_ca_fio"], "__________________________________________________")},
			"provider_ca_fullname":        {"Общество с ограниченной ответственностью «Энергосбыт Донецк»"},
			"provider_ca_inn":             {"9725129199"},
			"provider_ca_kpp":             {"772501001"},
			"provider_ca_name":            {"ООО «Энергосбыт Донецк»"},
			"provider_ca_post":            {valOrDef(record["provider_ca_post"], "__________________________________________________")},
			"provider_document":           {valOrDef(record["provider_f_rs"], "__________________________________________________")},
			"provider_f_address":          {"Донецкая Народная Республика, городской округ Донецк, г. Донецк, ул. Щорса, д. 87"},
			"provider_f_bank":             {valOrDef(record["provider_f_bank"], "_")},
			"provider_f_bik":              {valOrDef(record["provider_f_bik"], "_")},
			"provider_f_email":            {"info@Donetsk.e-sbt.ru"},
			"provider_f_inn":              {"9725129199"},
			"provider_f_kpp":              {"930945001"},
			"provider_f_ks":               {valOrDef(record["provider_f_ks"], "_")},
			"provider_f_name":             {"Обособленное подразделение «Энергосбыт Донецк»"},
			"provider_f_ogrn":             {valOrDef(record["provider_f_ogrn"], "_")},
			"provider_f_okpo":             {valOrDef(record["provider_f_okpo"], "_")},
			"provider_f_okved":            {"35.14"},
			"provider_f_phone":            {valOrDef(record["provider_f_phone"], "_")},
			"provider_f_rs":               {valOrDef(record["provider_f_rs"], "_")},
			"provider_f_www":              {"Donetsk.e-sbt.ru"},
			"provider_post":               {valOrDef(record["provider_f_rs"], "__________________________________________________")},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return makeDocument(body, "3034c7a4-40c2-11ee-a980-02c163b943db")
}

func makeAttachment11(record map[string]interface{}, transformers []map[string]interface{}) (string, string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"Data": map[string][]string{
			"contract_date":               {record["contract_date"].(string)},
			"contract_number":             {record["contract_number"].(string)},
			"hours_per_month_2":           unpackFieldIfFieldKeyExist(transformers, "hours_per_month_2", "scheme_2"),
			"hours_per_month_3":           unpackFieldIfFieldKeyExist(transformers, "hours_per_month_3", "scheme_3"),
			"hours_per_month_line":        unpackFieldIfFieldKeyExist(transformers, "hours_per_month_line", "scheme_line"),
			"idle_2":                      unpackFieldIfFieldKeyExist(transformers, "idle_2", "scheme_2"),
			"idle_3":                      unpackFieldIfFieldKeyExist(transformers, "idle_3", "scheme_3"),
			"length_km":                   unpackFieldIfFieldKeyExist(transformers, "length_km", "scheme_line"),
			"mark":                        unpackFieldIfFieldKeyExist(transformers, "mark", "scheme_line"),
			"name_2":                      unpackFieldIfFieldKeyExist(transformers, "name_2", "scheme_2"),
			"name_3":                      unpackFieldIfFieldKeyExist(transformers, "name_3", "scheme_3"),
			"name_line":                   unpackFieldIfFieldKeyExist(transformers, "name_line", "scheme_line"),
			"nominal_power_3":             unpackFieldIfFieldKeyExist(transformers, "nominal_power_3", "scheme_line"),
			"number_of_parallel_circuits": unpackFieldIfFieldKeyExist(transformers, "number_of_parallel_circuits", "scheme_line"),
			"power_2":                     unpackFieldIfFieldKeyExist(transformers, "power_2", "scheme_2"),
			"power_3":                     unpackFieldIfFieldKeyExist(transformers, "power_3", "scheme_3"),
			"power_line":                  unpackFieldIfFieldKeyExist(transformers, "power_line", "scheme_line"),
			"resistivity":                 unpackFieldIfFieldKeyExist(transformers, "resistivity", "scheme_line"),
			"scheme_2":                    unpackFieldIfFieldKeyExist(transformers, "scheme_2", "scheme_2"),
			"scheme_3":                    unpackFieldIfFieldKeyExist(transformers, "scheme_3", "scheme_3"),
			"scheme_line":                 unpackFieldIfFieldKeyExist(transformers, "scheme_line", "scheme_line"),
			"section":                     unpackFieldIfFieldKeyExist(transformers, "section", "scheme_line"),
			"short_2":                     unpackFieldIfFieldKeyExist(transformers, "short_2", "scheme_2"),
			"short_3_ch":                  unpackFieldIfFieldKeyExist(transformers, "short_3_ch", "scheme_3"),
			"short_3_hh":                  unpackFieldIfFieldKeyExist(transformers, "short_3_hh", "scheme_3"),
			"short_3_vh":                  unpackFieldIfFieldKeyExist(transformers, "short_3_vh", "scheme_3"),
			"voltage_2":                   unpackFieldIfFieldKeyExist(transformers, "voltage_2", "scheme_2"),
			"voltage_3":                   unpackFieldIfFieldKeyExist(transformers, "voltage_3", "scheme_3"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return makeDocument(body, "e8ff2b1c-40c7-11ee-a1c8-02c163b943db")
}

func makeAttachment1(record map[string]interface{}) (string, string, error) {
	counters := record["counters"].([]map[string]interface{})
	body, err := json.Marshal(map[string]interface{}{
		"Data": map[string][]string{
			"contract_date":                      {record["contract_date"].(string)},
			"contract_number":                    {record["contract_number"].(string)},
			"line":                               unpackIndex(counters),
			"accuracy_class":                     unpackField(counters, "accuracy_class"),
			"calibration_interval":               unpackField(counters, "calibration_interval"),
			"connection_scheme":                  unpackField(counters, "connection_scheme"),
			"kWh":                                unpackField(counters, "kWh"),
			"date_future":                        unpackField(counters, "date_future"),
			"date_past":                          unpackField(counters, "date_past"),
			"delivery_point":                     unpackField(counters, "delivery_point"),
			"estimated_coefficient":              unpackField(counters, "estimated_coefficient"),
			"estimated_voltage_level":            unpackField(counters, "estimated_voltage_level"),
			"max_point_1":                        unpackField(counters, "max_point_1"),
			"max_point_2":                        unpackField(counters, "max_point_2"),
			"name_address_device":                unpackField(counters, "name_address_device"),
			"note":                               unpackField(counters, "note"),
			"device_number":                      unpackField(counters, "device_number"),
			"place_of_installation_of_the_meter": unpackField(counters, "place_of_installation_of_the_meter"),
			"percent":                            unpackField(counters, "percent"),
			"purpose_of_accounting":              unpackField(counters, "purpose_of_accounting"),
			"significance":                       unpackField(counters, "significance"),
			"th_power":                           unpackField(counters, "th_power"),
			"tt_amperage":                        unpackField(counters, "tt_amperage"),
			"type_device":                        unpackField(counters, "type_device"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return makeDocument(body, "efbb2960-40c7-11ee-a34c-02c163b943db")
}

func makeAttachment2(record map[string]interface{}) (string, string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"Data": map[string][]string{
			"contract_calc":               {record["contract_calc"].(string)},
			"contract_calc_email":         {record["contract_calc_email"].(string)},
			"contract_calc_phone":         {record["contract_calc_phone"].(string)},
			"contract_curator":            {record["contract_curator"].(string)},
			"contract_curator_email":      {record["contract_curator_email"].(string)},
			"contract_curator_head":       {record["contract_curator_head"].(string)},
			"contract_curator_head_email": {record["contract_curator_head_email"].(string)},
			"contract_curator_head_fax":   {record["contract_curator_head_fax"].(string)},
			"contract_curator_head_phone": {record["contract_curator_head_phone"].(string)},
			"contract_curator_phone":      {record["contract_curator_phone"].(string)},
			"contract_date":               {record["contract_date"].(string)},
			"contract_number":             {record["contract_number"].(string)},
			"contract_tech_audit":         {record["contract_tech_audit"].(string)},
			"contract_tech_audit_email":   {record["contract_tech_audit_email"].(string)},
			"contract_tech_audit_phone":   {record["contract_tech_audit_phone"].(string)},
			"network_company1":            {record["network_company1"].(string)},
			"network_company1_email":      {record["network_company1_email"].(string)},
			"network_company1_fax":        {record["network_company1_fax"].(string)},
			"network_company1_phone":      {record["network_company1_phone"].(string)},
			"network_company1_place":      {record["network_company1_place"].(string)},
			"network_company1_post":       {record["network_company1_post"].(string)},
			"provider_count_address":      {record["provider_count_address"].(string)},
			"provider_count_email":        {record["provider_count_email"].(string)},
			"provider_count_fax":          {record["provider_count_fax"].(string)},
			"provider_count_phone":        {record["provider_count_phone"].(string)},
			"provider_f_address":          {"Донецкая Народная Республика, городской округ Донецк, г. Донецк, ул. Щорса, д. 87"},
			"provider_f_email":            {"info@Donetsk.e-sbt.ru"},
			"provider_f_www":              {"Donetsk.e-sbt.ru"},
			// "network_company2":            {record["network_company2"].(string)},
			// "network_company2_email":      {record["network_company2_email"].(string)},
			// "network_company2_phone":      {record["network_company2_phone"].(string)},
			// "network_company2_place":      {record["network_company2_place"].(string)},
			// "network_company2_post":       {record["network_company2_post"].(string)},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return makeDocument(body, "0c485c4c-40c8-11ee-8de3-02c163b943db")
}

func main() {
	logName := "./logs/contract_maker_" + time.Now().Format("20060102") + ".log"
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		err := os.Mkdir("./logs", 0700)
		if err != nil {
			logName = "./contract_maker_" + time.Now().Format("20060102") + ".log"
		}
	}

	_log, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	wrt := io.MultiWriter(os.Stdout, _log)
	log.SetOutput(wrt)
	log.SetFlags(log.LstdFlags | log.Ldate | log.Ltime | log.Lmicroseconds)

	log.Println("==========")

	mapDataContracts, err := parse()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range mapDataContracts {
		log.Printf("[INFO] Create contract: %v", record["contract_number"])

		dirName := record["contract_number"].(string)
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}

		filename := fmt.Sprintf("./%s/Договор.pdf", dirName)
		link, format, _ := makeContract(record)
		download(link, format, filename)

		filename = fmt.Sprintf("./%s/Приложение_2.pdf", dirName)
		link, format, _ = makeAttachment2(record)
		download(link, format, filename)

		filename = fmt.Sprintf("./%s/Приложение_1.pdf", dirName)
		link, format, _ = makeAttachment1(record)
		download(link, format, filename)

		for _, counter := range record["counters"].([]map[string]interface{}) {
			filename = fmt.Sprintf("./%s/Приложение_1.1_%v.pdf", dirName, counter["device_number"])
			link, format, _ = makeAttachment11(record, counter["transformers"].([]map[string]interface{}))
			download(link, format, filename)
		}
	}
}
