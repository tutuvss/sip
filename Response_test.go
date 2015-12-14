package sip

import "testing"

func TestReadResponse(t *testing.T) {
	var tvi = []string{
		"SIP/2.0 180 Ringing\r\n" +
			"Via: SIP/2.0/UDP 172.18.1.29:5060;branch=z9hG4bK43fc10fb4446d55fc5c8f969607991f4\r\n" +
			"To: \"0440\" <sip:0440@212.209.220.131>;tag=2600\r\n" +
			"From: \"Andreas\" <sip:andreas@e-horizon.se>;tag=8524\r\n" +
			"Call-ID: f51a1851c5f570606140f14c8eb64fd3@172.18.1.29\r\n" +
			"CSeq: 1 INVITE\r\n" +
			"Max-Forwards: 70\r\n" +
			"Record-Route: <sip:212.209.220.131:5060>\r\n" +
			"Content-Length: 0\r\n\r\n",
	}
	//	var tvo = []string{
	//		"SIP/2.0 180 Ringing\r\n" +
	//			"Via: SIP/2.0/UDP 172.18.1.29:5060;branch=z9hG4bK43fc10fb4446d55fc5c8f969607991f4\r\n" +
	//			"To: \"0440\" <sip:0440@212.209.220.131>;tag=2600\r\n" +
	//			"From: \"Andreas\" <sip:andreas@e-horizon.se>;tag=8524\r\n" +
	//			"Call-ID: f51a1851c5f570606140f14c8eb64fd3@172.18.1.29\r\n" +
	//			"CSeq: 1 INVITE\r\n" +
	//			"Max-Forwards: 70\r\n" +
	//			"Record-Route: <sip:212.209.220.131:5060>\r\n" +
	//			"Content-Length: 0\r\n\r\n",
	//	}

	for i := 0; i < len(tvi); i++ {
		//		smp := NewStringMsgParser()
		//		if sm, err := smp.ParseSIPMessage(tvi[i]); err != nil {
		//			t.Log(err)
		//			t.Fail()
		//		} else {
		//			d := sm.String()
		//			s := tvo[i]

		//			if strings.TrimSpace(d) != strings.TrimSpace(s) {
		//				t.Log("origin = " + s)
		//				t.Log("failed = " + d)

		//				// for j := 0; j < len(s); j++ {
		//				// 	if d[j] != s[j] {
		//				// 		t.Logf("%d:%c vs %c", j, d[j], s[j])
		//				// 	}
		//				// }

		//				t.Fail()
		//			}
		//		}
	}
}