package rtu

import (
	"reflect"
	"strings"
	"log"
	"os"
)

type rtuInserter interface {
	Insert(*RTUClient) (int, error)
}

type rtuSelecter interface {
	Select(*RTUClient) (*responseRowset, error)
}

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
	SRC_GATEKEEPER_ADDRESS   string
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
	RADIUS_USER              string
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
	Q850_REASON              string
	IN_CPC                   string
	OUT_CPC                  string
	PASS_FROM                string
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

func (c *CDR) Select(*RTUClient) (*responseRowset, error) {
	return &responseRowset{}, nil
}

type Prerouting struct {
	RULE_ID             string
	RULE_NAME           string
	DESCRIPTION         string
	ENABLE              string
	PRIORITY            string
	ANI_PATTERN         string
	DNIS_PATTERN        string
	ORIG_DNIS_PATTERN   string
	ANI_EXCLUDE         string
	DNIS_EXCLUDE        string
	ORIG_DNIS_EXCLUDE   string
	GROUP_ALLOW         string
	GROUP_DENY          string
	ANI_TRANSLATE       string
	ANI_BILL_TRANSLATE  string
	DNIS_TRANSLATE      string
	DNIS_BILL_TRANSLATE string
	ORIG_DNIS_TRANSLATE string
	ANI_SORM_TRANSLATE  string
	DNIS_SORM_TRANSLATE string
	GROUP_TRANSLATE     string
	CPC_ALLOW           string
	CPC_DENY            string
	CPC_TRANSLATE       string
	ACTION              string
	DISCONNECT_CODE     string
	SCHED_TYPE          string
	SCHED_TOD           string
	SCHED_TOD_MON       string
	SCHED_TOD_TUE       string
	SCHED_TOD_WED       string
	SCHED_TOD_THU       string
	SCHED_TOD_FRI       string
	SCHED_TOD_SAT       string
	SCHED_TOD_SUN       string
	SCHED_TOM           string
	SCHED_TOY           string
}

func (p *Prerouting) Insert(rc *RTUClient) (count int, err error) {
	rm, err := rowStructToMap(p)
	rsm := []map[string]string{
		0: rm,
	}
	if err != nil {
		return 0, err
	}
	count, err = rc.Query().Insert().Into(TablePrerouting).Values(rsm).GetInt()
	return count, err
}

func (p *Prerouting) Select(*RTUClient) (*responseRowset, error) {
	return &responseRowset{}, nil
}

//Interface functions

func fillStruct(s rtuSelecter, r *responseRow) {
	for _, item := range r.Items {
		setStructField(s, item.Key, item.Value)
	}
	return
}

func setStructField(s rtuSelecter, fn string, fv string) {
	v := reflect.ValueOf(s).Elem()
	if v.Kind() == reflect.Struct {
		f := v.FieldByName(strings.ToUpper(fn))
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
	return
}

func rowStructToMap(s rtuInserter) (map[string]string, error) {
	out := make(map[string]string)
	v := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		value := v.Field(i).String()
		if value != "" {
			out[key] = value
		}
	}
	return out, nil
}
