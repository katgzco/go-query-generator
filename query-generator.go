var tables = map[string]string{
	"structName": "db_table_name",
}

/*
*
It builds the query from the structure that is passed to the function from the db tags that are present in the function.
From these tags it takes the columns, the placeholders and the storage values.
This implementation is done to have an abstraction of the storage action performed by the various entities present in the application.
*/

func QueryConstructor(i interface{}) (string, string, string, []interface{}, error) {
	var columns []string
	var values []interface{}
	var placeHolders []string
	var table string
	var lenFieldStruct int
	if i == nil {
		return "", "", "", nil, errors.New("[QueryConstructor] nil interface")
	}
	if reflect.TypeOf(i).Kind() == reflect.Ptr || reflect.TypeOf(i).Kind() != reflect.Struct { //Check if variable is a pointer or nil
		return "", "", "", nil, errors.New("[QueryConstructor] not a struct or is a pointer")
	}
	valueType := reflect.ValueOf(i) // Get the type
	fmt.Println("valueType: ", valueType)
	lenFieldStruct = int(valueType.NumField()) // len of struct fields
	for i := 0; i < lenFieldStruct; i++ {
		tag := valueType.Type().Field(i).Tag.Get(structTag)
		if tag != "" && tag != "-" {
			columns = append(columns, tag)
			fieldValues := valueType.Field(i).Interface() // The interface is invoked to store the primitive datatype and not the reflect type
			values = append(values, fieldValues)
			if i < lenFieldStruct-1 { // last character should be ?
				placeHolders = append(placeHolders, "? ,")
			} else {
				placeHolders = append(placeHolders, "?")
			}
		}
	}
	// Convert []string{values} to `values`
	columnsStr := strings.Join(columns, ", ")
	placeHoldersStr := strings.Join(placeHolders, "")
	typeName := reflect.TypeOf(i).Name() // get the name of the type to get the table
	table = tables[typeName]
	return columnsStr, placeHoldersStr, table, values, nil
}
