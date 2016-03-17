package rtu

import (
	"reflect"
	"strings"
	"log"
	"os"
)

type CDR struct {
	CDR_ID                   string
	CDR_DATE                 string
	RECORD_TYPE              string
	IN_ANI                   string
	IN_DNIS                  string
	OUT_ANI                  string
	OUT_DNIS                 string
	BILL_ANI                 string
	BILL_DNIS                string
	SIG_NODE_NAME            string
	REMOTE_SRC_SIG_ADDRESS   string
	REMOTE_DST_SIG_ADDRESS   string
	REMOTE_SRC_MEDIA_ADDRESS string
	REMOTE_DST_MEDIA_ADDRESS string
	LOCAL_SRC_SIG_ADDRESS    string
	LOCAL_DST_SIG_ADDRESS    string
	LOCAL_SRC_MEDIA_ADDRESS  string
	LOCAL_DST_MEDIA_ADDRESS  string
	IN_LEG_PROTO             string
	OUT_LEG_PROTO            string
	CONF_ID                  string
	IN_LEG_CALL_ID           string
	OUT_LEG_CALL_ID          string
	SRC_IN_LEG_CONF_ID       string
	SRC_IN_LEG_CALL_ID       string
	SRC_OUT_LEG_CALL_ID      string
	SRC_USER                 string
	DST_USER                 string
	SRC_NAME                 string
	DST_NAME                 string
	DP_NAME                  string
	ELAPSED_TIME             string
	SETUP_TIME               string
	CONNECT_TIME             string
	DISCONNECT_TIME          string
	DISCONNECT_CODE          string
	IN_LEG_CODECS            string
	OUT_LEG_CODECS           string
	SRC_FASTSTART_PRESENT    string
	DST_FASTSTART_PRESENT    string
	SRC_TUNNELING_PRESENT    string
	DST_TUNNELING_PRESENT    string
	PROXY_MODE               string
	LAR_FAULT_REASON         string
	ROUTE_RETRIES            string
	SCD                      string
	PDD                      string
	SRC_MEDIA_BYTES_IN       string
	SRC_MEDIA_BYTES_OUT      string
	DST_MEDIA_BYTES_IN       string
	DST_MEDIA_BYTES_OUT      string
	SRC_MEDIA_PACKETS        string
	DST_MEDIA_PACKETS        string
	SRC_MEDIA_PACKETS_LATE   string
	DST_MEDIA_PACKETS_LATE   string
	SRC_MEDIA_PACKETS_LOST   string
	DST_MEDIA_PACKETS_LOST   string
	SRC_MIN_JITTER_SIZE      string
	SRC_MAX_JITTER_SIZE      string
	DST_MIN_JITTER_SIZE      string
	DST_MAX_JITTER_SIZE      string
	LAST_CDR                 string
	IN_CPC                   string
	OUT_CPC                  string
	IN_ZONE                  string
	OUT_ZONE                 string
	DISCONNECT_INITIATOR     string
	DIVERSION                string
	IN_ANI_TYPE_OF_NUMBER    string
	IN_DNIS_TYPE_OF_NUMBER   string
	OUT_ANI_TYPE_OF_NUMBER   string
	OUT_DNIS_TYPE_OF_NUMBER  string
	IN_ORIG_DNIS             string
	OUT_ORIG_DNIS            string
	SRC_DISCONNECT_CODES     string
	DST_DISCONNECT_CODES     string
}

func (c *CDR) SetField(fn string, fv string) {
	s := reflect.ValueOf(c).Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName(strings.ToUpper(fn))
		if f.IsValid() {
			if f.CanSet() {
				if f.Kind() == reflect.String {
					//log.Printf("%v", strings.ToUpper(field))
					if fv != "" {
						f.SetString(fv)
					}
				}
			}
		} else {
			log.Printf("%v not present in CDR struct", strings.ToUpper(fn))
			os.Exit(0)
		}
	}
}
