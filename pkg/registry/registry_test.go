package registry

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockClient struct {
	basicHTTPClient
}

func (c *mockClient) AuthHeader(image string) (string, error) {
	return "TOKEN", nil
}

func TestGetTags(t *testing.T) {
	ts := mockServer()
	defer ts.Close()

	client := ts.Client()
	c := mockClient{basicHTTPClient{client: client, baseURL: ts.URL}}

	tests := []struct {
		image        string
		expectedTags int
		notFound     bool
	}{
		{image: "alpine", expectedTags: 12},
		{image: "alpine", expectedTags: 12},
		{image: "nginx", notFound: true},
		{image: "gcr.io/google-containers/kube-apiserver", expectedTags: 302},
		{image: "quay.io/coreos/etcd", expectedTags: 26},
	}

	for _, test := range tests {
		tags, err := getTags(test.image, &c)
		if test.notFound {
			if err == nil || !strings.Contains(err.Error(), "got a bad return code") {
				t.Errorf("expected an error to contain 'got a bad return code' instead got: %v", err)
			}
			continue
		}
		if err != nil {
			t.Errorf("error getting tags for %q: %v", test.image, err)
		}
		if len(tags) != test.expectedTags {
			t.Errorf("unexpected number of tags returned for  %q %d", test.image, len(tags))
		}
	}
}

func mockServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("/v2/%s/tags/", "alpine"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, dockerhubAlpineResp)
	})
	mux.HandleFunc(fmt.Sprintf("/v2/%s/tags/", "gcr.io/google-containers/kube-apiserver"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, gcrKubeApiserverResp)
	})
	mux.HandleFunc(fmt.Sprintf("/v2/%s/tags/", "quay.io/coreos/etcd"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, quayEtcdResp)
	})
	return httptest.NewTLSServer(mux)
}

var dockerhubAlpineResp = `{
    "name": "library/alpine",
    "tags": [
        "2.6",
        "2.7",
        "3.1",
        "3.2",
        "3.3",
        "3.4",
        "3.5",
        "3.6",
        "3.7",
        "3.8",
        "edge",
        "latest"
    ]
}`

var gcrKubeApiserverResp = `{
    "child": [],
    "manifest": {
        "sha256:009083f7d4569e56826bf10a4cca87a51228c90c59dfafea6955aeedef2df266": {
            "imageSizeBytes": "27735631",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.3"
            ],
            "timeCreatedMs": "1476600931051",
            "timeUploadedMs": "1476605657097"
        },
        "sha256:00ac3585f4e3a41bda30286be1d3b812b2d9bbb6193682a7b5f46dbdadfe41da": {
            "imageSizeBytes": "31142483",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.1-beta.0"
            ],
            "timeCreatedMs": "1513375496673",
            "timeUploadedMs": "1513378485105"
        },
        "sha256:01000783487d958867f7de83b5ed915e1abf98caafd5d85e12167d588070cca5": {
            "imageSizeBytes": "13039845",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.5"
            ],
            "timeCreatedMs": "1467051192707",
            "timeUploadedMs": "1467051704930"
        },
        "sha256:01b8d1436f32c78e19aff309055080d585897f4f5b858583c782975fc3c60ee1": {
            "imageSizeBytes": "35019702",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0"
            ],
            "timeCreatedMs": "1530132957220",
            "timeUploadedMs": "1530135858556"
        },
        "sha256:01cb88579277ba9731885e127681f1ad97d99705914218aa54f7b1c5e5185ff6": {
            "imageSizeBytes": "31736483",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.7-beta.0"
            ],
            "timeCreatedMs": "1505381097921",
            "timeUploadedMs": "1505388241239"
        },
        "sha256:03e061706294225108829226b3b047bbd48e3c5e029d3d67e47de176e10f9a00": {
            "imageSizeBytes": "27682882",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.5.0-alpha.0"
            ],
            "timeCreatedMs": "1472763974729",
            "timeUploadedMs": "1472766246216"
        },
        "sha256:04291a8c4954b6f0e6f40dcd3ce61ae2d0cb466c768504667b36b49cc184fa6e": {
            "imageSizeBytes": "30210759",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.12-beta.0"
            ],
            "timeCreatedMs": "1522951426431",
            "timeUploadedMs": "1522960032973"
        },
        "sha256:05583286993b43d58b5e9736fbd499588bf58a83ee3adb67009afbe5ae5670ad": {
            "imageSizeBytes": "20569997",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-beta.3"
            ],
            "timeCreatedMs": "1467333861598",
            "timeUploadedMs": "1467334574815"
        },
        "sha256:058f9327867fb8a61df2de5f98cd690ef9ef8824cf341851bee68c459af84203": {
            "imageSizeBytes": "30210395",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.13-beta.0"
            ],
            "timeCreatedMs": "1524527870958",
            "timeUploadedMs": "1524544013219"
        },
        "sha256:071c9dfc9e2377331733a5b19f0db75a6551395008195b662342c2d65855d33c": {
            "imageSizeBytes": "20611673",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.8"
            ],
            "timeCreatedMs": "1477509123094",
            "timeUploadedMs": "1477510087514"
        },
        "sha256:07599038d25daaffc08a1edf6983fa9abb01ac8307b7efe2c88e81515d562a2a": {
            "imageSizeBytes": "27616028",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.11"
            ],
            "timeCreatedMs": "1506634932095",
            "timeUploadedMs": "1506641050211"
        },
        "sha256:07b76a175296b89d3bc363003518de1704f87926476d938d1b7eab682580cc71": {
            "imageSizeBytes": "22809632",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.3-beta.0"
            ],
            "timeCreatedMs": "1484201682378",
            "timeUploadedMs": "1484205682891"
        },
        "sha256:07dc7d7afbe8e746a6c94ba30825fe73fdb022f7d497005e31efa35d9b0636cf": {
            "imageSizeBytes": "27570991",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.3-beta.1"
            ],
            "timeCreatedMs": "1494887842944",
            "timeUploadedMs": "1494893965959"
        },
        "sha256:07efc9d98b29ffe0bc174cfe9acccf2ea4e9259a8ea40cb4d60805cfb9a0494d": {
            "imageSizeBytes": "30210903",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.11"
            ],
            "timeCreatedMs": "1522950660745",
            "timeUploadedMs": "1522959466679"
        },
        "sha256:08d145d28eb681bff15f4a1ba081f940d150617273abd0038f28d7c179034571": {
            "imageSizeBytes": "29778283",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.2-beta.0"
            ],
            "timeCreatedMs": "1507767133952",
            "timeUploadedMs": "1507769271845"
        },
        "sha256:08e3ef73efc50526973cb9e1282cdb6c6be0cb1b7ae46cc156019e940b385e00": {
            "imageSizeBytes": "31493549",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.1-beta.0"
            ],
            "timeCreatedMs": "1498784833633",
            "timeUploadedMs": "1498799741252"
        },
        "sha256:0a3aa91344a3da4540a7ee0d006cae9c3380c4b0e907f687cac495772299bf14": {
            "imageSizeBytes": "29792945",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.5-beta.0"
            ],
            "timeCreatedMs": "1511158352471",
            "timeUploadedMs": "1511203950692"
        },
        "sha256:0d63e91a80b84d5323ebf5852e9ebbda311300a59f4a7a0e5644cfccb9af7385": {
            "imageSizeBytes": "31671154",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.6-beta.0"
            ],
            "timeCreatedMs": "1521477828861",
            "timeUploadedMs": "1521484169796"
        },
        "sha256:0d7870b099909355292d08fc5a1f60e2dd12db056f54e15942c99d64cbe3477d": {
            "imageSizeBytes": "28056893",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-beta.2"
            ],
            "timeCreatedMs": "1489005324832",
            "timeUploadedMs": "1489007823948"
        },
        "sha256:0dcad4c5c58929a91cefac77f4a278bd6edb77347d0c8bab9cbfdc67a8850d90": {
            "imageSizeBytes": "27616349",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.13-beta.0"
            ],
            "timeCreatedMs": "1508958171341",
            "timeUploadedMs": "1508961815771"
        },
        "sha256:0fbe8134aefc3e8c0ce54cf03ef77e28eda22d9afbb1193c14d3e1bbe7e9d97f": {
            "imageSizeBytes": "22809596",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.2"
            ],
            "timeCreatedMs": "1484197681519",
            "timeUploadedMs": "1484205059856"
        },
        "sha256:103f5938cd686096075f50f045eac58e4a6dc57b949ab8a6382750b99aac3f17": {
            "imageSizeBytes": "32941562",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.5-beta.0"
            ],
            "timeCreatedMs": "1528276876112",
            "timeUploadedMs": "1528283287845"
        },
        "sha256:10d693e7180580e2bbc8d22eebb8fb2def88b69df184da8938f3984f9018bb6e": {
            "imageSizeBytes": "27613594",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.9-beta.0"
            ],
            "timeCreatedMs": "1501787894169",
            "timeUploadedMs": "1501791105710"
        },
        "sha256:12b554e562fb1677619df2b31f66371086c182d7ba936cd345bb9ef868d01205": {
            "imageSizeBytes": "27740211",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.4"
            ],
            "timeCreatedMs": "1477018521049",
            "timeUploadedMs": "1477020158889"
        },
        "sha256:15c8bd68be1eb9556f3512c62714d34cc97d73ef1d00046715709d56e0ba5f59": {
            "imageSizeBytes": "13038934",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.5-beta.0"
            ],
            "timeCreatedMs": "1462557080464",
            "timeUploadedMs": "1462560526613"
        },
        "sha256:18feec6fc5862367bd5e0d31ec551bfffae84ee00e955debdcb62e7926bb319d": {
            "imageSizeBytes": "33379706",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.7-beta.0"
            ],
            "timeCreatedMs": "1532604896882",
            "timeUploadedMs": "1532609998991"
        },
        "sha256:19274315cd5e6dd61d8476994117ca3b312419346633a190ce29726c326a7960": {
            "imageSizeBytes": "19638890",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-alpha.4"
            ],
            "timeCreatedMs": "1463522667982",
            "timeUploadedMs": "1463523436341"
        },
        "sha256:1969f8075587638fe213940d79a4175abfd45136a304ef6e83b737f6573758f4": {
            "imageSizeBytes": "27752489",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.7"
            ],
            "timeCreatedMs": "1474149585877",
            "timeUploadedMs": "1474150497470"
        },
        "sha256:197ac40e9c3cea013f3b2adfd6a2c8c54168b898cc618a19db997cfaeef3880d": {
            "imageSizeBytes": "27559139",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.3-beta.0"
            ],
            "timeCreatedMs": "1492637738327",
            "timeUploadedMs": "1492642276025"
        },
        "sha256:1988ec204ca3c5428d4b1601e0b0c2168e6d2b71ce127157bec44f814a6e8c33": {
            "imageSizeBytes": "31718655",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.3-beta.0"
            ],
            "timeCreatedMs": "1500633485296",
            "timeUploadedMs": "1500639511658"
        },
        "sha256:198cd724ccf16a7c12b17c1b6f4898eabefa7265aac2016941e449f5734582b2": {
            "imageSizeBytes": "26795634",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-alpha.0"
            ],
            "timeCreatedMs": "1487733969546",
            "timeUploadedMs": "1487740942440"
        },
        "sha256:1ab992df89cc24f92ed99ec96636d8d214a1c06579c0e723640d5a5f752650d2": {
            "imageSizeBytes": "20557503",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.4"
            ],
            "timeCreatedMs": "1470070260751",
            "timeUploadedMs": "1470071914114"
        },
        "sha256:1ba863c8e9b9edc6d1329ebf966e4aa308ca31b42a937b4430caf65aa11bdd12": {
            "imageSizeBytes": "32912344",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.2"
            ],
            "timeCreatedMs": "1524824094199",
            "timeUploadedMs": "1524828670365"
        },
        "sha256:1bb16ddef9edd8142125ac2443e36986c051b0df808db907648ac9d68e0d5f23": {
            "imageSizeBytes": "31672140",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.6"
            ],
            "timeCreatedMs": "1521647500996",
            "timeUploadedMs": "1521656330699"
        },
        "sha256:1dddbf3fe3140575f7af35c01b36c57adbdf6e4090bde623d17203794f19948e": {
            "imageSizeBytes": "27735638",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.2"
            ],
            "timeCreatedMs": "1476490814787",
            "timeUploadedMs": "1476492496766"
        },
        "sha256:1e256ad9508a48d31f864baec538bd27b1c88bbae1f4448e024e7d9c863a5f43": {
            "imageSizeBytes": "20565595",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.3"
            ],
            "timeCreatedMs": "1469219754711",
            "timeUploadedMs": "1469223974163"
        },
        "sha256:1e936fda1642311dfcbafde363a07b8abba86a55e1a2caff5aee8fcfb56c3f75": {
            "imageSizeBytes": "29771564",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-alpha.1"
            ],
            "timeCreatedMs": "1506131497572",
            "timeUploadedMs": "1506133334154"
        },
        "sha256:1ec736c48ac0e351df2ca1a944d0c8a435ba5b36e2b3cdb72e59eabf7bd2fc6d": {
            "imageSizeBytes": "30561345",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-alpha.0"
            ],
            "timeCreatedMs": "1496374897484",
            "timeUploadedMs": "1496387503513"
        },
        "sha256:20644a3434b682fb8c0e4a6084ce8135995e49cce5d70cace5e8bd0cd3a584aa": {
            "imageSizeBytes": "31497540",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.2-beta.0"
            ],
            "timeCreatedMs": "1500002084858",
            "timeUploadedMs": "1500005689925"
        },
        "sha256:20fa86aa7d210d948946785ad6875f07c01cb62e522660bbb6ccd7fe73280b9b": {
            "imageSizeBytes": "30357004",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.15"
            ],
            "timeCreatedMs": "1531333636554",
            "timeUploadedMs": "1531349493613"
        },
        "sha256:210241064ac18e9940b92c25868be9521298cc8414cb7653d48adfefd208db8b": {
            "imageSizeBytes": "27598398",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.5"
            ],
            "timeCreatedMs": "1497471972479",
            "timeUploadedMs": "1497477990279"
        },
        "sha256:212dba7a7a4dae2a2a735b5aadd234afe569d266f17fa2550f25b825f8fe5ed8": {
            "imageSizeBytes": "20611760",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.10-beta.0"
            ],
            "timeCreatedMs": "1477001063561",
            "timeUploadedMs": "1477432119539"
        },
        "sha256:240f5a77f97fa92a556a532dccf46d635d0de802894ea62a3cf1a60e1d95dd57": {
            "imageSizeBytes": "13209867",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-alpha.1"
            ],
            "timeCreatedMs": "1459289696293",
            "timeUploadedMs": "1459289940405"
        },
        "sha256:24609093d42666f408465e0f74ddc913c4d475edb5a78bd7876e3b387916c324": {
            "imageSizeBytes": "29110187",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-alpha.2"
            ],
            "timeCreatedMs": "1492647866747",
            "timeUploadedMs": "1492704365229"
        },
        "sha256:24841cbae54a535101ca84fb60a84f4e9b904c9ed376accade4b440654c18821": {
            "imageSizeBytes": "31546902",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.4-beta.0"
            ],
            "timeCreatedMs": "1518014618543",
            "timeUploadedMs": "1518212336650"
        },
        "sha256:2567450273a44bdbb2e9ea375c107e543af1e3cbcd1f39fa702a23ea98910d89": {
            "imageSizeBytes": "33942446",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-alpha.1"
            ],
            "timeCreatedMs": "1524167638270",
            "timeUploadedMs": "1524169438556"
        },
        "sha256:25cb990da033028596b53ac641f5e6ec555e9e1bef998ba3934ef10d7c6140bc": {
            "imageSizeBytes": "32168502",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.14"
            ],
            "timeCreatedMs": "1520873267305",
            "timeUploadedMs": "1520880346438"
        },
        "sha256:296b7febb48f67aeda984bf0b7b9cb1e8d0d0d7d4389e3f7e9d119b696c025d2": {
            "imageSizeBytes": "27684510",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.0"
            ],
            "timeCreatedMs": "1472765606107",
            "timeUploadedMs": "1472766747900"
        },
        "sha256:2aa256ff65e607c6bc9383208c9e6940a1feab5b5a25b83738ff943ed4ddc689": {
            "imageSizeBytes": "23801517",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.9-beta.0"
            ],
            "timeCreatedMs": "1506822642638",
            "timeUploadedMs": "1506841835339"
        },
        "sha256:2b1881caaf447095e63fbc3bb26b47f74ea2bd4e8f644176b8df63a1bcedfaae": {
            "imageSizeBytes": "20570013",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0"
            ],
            "timeCreatedMs": "1467401579137",
            "timeUploadedMs": "1467403669600"
        },
        "sha256:2c27af79f314a9a030b4f3d1787bcd1b931779bcc3d229c7fc3fa83df0bbd7b1": {
            "imageSizeBytes": "22811827",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.4"
            ],
            "timeCreatedMs": "1488932575806",
            "timeUploadedMs": "1488938943619"
        },
        "sha256:2cf2826f8aebf4a9fb7fb3233524ee01a7bd5248e058de2be258c9ddfe640b07": {
            "imageSizeBytes": "31719506",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.4-beta.0"
            ],
            "timeCreatedMs": "1501751874672",
            "timeUploadedMs": "1501758704719"
        },
        "sha256:2d26b84ecad8f7a2189a88d4c70c6f184b6cf89698c7aef2f7686b9e628bb81b": {
            "imageSizeBytes": "32514895",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-alpha.0"
            ],
            "timeCreatedMs": "1518632003688",
            "timeUploadedMs": "1518642918053"
        },
        "sha256:2d350fdea685f6e2eb8d68a46db88347ee3c6bab314febe67d41152547f7583f": {
            "imageSizeBytes": "27572896",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.4-beta.0"
            ],
            "timeCreatedMs": "1494435185952",
            "timeUploadedMs": "1494440313874"
        },
        "sha256:2d8075baa50796ec18302b7b7108093fc9ab9499f192bf93348b149b0d972e0e": {
            "imageSizeBytes": "20412472",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-alpha.0"
            ],
            "timeCreatedMs": "1465594373510",
            "timeUploadedMs": "1465595918445"
        },
        "sha256:2ef812f77eaa0cd46aba04c51315c609412fb0b72e07c6e91509a741d9eab543": {
            "imageSizeBytes": "32146266",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.14-beta.0"
            ],
            "timeCreatedMs": "1519655169780",
            "timeUploadedMs": "1519895861183"
        },
        "sha256:3024718df65a431b6f1a540fb3eefc306f3066446428518999d9f299a5263944": {
            "imageSizeBytes": "22357858",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.5.0-beta.0"
            ],
            "timeCreatedMs": "1478398077573",
            "timeUploadedMs": "1478401806147"
        },
        "sha256:30ace5133327838037dac781dee4691a12bf57e4111dc9974fd570aaaae45ab5": {
            "imageSizeBytes": "20611664",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [],
            "timeCreatedMs": "1477444811848",
            "timeUploadedMs": "1477506331892"
        },
        "sha256:31eac037d316e2002bee0266a7f5cae7e057251d03a95852a62b961c65834533": {
            "imageSizeBytes": "22304531",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.5.0-alpha.2"
            ],
            "timeCreatedMs": "1477606960092",
            "timeUploadedMs": "1477610191158"
        },
        "sha256:3283c534aab386bc678986c74c2fc855db46b3f424884d547641c82dba6ec5de": {
            "imageSizeBytes": "33379865",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.6"
            ],
            "timeCreatedMs": "1532603987043",
            "timeUploadedMs": "1532609479155"
        },
        "sha256:32901915205c4c3be4960638f8f6440b0c924d8d925ad18694db8e4f00036ed0": {
            "imageSizeBytes": "13038914",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.4"
            ],
            "timeCreatedMs": "1462556777202",
            "timeUploadedMs": "1462559992729"
        },
        "sha256:329f116e68364e241e35c9c4ca9a4431f8b9bf2b4e098e3af446f8415548b0ad": {
            "imageSizeBytes": "31825452",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.9"
            ],
            "timeCreatedMs": "1530236585695",
            "timeUploadedMs": "1530251560346"
        },
        "sha256:34e399e9f219bc2984df254d32b91001fcc58a14a6992c0c04710ac0951c13e9": {
            "imageSizeBytes": "30209168",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.10"
            ],
            "timeCreatedMs": "1521483480145",
            "timeUploadedMs": "1521488554495"
        },
        "sha256:3777a5c0e36dd8e96d8ea8f54b99c2278139c92708263652ed0aa19ea284cb28": {
            "imageSizeBytes": "35675519",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.12.0-alpha.1"
            ],
            "timeCreatedMs": "1533167315148",
            "timeUploadedMs": "1533169297301"
        },
        "sha256:3a32cfd99aa07ca00fc32d815c8812c4cebfd167bf4aacb4cf25e2299fa7a4ae": {
            "imageSizeBytes": "20554770",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-beta.2"
            ],
            "timeCreatedMs": "1466541851688",
            "timeUploadedMs": "1466542508545"
        },
        "sha256:3b427d8178f3064a931b8eb3851e87cd920b63c0bd652d5d807991a7b03372e6": {
            "imageSizeBytes": "31737660",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.7"
            ],
            "timeCreatedMs": "1506645377414",
            "timeUploadedMs": "1506661869245"
        },
        "sha256:3d2dcf9386221b26ff4a2e57462361ecd2417ed7c91f3f77a26d582170577561": {
            "imageSizeBytes": "27709700",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.8-beta.0"
            ],
            "timeCreatedMs": "1481346968174",
            "timeUploadedMs": "1481348030036"
        },
        "sha256:3de9e6d3d1161a0ea5c0d97f3c4753896402970f26c3522d752fb2a9544a7d0e": {
            "imageSizeBytes": "30211389",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.13"
            ],
            "timeCreatedMs": "1526409842148",
            "timeUploadedMs": "1526413021425"
        },
        "sha256:3e883b87bb49dca2acdf0cf8d3f05c1fe315c249c4fe52ceba9172cfc40f8ea3": {
            "imageSizeBytes": "31745520",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.9"
            ],
            "timeCreatedMs": "1508433725809",
            "timeUploadedMs": "1508441389174"
        },
        "sha256:3e980f4b57292568ea8c87be462cf0583e40bbc2dbfff71d0d9e19beda3cb74b": {
            "imageSizeBytes": "29781670",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.2"
            ],
            "timeCreatedMs": "1508875188235",
            "timeUploadedMs": "1508878510148"
        },
        "sha256:407ca1ffe4e198542f5714f6333e97c2158aabec59746b7479560f5bac8bbdde": {
            "imageSizeBytes": "27734750",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.5"
            ],
            "timeCreatedMs": "1473917430899",
            "timeUploadedMs": "1473922229650"
        },
        "sha256:40fece6b61de6e30bc052238d10b000a25acac1b86e7cd3452551f00a7290149": {
            "imageSizeBytes": "27531826",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-beta.4"
            ],
            "timeCreatedMs": "1489773847445",
            "timeUploadedMs": "1489776278398"
        },
        "sha256:41628d9f4494a5b1a9d46d5b989a635b14eaa24ecbd6836da6646e8ce53bc083": {
            "imageSizeBytes": "27614092",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.9"
            ],
            "timeCreatedMs": "1503508833331",
            "timeUploadedMs": "1503517039063"
        },
        "sha256:427a1a39d63da9c5a010daff3025cdb4b0ef8a1676950391f083a404794d131d": {
            "imageSizeBytes": "31492106",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-rc.1"
            ],
            "timeCreatedMs": "1498284486415",
            "timeUploadedMs": "1498342677607"
        },
        "sha256:42c6118449fa23c20e70051a57fe0fee8f9ff4d241607e8249ffe602fcb8ae2c": {
            "imageSizeBytes": "20570009",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.1-beta.0"
            ],
            "timeCreatedMs": "1467402698688",
            "timeUploadedMs": "1467404329540"
        },
        "sha256:4462b7748893fe0aa53d25c2091ec13c8341242baa7b8c5a6a13c3199d08dabf": {
            "imageSizeBytes": "31129739",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-beta.1"
            ],
            "timeCreatedMs": "1512072507528",
            "timeUploadedMs": "1512076714993"
        },
        "sha256:44be264af2d412ee69ddc000979a9487a3cd26da3cff950fd2716d5d941d05ea": {
            "imageSizeBytes": "22671324",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.5.0-beta.2"
            ],
            "timeCreatedMs": "1480027438636",
            "timeUploadedMs": "1480036819778"
        },
        "sha256:45863fe3af3cb99f38afeed84772e5a37650d4dd7a455f6bf60643df582a3d67": {
            "imageSizeBytes": "30196920",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.8"
            ],
            "timeCreatedMs": "1518213585807",
            "timeUploadedMs": "1518217463482"
        },
        "sha256:45a46bee201b184335ed893909335ba208fca676c9ecee36acf6db0deafee893": {
            "imageSizeBytes": "24943327",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-alpha.1"
            ],
            "timeCreatedMs": "1485831394395",
            "timeUploadedMs": "1485834459376"
        },
        "sha256:462ae77a32db30a08298bd1e365ef29d8d06788096653bee7ac40b10b0087130": {
            "imageSizeBytes": "35032944",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.1"
            ],
            "timeCreatedMs": "1531856092690",
            "timeUploadedMs": "1531933909732"
        },
        "sha256:47377f5df3ea25cb677fe081be78dce9ffe4a58dcaaa4191a0a9e8295cfb848b": {
            "imageSizeBytes": "31142413",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0"
            ],
            "timeCreatedMs": "1513372828436",
            "timeUploadedMs": "1513377547310"
        },
        "sha256:48ede5f3211fed630cf1ada49ff844779d7c77c0c7cd8dbe3d1f3495faa7f40e": {
            "imageSizeBytes": "20611747",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.10"
            ],
            "timeCreatedMs": "1477949609640",
            "timeUploadedMs": "1477952557877"
        },
        "sha256:49d456b1fe6f6b5d69f0ecbfa440cc5e230847d34defd9e2605f2c9b8a8061fa": {
            "imageSizeBytes": "20616545",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-alpha.1"
            ],
            "timeCreatedMs": "1468268440950",
            "timeUploadedMs": "1468271079644"
        },
        "sha256:4a84f8917f96b838160da1e9403518e33563f2f71db64879b88d1068df1e51b4": {
            "imageSizeBytes": "27733235",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.3"
            ],
            "timeCreatedMs": "1473704319978",
            "timeUploadedMs": "1473705352091"
        },
        "sha256:4cdeeec61a0b35563ba152b4df7c8ce556508a7950d56e14c68482bc21a66519": {
            "imageSizeBytes": "32165441",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.16"
            ],
            "timeCreatedMs": "1522833547990",
            "timeUploadedMs": "1522847440526"
        },
        "sha256:4d65f12f09b49838a80f45f20f8aa31a609b828bb5371abc0a2beda717ca784e": {
            "imageSizeBytes": "20568388",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.6"
            ],
            "timeCreatedMs": "1472235579904",
            "timeUploadedMs": "1472237292564"
        },
        "sha256:4f59b5ff9c4c7c0a7d836592dce95e4be9f6d91af0da291d5fbc1ec1c64466ff": {
            "imageSizeBytes": "20608891",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.8-beta.0"
            ],
            "timeCreatedMs": "1473723516990",
            "timeUploadedMs": "1473724712997"
        },
        "sha256:4f63662eab70db8f9c2cbb6185cc2c76063b4be59d2056951260d4ffe101bdfa": {
            "imageSizeBytes": "30356790",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.16-beta.0"
            ],
            "timeCreatedMs": "1531334446971",
            "timeUploadedMs": "1531350021096"
        },
        "sha256:5002411957b90e0ff4f35194f6be01258b7e92055876142b87cc8d29a38488b4": {
            "imageSizeBytes": "30603921",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-alpha.0"
            ],
            "timeCreatedMs": "1510872622989",
            "timeUploadedMs": "1510880980707"
        },
        "sha256:5161dda759c3c4b35d1eccbb89a9822085abe39a5669329ac09234b328bb8c83": {
            "imageSizeBytes": "20409135",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-beta.1"
            ],
            "timeCreatedMs": "1466125486293",
            "timeUploadedMs": "1466126706463"
        },
        "sha256:517931adb87128ba75ad96c177eb08b64e919561a1d24c5b316666b4d745ab3f": {
            "imageSizeBytes": "31825290",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.10-beta.0"
            ],
            "timeCreatedMs": "1530237412059",
            "timeUploadedMs": "1530252121541"
        },
        "sha256:527dd368b116b6e08d00e94aa5fab7f8fa22be7ab339e67aa4cdc64f70373b2b": {
            "imageSizeBytes": "31672363",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.7-beta.0"
            ],
            "timeCreatedMs": "1521648322942",
            "timeUploadedMs": "1521657145942"
        },
        "sha256:53b987e5a2932bdaff88497081b488e3b56af5b6a14891895b08703129477d85": {
            "imageSizeBytes": "31493847",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0"
            ],
            "timeCreatedMs": "1498779153329",
            "timeUploadedMs": "1498798341468"
        },
        "sha256:55d5084015f2b4cf43c3d0290dfe2a26706643b1dee4a1ae94e73db1ab9089f6": {
            "imageSizeBytes": "31839984",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.10"
            ],
            "timeCreatedMs": "1533232156805",
            "timeUploadedMs": "1533259108978"
        },
        "sha256:56571b3108018bd3d99824d46de2ea54b83053155c7d7d9c779f4a68f5bebc26": {
            "imageSizeBytes": "30561352",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-beta.0"
            ],
            "timeCreatedMs": "1496382264895",
            "timeUploadedMs": "1496388636699"
        },
        "sha256:56ac2a86a613a9b56a8d3363e3fd2ddf003d1b9c935bc9a9d1eeecf97b55832a": {
            "imageSizeBytes": "31417237",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-beta.1"
            ],
            "timeCreatedMs": "1496887795100",
            "timeUploadedMs": "1496891743137"
        },
        "sha256:57e36016d973776fedbeaeaedede45a97dbe31f64bc3624740839bb57d741f4e": {
            "imageSizeBytes": "31743164",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.9-beta.0"
            ],
            "timeCreatedMs": "1507196387911",
            "timeUploadedMs": "1507203964219"
        },
        "sha256:57e482529b95d32730d1bcd2e374199f27eab4abcf1ff49c5db2a7a7e2231cc8": {
            "imageSizeBytes": "27598104",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.7"
            ],
            "timeCreatedMs": "1499274065771",
            "timeUploadedMs": "1499279817299"
        },
        "sha256:582c2a5e1361ba76fbe8f92f6ca5fbf14aded7a81d03d55f11d0aa8ff93f4ba3": {
            "imageSizeBytes": "31719310",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.3"
            ],
            "timeCreatedMs": "1501744965932",
            "timeUploadedMs": "1501757223589"
        },
        "sha256:584b6503e8f59cac6c2730b849ad2e68f7a77d7354054a4f11ff3b1756dd72f9": {
            "imageSizeBytes": "32903627",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.2-beta.0"
            ],
            "timeCreatedMs": "1523547022797",
            "timeUploadedMs": "1523556209704"
        },
        "sha256:58f02d5be55482f214366aa4940faa08520cddbc7c3f83ca05b0643fcc312199": {
            "imageSizeBytes": "27759997",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0"
            ],
            "timeCreatedMs": "1474914265726",
            "timeUploadedMs": "1474915876517"
        },
        "sha256:593d401aea8f09258f492b57b65fe66cfeba47d49328405fdd575a0d1a278f00": {
            "imageSizeBytes": "27735126",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.2-beta.1"
            ],
            "timeCreatedMs": "1476374889262",
            "timeUploadedMs": "1476375963654"
        },
        "sha256:5a1623ec941d6fcde6922dd501b8a03f7dd9ffa0f612d00dae7920321ef3681e": {
            "imageSizeBytes": "27572939",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.3"
            ],
            "timeCreatedMs": "1494431898864",
            "timeUploadedMs": "1494439552663"
        },
        "sha256:5a1954f8c54bf56f863f9fd600c97f35362029e26d963f610b7a70b63f98802b": {
            "imageSizeBytes": "27735657",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.4-beta.0"
            ],
            "timeCreatedMs": "1476603879706",
            "timeUploadedMs": "1476606477743"
        },
        "sha256:5a1a7b81b3a6212bea7485e5d0c4f0c97eb240191e84151d0a6c2b17ad20aea5": {
            "imageSizeBytes": "32822709",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-beta.1"
            ],
            "timeCreatedMs": "1519890051614",
            "timeUploadedMs": "1519891617916"
        },
        "sha256:5a3769083f1b1a8f3cf34695f07a054f36f3f8957d6ff56d3f3f9ac40a504db6": {
            "imageSizeBytes": "27614119",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.10-beta.0"
            ],
            "timeCreatedMs": "1503512478417",
            "timeUploadedMs": "1503517847553"
        },
        "sha256:5a85678740eace54c75e4aaaca9dce212518486101b965c362c3278cf4ba3597": {
            "imageSizeBytes": "31547626",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.3-beta.0"
            ],
            "timeCreatedMs": "1516277732317",
            "timeUploadedMs": "1516305401569"
        },
        "sha256:5b287e9241d718af87907b96bd472778b0500ca1cc78d5b157043d1f1081e177": {
            "imageSizeBytes": "30604046",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-beta.0"
            ],
            "timeCreatedMs": "1510875225461",
            "timeUploadedMs": "1510881956456"
        },
        "sha256:5beec729da2ab7a60fc860a31488cee1879616b22c3b1f52e8c306f80fb7b3d1": {
            "imageSizeBytes": "34138995",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-beta.0"
            ],
            "timeCreatedMs": "1526529421731",
            "timeUploadedMs": "1526531801216"
        },
        "sha256:5caa4c6724570201cb4b9bf75172a39ff5aa94189ff5607ee1de5c52b0184b17": {
            "imageSizeBytes": "32941867",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.4"
            ],
            "timeCreatedMs": "1528275983249",
            "timeUploadedMs": "1528282776328"
        },
        "sha256:5e2bbbd55886c799affc9c6f7fc91b2bbb35c4ac4f26949d0f20cc17c32b2c65": {
            "imageSizeBytes": "32910465",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.1-beta.0"
            ],
            "timeCreatedMs": "1522087331499",
            "timeUploadedMs": "1522106058550"
        },
        "sha256:5ffb40ac1f547fb152b28d0c6317e14224335f5197ea7be99b45cf495f61b452": {
            "imageSizeBytes": "30211344",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.14-beta.0"
            ],
            "timeCreatedMs": "1526410676015",
            "timeUploadedMs": "1526413518151"
        },
        "sha256:60fad63f7608fcf7d1c2df5e01884f07f37e82526045e336b7d9c1abec0f5bc4": {
            "imageSizeBytes": "29773105",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [],
            "timeCreatedMs": "1506639293823",
            "timeUploadedMs": "1506648670129"
        },
        "sha256:61cdfa3bf94fb845c54625c1bedef4aa87d61378724ddf2316c14e8ee2ca98e8": {
            "imageSizeBytes": "32146230",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.13"
            ],
            "timeCreatedMs": "1519652664493",
            "timeUploadedMs": "1519895368706"
        },
        "sha256:6273c7c6146212da7c2eb8e5856473f9cc054413d3ffc3a04d0377f81b6d6200": {
            "imageSizeBytes": "30583183",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-alpha.3"
            ],
            "timeCreatedMs": "1510787925521",
            "timeUploadedMs": "1510792098348"
        },
        "sha256:6425c124a58bb1db4c69a1c9019bbd0d5e74587ad83a785e3bc586dc1f8d7ce9": {
            "imageSizeBytes": "31484502",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-beta.2"
            ],
            "timeCreatedMs": "1497547937251",
            "timeUploadedMs": "1497552110101"
        },
        "sha256:646f9c24fc1b9b99d80b9a3cc6d14ad4f80455642eca0c0e88ff2cb9e1b9d1fa": {
            "imageSizeBytes": "22676475",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.5.0-beta.1"
            ],
            "timeCreatedMs": "1479507279200",
            "timeUploadedMs": "1479509891143"
        },
        "sha256:647b02e962848f3996429cc9426832164a57e1f814dd240dbadb0adb03072e6a": {
            "imageSizeBytes": "27598207",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.6"
            ],
            "timeCreatedMs": "1497638709648",
            "timeUploadedMs": "1497644934760"
        },
        "sha256:648d1bfb13eadb6a481c23da1b9732ac7d5f9ad84cb4ac6b7a86024cbcd55b13": {
            "imageSizeBytes": "31745371",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.10-beta.0"
            ],
            "timeCreatedMs": "1508437908697",
            "timeUploadedMs": "1508442294287"
        },
        "sha256:6652806864d1d3532ff4739a9d6cae9a05a930f606ca33f6d05903e2a587bf38": {
            "imageSizeBytes": "29815354",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.8-beta.0"
            ],
            "timeCreatedMs": "1516142410712",
            "timeUploadedMs": "1516151421432"
        },
        "sha256:67bc54a4f55103ef50ddb662e210f1489cbbfad8174b055fe21cfb2f8d0e3a3f": {
            "imageSizeBytes": "20570309",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.1"
            ],
            "timeCreatedMs": "1468778350564",
            "timeUploadedMs": "1468778940869"
        },
        "sha256:67f602ab154391ed819a292feab12d24480bd95d35214f5744967b158bb44e36": {
            "imageSizeBytes": "8863788",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "9680e782e08a1a1c94c656190011bd02"
            ],
            "timeCreatedMs": "1433045619235",
            "timeUploadedMs": "1484178888522"
        },
        "sha256:6841027e906b44e26339b097becba5695f4e1a1fa2d94ea8c85b19af1b91e130": {
            "imageSizeBytes": "13032729",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.0-beta.1"
            ],
            "timeCreatedMs": "1457736345327",
            "timeUploadedMs": "1457736584826"
        },
        "sha256:693eeb779cbc974dad06a7f3bcc57e6fd4923d19d812cb6437fa8f478d68e122": {
            "imageSizeBytes": "20569995",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.1-beta.1"
            ],
            "timeCreatedMs": "1468626660506",
            "timeUploadedMs": "1468628332265"
        },
        "sha256:6c667cc854e8c5de12ce55ca92b5f9a6c2564ee48405869993c91015e003bc92": {
            "imageSizeBytes": "34497510",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-beta.1"
            ],
            "timeCreatedMs": "1527810970721",
            "timeUploadedMs": "1527813483079"
        },
        "sha256:6c74eb9d3aed062d604a518f7e346ea92ee4466ede47671b756d780e3993e6f7": {
            "imageSizeBytes": "29797940",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.5"
            ],
            "timeCreatedMs": "1512663990151",
            "timeUploadedMs": "1512668609984"
        },
        "sha256:6d5aa429c2b0806e4b6d1d179054d6deee46eec0aabe7bd7bd6abff97be36ae7": {
            "imageSizeBytes": "27571026",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.4"
            ],
            "timeCreatedMs": "1495220057373",
            "timeUploadedMs": "1495223454843"
        },
        "sha256:6edc40dbb75ebb57b9e281c3a7dec67109676ff0387264906859df068f6e76b4": {
            "imageSizeBytes": "27721807",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.10-beta.0"
            ],
            "timeCreatedMs": "1487167164567",
            "timeUploadedMs": "1487168168088"
        },
        "sha256:71273b57d811654620dc7a0d22fd893d9852b6637616f8e7e3f4507c60ea7357": {
            "imageSizeBytes": "29815329",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.7"
            ],
            "timeCreatedMs": "1516141621414",
            "timeUploadedMs": "1516150669644"
        },
        "sha256:73125d87f45b94ab85c95db74e4e31803d139fc6f3571ca50ce08337a41e25ae": {
            "imageSizeBytes": "30210479",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.12"
            ],
            "timeCreatedMs": "1524527063143",
            "timeUploadedMs": "1524543415336"
        },
        "sha256:734ed936d8a199d6c65ea5c522e49426156f6252ed8f799f0964b9fce9781652": {
            "imageSizeBytes": "34121854",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-alpha.2"
            ],
            "timeCreatedMs": "1525263694536",
            "timeUploadedMs": "1525276254121"
        },
        "sha256:73d4a6883a4f4f78fce1829afad2dce54c31af2ec689d45728df11506283f1a2": {
            "imageSizeBytes": "31497890",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.1"
            ],
            "timeCreatedMs": "1499998285649",
            "timeUploadedMs": "1500004943459"
        },
        "sha256:74893e99fe1977286c6f18eb83a67f843740bc4868a97929955f8534e5474f84": {
            "imageSizeBytes": "20412716",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-beta.0"
            ],
            "timeCreatedMs": "1465595435369",
            "timeUploadedMs": "1465596421395"
        },
        "sha256:7612e0594cbc78dba21f22299d40fd0413ad1e8b54b7d6f8b36f6852d5b910f2": {
            "imageSizeBytes": "20565543",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.4-beta.0"
            ],
            "timeCreatedMs": "1469220874486",
            "timeUploadedMs": "1469225012946"
        },
        "sha256:762a4da324039708487db08b1faa870ceb9827e49d05c51b24b80be10c79922a": {
            "imageSizeBytes": "32166257",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.16-beta.0"
            ],
            "timeCreatedMs": "1521473248497",
            "timeUploadedMs": "1521491209382"
        },
        "sha256:77230dfef485cdc4149ed04fe2b9ed4d332499f726760f8e96a84198ad5cfcea": {
            "imageSizeBytes": "20571013",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.2-beta.0"
            ],
            "timeCreatedMs": "1468627735138",
            "timeUploadedMs": "1468628946438"
        },
        "sha256:772a92eb61cdf7be3014e08c66ec08b5cc4e6fbd310b2bf1a805f75f2f826a33": {
            "imageSizeBytes": "20557647",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.5-beta.0"
            ],
            "timeCreatedMs": "1470071281169",
            "timeUploadedMs": "1470072541120"
        },
        "sha256:781f3f3fe06aa7e80bfe78df708639d81684ef0e809c536a0773c02b176f7552": {
            "imageSizeBytes": "22811850",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.5-beta.0"
            ],
            "timeCreatedMs": "1488936898352",
            "timeUploadedMs": "1488939683941"
        },
        "sha256:783191f3db9a4b51c8aac432c767fb783d6e678d50965200d2db5cfb5c0dd7da": {
            "imageSizeBytes": "27759950",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.11"
            ],
            "timeCreatedMs": "1474671627362",
            "timeUploadedMs": "1474672450672"
        },
        "sha256:7841948c6d0d84f9a9d7192d4a5fb685927ca9829041d9f7a65b01057aab68e9": {
            "imageSizeBytes": "31487464",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-alpha.1"
            ],
            "timeCreatedMs": "1497912946903",
            "timeUploadedMs": "1497916932231"
        },
        "sha256:78a2350eaab62dfa23ceb21f2f45a2ed5c3ef1939ba08917937e1ede6a79a7a2": {
            "imageSizeBytes": "29771551",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-rc.1"
            ],
            "timeCreatedMs": "1506132717312",
            "timeUploadedMs": "1506134336730"
        },
        "sha256:796d2d20ed83ef2df39c05bfe34541d24c97e0b3190dc94c504fc92e2c36ee3f": {
            "imageSizeBytes": "23786206",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.8-beta.0"
            ],
            "timeCreatedMs": "1493293161512",
            "timeUploadedMs": "1493295787273"
        },
        "sha256:79d444e6cb940079285109aaa5f6a97e5c0a5568f6606e003ed279cd90bcf1ca": {
            "imageSizeBytes": "31689595",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.8"
            ],
            "timeCreatedMs": "1526930977230",
            "timeUploadedMs": "1526939145283"
        },
        "sha256:7d77bee7d4f9822692a30d419cbe6f226a0db38a3166f343258c7a34f3817b83": {
            "imageSizeBytes": "14998627",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-alpha.3"
            ],
            "timeCreatedMs": "1461802134450",
            "timeUploadedMs": "1461864341997"
        },
        "sha256:7dd91e4670b3563de04fee77c3206d84ca556ea88f95a40ee510ca2b77dc49d7": {
            "imageSizeBytes": "22812218",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.5"
            ],
            "timeCreatedMs": "1490142810585",
            "timeUploadedMs": "1490145656693"
        },
        "sha256:7fd5204f194111fdae83e0677378778ada6b6b729e9f0479b256e804da366f67": {
            "imageSizeBytes": "30212333",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.14"
            ],
            "timeCreatedMs": "1529360737872",
            "timeUploadedMs": "1529427550040"
        },
        "sha256:80d3750f3e70db2cdc9d1154b78f6c9961ea41513f7f1d3ed77c461ee54d5734": {
            "imageSizeBytes": "27553227",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0"
            ],
            "timeCreatedMs": "1490719646544",
            "timeUploadedMs": "1490728092896"
        },
        "sha256:8114f937a1d70adcabd01d49b81ad7b96bbdfd03418f5e505e8e62b43c4fcac2": {
            "imageSizeBytes": "27611763",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.10"
            ],
            "timeCreatedMs": "1505330364713",
            "timeUploadedMs": "1505339028175"
        },
        "sha256:821751c5de9395fe7521943ca81aefcfcc2535577669274fee7b1a07e3935bd3": {
            "imageSizeBytes": "26353031",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-alpha.2"
            ],
            "timeCreatedMs": "1487120597075",
            "timeUploadedMs": "1487123030435"
        },
        "sha256:82b2d22fe5575b30e1a96b1a7a739fad46f1fc7d6cb79df50b28ef293c1e2e8c": {
            "imageSizeBytes": "29146234",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-alpha.4"
            ],
            "timeCreatedMs": "1495129308988",
            "timeUploadedMs": "1495135312168"
        },
        "sha256:82eb9137fac472365a2347eba6cbdff23c829f9a750e7b5e54604c61eb04a68e": {
            "imageSizeBytes": "13038778",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.3-beta.0"
            ],
            "timeCreatedMs": "1460077157739",
            "timeUploadedMs": "1460077434087"
        },
        "sha256:83fedf1fdedfd0efd1cca7ea6de0955915b2ff36a6370d93578b4ff008792ee7": {
            "imageSizeBytes": "27553516",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.2-beta.0"
            ],
            "timeCreatedMs": "1491256222723",
            "timeUploadedMs": "1491259656182"
        },
        "sha256:84c92820bd3532aa66a62d79391aca85c75c3d595cd191ad66233b5e40f08506": {
            "imageSizeBytes": "31206770",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-alpha.1"
            ],
            "timeCreatedMs": "1513631619465",
            "timeUploadedMs": "1513633523556"
        },
        "sha256:84f6550b0f7710612c72d44a5732a666d49ef1a99d88efb577b123252cbdc6ac": {
            "imageSizeBytes": "28122838",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-beta.1"
            ],
            "timeCreatedMs": "1488496569993",
            "timeUploadedMs": "1488499128769"
        },
        "sha256:858e34bb3495c2f6f45df8c7ee5ca5d3afade6329dfddd9d71537667aba5a824": {
            "imageSizeBytes": "29788704",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-alpha.3"
            ],
            "timeCreatedMs": "1503529717329",
            "timeUploadedMs": "1503530935603"
        },
        "sha256:85b625f8cdcd067550e1c512de9fdb76c490578a7457014ab4424088426e470f": {
            "imageSizeBytes": "32912324",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.3-beta.0"
            ],
            "timeCreatedMs": "1524824918443",
            "timeUploadedMs": "1524829241737"
        },
        "sha256:869c32a3654e7e3dfe93c51ab0af05d177a156578876c5ff28b4bc834f1770cc": {
            "imageSizeBytes": "20571031",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.2"
            ],
            "timeCreatedMs": "1468780563092",
            "timeUploadedMs": "1468782222947"
        },
        "sha256:872e3d4286a8ef4338df59945cb0d64c2622268ceb3e8a2ce7b52243279b02d0": {
            "imageSizeBytes": "29793060",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.3"
            ],
            "timeCreatedMs": "1510167033842",
            "timeUploadedMs": "1510173974983"
        },
        "sha256:8757c7494c8e2c9a725cd41ca8f63dec564dd9d2db64a629d9da12cbe9e858be": {
            "imageSizeBytes": "30196740",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.9-beta.0"
            ],
            "timeCreatedMs": "1518214387961",
            "timeUploadedMs": "1518218198917"
        },
        "sha256:89ef9bb1b81a268e1bc1a0b10edf1b2aa84ac3dd4fac06efedab7ab24abcf156": {
            "imageSizeBytes": "31740380",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.11",
                "v1.7.12-beta.0"
            ],
            "timeCreatedMs": "1511637119313",
            "timeUploadedMs": "1511649822769"
        },
        "sha256:8ac6a18f2d6af3d178f34a4457f59c0879fbf7590bc5ad7a3357725ba8140285": {
            "imageSizeBytes": "22680629",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.0-beta.3"
            ],
            "timeCreatedMs": "1481234903643",
            "timeUploadedMs": "1481238165755"
        },
        "sha256:8b704385ae24e1782c8301194835f4fe44531d061b1626e371d88d20244d98da": {
            "imageSizeBytes": "27574925",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.5-beta.0"
            ],
            "timeCreatedMs": "1494891364755",
            "timeUploadedMs": "1494894896210"
        },
        "sha256:8b7a675c6fdda0469e971e5b1f3e902bb71c36396faf506f1511052705a5b0ee": {
            "imageSizeBytes": "32910426",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0"
            ],
            "timeCreatedMs": "1522086439342",
            "timeUploadedMs": "1522105395384"
        },
        "sha256:8bfc3dbe321d4bf4b5b9fd2f13053bb9954f1379df8813f53dc799251dc56b5d": {
            "imageSizeBytes": "13039815",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.6-beta.0"
            ],
            "timeCreatedMs": "1467051459705",
            "timeUploadedMs": "1467051983718"
        },
        "sha256:8cabfffd63265ef40b1aeae36756b91bae46da0f19502394584f9fac8f1c11bf": {
            "imageSizeBytes": "27734698",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.6"
            ],
            "timeCreatedMs": "1474049266080",
            "timeUploadedMs": "1474060354033"
        },
        "sha256:8dfc23c126a60f53e41b8d4dcbf638df0b80c0fc2959080b1a10b0fced1299c6": {
            "imageSizeBytes": "20557185",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.5"
            ],
            "timeCreatedMs": "1470947712024",
            "timeUploadedMs": "1470949555930"
        },
        "sha256:8e127857b74640fe4f9cd35597f5430e5af7ba440ebc27883a79d8384d5847bf": {
            "imageSizeBytes": "27733967",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.1"
            ],
            "timeCreatedMs": "1476124068185",
            "timeUploadedMs": "1476126483853"
        },
        "sha256:8ec52cd0d6840c1cffd622ef2957364aa5ba87372562cdce89aad9bd19cfc64a": {
            "imageSizeBytes": "27611966",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.11-beta.0"
            ],
            "timeCreatedMs": "1505334108213",
            "timeUploadedMs": "1505339814404"
        },
        "sha256:8ec98f05bd8c2e2ce459d0b85474a24985dfdbd5136953ac0a864fcf4f71375f": {
            "imageSizeBytes": "13032730",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.8-beta.0"
            ],
            "timeCreatedMs": "1477168247684",
            "timeUploadedMs": "1477168749030"
        },
        "sha256:8fea627a5a9f6a5edf400aa1c2d0758ce8b7b5fbee43c0304fb81a64748fc76d": {
            "imageSizeBytes": "27598267",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.8-beta.0"
            ],
            "timeCreatedMs": "1499277497184",
            "timeUploadedMs": "1499280581006"
        },
        "sha256:9011af032a490c6db9c1708c1bace21a79b097c67fe2aa4a191712368cb021ae": {
            "imageSizeBytes": "20557104",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.6-beta.0"
            ],
            "timeCreatedMs": "1470948785499",
            "timeUploadedMs": "1470950218801"
        },
        "sha256:92cf7632e4bca87974aac6dd53fd608dba8830dd75d73ed5e7aef9c7b2f152f5": {
            "imageSizeBytes": "22680465",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.0"
            ],
            "timeCreatedMs": "1481586356785",
            "timeUploadedMs": "1481592759146"
        },
        "sha256:94fbefcd1473f45dbfe4953628118f55217a1f77b814de125ec64bdcfabad2cf": {
            "imageSizeBytes": "32907837",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-beta.3"
            ],
            "timeCreatedMs": "1520891397225",
            "timeUploadedMs": "1520893239327"
        },
        "sha256:960ce3bb6cf66ace03fd5897beb01db73be42c7db567d42b9bf1be17e42fbac3": {
            "imageSizeBytes": "32863903",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-beta.2"
            ],
            "timeCreatedMs": "1520456583695",
            "timeUploadedMs": "1520458726665"
        },
        "sha256:9649ccad4d5f5771ffec1ca185eb83f8bd7a79faab74928b657411a8aa25362d": {
            "imageSizeBytes": "27553526",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-rc.1"
            ],
            "timeCreatedMs": "1490383816191",
            "timeUploadedMs": "1490386325963"
        },
        "sha256:965c830a29f8e40aad38607bc211346170116ab986cdbe1e1d638d7139662f7a": {
            "imageSizeBytes": "31740159",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.10"
            ],
            "timeCreatedMs": "1509731546501",
            "timeUploadedMs": "1509742069132"
        },
        "sha256:97200ebf99d11a6682eb849a5f6113f0272ab625ecdb068c11bee01f4922e0c8": {
            "imageSizeBytes": "13040352",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.6"
            ],
            "timeCreatedMs": "1468600172753",
            "timeUploadedMs": "1468600751266"
        },
        "sha256:9738f7c6fd44253513e7a11fb05239f8e0adda957ea502d8f1e7428374233fbf": {
            "imageSizeBytes": "23801561",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.8"
            ],
            "timeCreatedMs": "1506818655976",
            "timeUploadedMs": "1506841191682"
        },
        "sha256:97f7a4cf94a0c32604eb42eb16b0ee78752d8600091461194fcecb8deb259499": {
            "imageSizeBytes": "13040367",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.7-beta.0"
            ],
            "timeCreatedMs": "1468600469227",
            "timeUploadedMs": "1468601030236"
        },
        "sha256:9931395c32a248d5c313c0b135d8158ed251fe5ab2027fac099308cdcc23d45c": {
            "imageSizeBytes": "23786217",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.7"
            ],
            "timeCreatedMs": "1493288779662",
            "timeUploadedMs": "1493295189423"
        },
        "sha256:995200084f9cdb3ea840513582ef6c7bdea204fc8e16449a36645d49651406bb": {
            "imageSizeBytes": "35098831",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-beta.2"
            ],
            "timeCreatedMs": "1528389360559",
            "timeUploadedMs": "1528391083283"
        },
        "sha256:99fe756ac35ab6f35ad09ddc09787574e063c1eed9282008498b8d1618c4cb9c": {
            "imageSizeBytes": "23569635",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.6"
            ],
            "timeCreatedMs": "1490710767431",
            "timeUploadedMs": "1490716767133"
        },
        "sha256:9a35d98a547e9504c6950cb652f75b49dd2d59d31c75236d5852691aaeabeb8e": {
            "imageSizeBytes": "28133465",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-alpha.1"
            ],
            "timeCreatedMs": "1491510044179",
            "timeUploadedMs": "1491512330444"
        },
        "sha256:9ba40816080fe51127be035fed767f2dd44ef451c42946a7c36d57d550f313db": {
            "imageSizeBytes": "32910483",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-beta.4"
            ],
            "timeCreatedMs": "1521008466559",
            "timeUploadedMs": "1521010082273"
        },
        "sha256:9cd35f406c0b3e74233a757820713bb98d931f6bd1ae4c24efabd49bbdd71062": {
            "imageSizeBytes": "13073051",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.0-beta.0"
            ],
            "timeCreatedMs": "1457136551157",
            "timeUploadedMs": "1457136772253"
        },
        "sha256:9d2e12ce12d3e7dc0a28b7fc98c678bb6f2b435b82635d28140fd390511db005": {
            "imageSizeBytes": "27709455",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.7"
            ],
            "timeCreatedMs": "1481345760082",
            "timeUploadedMs": "1481347541565"
        },
        "sha256:9e3fe458aa708a40c176005216ceab80e9ec1d51fea71c6fb916e10de97aa8f9": {
            "imageSizeBytes": "29773181",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.1-beta.0"
            ],
            "timeCreatedMs": "1506642556949",
            "timeUploadedMs": "1506658918257"
        },
        "sha256:9fc0d2703d28212e4a391138b0066d548b0ad3726c44eb98efe45496e27aca42": {
            "imageSizeBytes": "32311098",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-alpha.2"
            ],
            "timeCreatedMs": "1516993642256",
            "timeUploadedMs": "1516995098623"
        },
        "sha256:a1359410a7cac64851e751c4e49b531a4e233e4919db59f219c80f6a6c03f1bc": {
            "imageSizeBytes": "29284513",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-beta.0"
            ],
            "timeCreatedMs": "1504314295353",
            "timeUploadedMs": "1504321637664"
        },
        "sha256:a1a12f95f18a95d4400a66f4ba1858804c9d30b1327bf0366de2111bebe3b48a": {
            "imageSizeBytes": "31668915",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.5-beta.0"
            ],
            "timeCreatedMs": "1520874736927",
            "timeUploadedMs": "1520880315082"
        },
        "sha256:a25dc2d57e784cf25228fcb3b6fdeba04cdea78b50d432e234b4c2c562387641": {
            "imageSizeBytes": "27598208",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.7-beta.0"
            ],
            "timeCreatedMs": "1497642566783",
            "timeUploadedMs": "1497645654667"
        },
        "sha256:a2cd61c7b5c88ecf7d9c76c86033b715616a726d50cfb38628532b7f7ae17354": {
            "imageSizeBytes": "27721791",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.9"
            ],
            "timeCreatedMs": "1487166159688",
            "timeUploadedMs": "1487167682589"
        },
        "sha256:a312a72821fb748ad45563f52b2f89136c3389f9b57e72f9e838cb333603bc03": {
            "imageSizeBytes": "27616273",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.12"
            ],
            "timeCreatedMs": "1508954720057",
            "timeUploadedMs": "1508960935034"
        },
        "sha256:a3fb9e02ad92f6f3c367da368a4868b503460d5ec8da0e0f9be0688f63067b00": {
            "imageSizeBytes": "13038674",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.4-beta.0"
            ],
            "timeCreatedMs": "1461803561479",
            "timeUploadedMs": "1461804030984"
        },
        "sha256:a52727531c1893c3dca20291afa45d098707f1a1f93c7d547b38c89e0ae22b88": {
            "imageSizeBytes": "29792872",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.4-beta.0"
            ],
            "timeCreatedMs": "1510169022147",
            "timeUploadedMs": "1510174765038"
        },
        "sha256:a52dd6d990bbdfa6ee82d60b83a95e1d24393b2727ae6199f51cf3e622da4c22": {
            "imageSizeBytes": "27725162",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.1"
            ],
            "timeCreatedMs": "1473538942285",
            "timeUploadedMs": "1473539608098"
        },
        "sha256:a5382344aa373a90bc87d3baa4eda5402507e8df5b8bfbbad392c4fff715f043": {
            "imageSizeBytes": "31546923",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.3"
            ],
            "timeCreatedMs": "1518008657645",
            "timeUploadedMs": "1518211837967"
        },
        "sha256:a6c4b6b2429d0a15d30a546226e01b1164118e022ad40f3ece2f95126f1580f5": {
            "imageSizeBytes": "32934456",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.3"
            ],
            "timeCreatedMs": "1526897289993",
            "timeUploadedMs": "1526900443170"
        },
        "sha256:a6d55acbdffbccfc85c4a9b64a031ffb17655cd20e625c90427b46a8f7bf4120": {
            "imageSizeBytes": "30212312",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.15-beta.0"
            ],
            "timeCreatedMs": "1529361518749",
            "timeUploadedMs": "1529428052790"
        },
        "sha256:a89ccf0ebe8af731e209584a9e66e03005a52c9249f032dbdb3945783535736d": {
            "imageSizeBytes": "13038674",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [],
            "timeCreatedMs": "1461360777827",
            "timeUploadedMs": "1461364874212"
        },
        "sha256:a89cfd287c9bdf588d3f1271a8ec5d69b93e8acd6b09279b2cb6474ce71e1262": {
            "imageSizeBytes": "26611226",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-alpha.3"
            ],
            "timeCreatedMs": "1487273890183",
            "timeUploadedMs": "1487276297691"
        },
        "sha256:a9ccc205760319696d2ef0641de4478ee90fb0b75fbe6c09b1d64058c8819f97": {
            "imageSizeBytes": "31718774",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.2"
            ],
            "timeCreatedMs": "1500626555579",
            "timeUploadedMs": "1500638077290"
        },
        "sha256:aa5674d9cfb2e7c445d10b25cda16f0db03455a1adf4550bdc9121dc7fd5b504": {
            "imageSizeBytes": "31726280",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.5"
            ],
            "timeCreatedMs": "1504172192028",
            "timeUploadedMs": "1504184101307"
        },
        "sha256:aba16db38351970921151cbb952f685c52178725e213b002e6d291a1a5052ad8": {
            "imageSizeBytes": "34139113",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.12.0-alpha.0"
            ],
            "timeCreatedMs": "1526528609165",
            "timeUploadedMs": "1526531208527"
        },
        "sha256:ad273ffc70fbc7b487bb05092e56b30736674275135c55db9fd65be048ea0b80": {
            "imageSizeBytes": "13072681",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-alpha.0"
            ],
            "timeCreatedMs": "1457135856626",
            "timeUploadedMs": "1457136158596"
        },
        "sha256:aff52721235b94ab19eba8e515e308196aa8f9c9a85fb8de191592e6d2037354": {
            "imageSizeBytes": "31688178",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.8-beta.0"
            ],
            "timeCreatedMs": "1524098858247",
            "timeUploadedMs": "1524155580029"
        },
        "sha256:b00ef533fd885e5bdb1072096abf474b048795f529b359e6c93e881032294603": {
            "imageSizeBytes": "13035223",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.0"
            ],
            "timeCreatedMs": "1458165906539",
            "timeUploadedMs": "1458166139238"
        },
        "sha256:b0d547e39e16801cba391375ec58fda6f388a6b397bcdf08080e667f0c23e905": {
            "imageSizeBytes": "20571015",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.3-beta.0"
            ],
            "timeCreatedMs": "1468781594059",
            "timeUploadedMs": "1468782852611"
        },
        "sha256:b29e1f6301ca7f05458d20ef7e31190d534344c3c44a87c57154b180751e24ad": {
            "imageSizeBytes": "27616336",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.13",
                "v1.6.14-beta.0"
            ],
            "timeCreatedMs": "1511383074213",
            "timeUploadedMs": "1511389662522"
        },
        "sha256:b398602f2df521249613e28b4008a65b8c4bede52c8ad1804d147c980e16f7cb": {
            "imageSizeBytes": "20568399",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.7-beta.0"
            ],
            "timeCreatedMs": "1472236672063",
            "timeUploadedMs": "1472237947282"
        },
        "sha256:b3ed1c08f89fe72ae318f244064a3f9b131d017f024fceafc0318039ecf763d7": {
            "imageSizeBytes": "27668830",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-alpha.3"
            ],
            "timeCreatedMs": "1472150868435",
            "timeUploadedMs": "1472153912248"
        },
        "sha256:b43fb501ce4ceae2f25a2ebd956d4fb9bb7b9c1d6a37dabd6b89a9b1f19bc1b7": {
            "imageSizeBytes": "33369376",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.6-beta.0"
            ],
            "timeCreatedMs": "1529585533052",
            "timeUploadedMs": "1529588835548"
        },
        "sha256:b4a5467c922ec6ffa4e3c4f5a476a2ef9732d60890dd834577cf17e7974692bc": {
            "imageSizeBytes": "30475552",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-alpha.2"
            ],
            "timeCreatedMs": "1509574804174",
            "timeUploadedMs": "1509575988049"
        },
        "sha256:b4da427ad79e4c594237b3e47f94dec1b8b8a6fdc45a8e3e56abce2aaaabf47e": {
            "imageSizeBytes": "31688088",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.7"
            ],
            "timeCreatedMs": "1524098034585",
            "timeUploadedMs": "1524154925263"
        },
        "sha256:b4f04c9077016f78a5a9862d2fe9019e613a496713e70920a036af15c5c22247": {
            "imageSizeBytes": "31745160",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.12",
                "v1.7.13-beta.0"
            ],
            "timeCreatedMs": "1514539457185",
            "timeUploadedMs": "1514555933666"
        },
        "sha256:b5955a6607e01583c3cf395d410ea17c08e97944b9bbe3f416010db887183809": {
            "imageSizeBytes": "12776720",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.0-alpha.8"
            ],
            "timeCreatedMs": "1455933588215",
            "timeUploadedMs": "1455936655255"
        },
        "sha256:b6498fd4b6a150f9378e1079242e8ed2b924851eca9d2a5239932875e45ebe70": {
            "imageSizeBytes": "35019714",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-rc.2"
            ],
            "timeCreatedMs": "1529902185594",
            "timeUploadedMs": "1529947573290"
        },
        "sha256:b664a5dabed80309eea8c577d485e9c53350819d2a172f1499e47be842086610": {
            "imageSizeBytes": "13032825",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.7"
            ],
            "timeCreatedMs": "1477167965851",
            "timeUploadedMs": "1477168547736"
        },
        "sha256:b66edf53b407b09497a8b7f7eac78f42a62b9cd6cbd5862c2849039e0f867275": {
            "imageSizeBytes": "27616002",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.12-beta.0"
            ],
            "timeCreatedMs": "1506638439299",
            "timeUploadedMs": "1506641890928"
        },
        "sha256:b6ce88a4be8d9c3a7b1feda6e1f4873735c0f00635ae31e5e903a1a8d50e3076": {
            "imageSizeBytes": "22357858",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.6.0-alpha.0"
            ],
            "timeCreatedMs": "1478394772451",
            "timeUploadedMs": "1478401203042"
        },
        "sha256:b73d1823e8358cc51617ce736268e47b1aa7a986ca6203745ab5313713b4466f": {
            "imageSizeBytes": "27796528",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.6-beta.0"
            ],
            "timeCreatedMs": "1477706508798",
            "timeUploadedMs": "1477707571525"
        },
        "sha256:b799af80fa8fea6a411f5c0784a4c5e28587656863e043bfefc6a4b877439a61": {
            "imageSizeBytes": "31671077",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.5"
            ],
            "timeCreatedMs": "1521477002513",
            "timeUploadedMs": "1521483337080"
        },
        "sha256:b8f66ad41940b4efd09f243541c7b8658f56e7d6450b9eef97194e0244c72b3b": {
            "imageSizeBytes": "13035282",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.1-beta.0"
            ],
            "timeCreatedMs": "1458166709023",
            "timeUploadedMs": "1458166938695"
        },
        "sha256:ba4d5464a81241a604e6e591367d09e9d874cfac278c79c9e14974fabf2b3b04": {
            "imageSizeBytes": "32289805",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-alpha.3"
            ],
            "timeCreatedMs": "1517512476586",
            "timeUploadedMs": "1517521277779"
        },
        "sha256:ba6da0ea3abf67a00ca037dcdc9601f2b921c9f2cba4cf5682a1e6675edb2f7e": {
            "imageSizeBytes": "13036891",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.1"
            ],
            "timeCreatedMs": "1459534523011",
            "timeUploadedMs": "1459534767891"
        },
        "sha256:ba6f72c19534c7974139805845554a15d0227f8a9053059d98152712193de27a": {
            "imageSizeBytes": "35019724",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-rc.3"
            ],
            "timeCreatedMs": "1530030219333",
            "timeUploadedMs": "1530047818134"
        },
        "sha256:be2fca37b708dddfa122d6a81f682b1c9394f1b984cb40055b2e0609edd72eab": {
            "imageSizeBytes": "22812516",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.4-beta.0"
            ],
            "timeCreatedMs": "1487145587794",
            "timeUploadedMs": "1487149397847"
        },
        "sha256:bed05fa1dd564269dd01d076a64e3604618990402aadad8d8b0d301e3f434e90": {
            "imageSizeBytes": "30207097",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.10-beta.0"
            ],
            "timeCreatedMs": "1520874058143",
            "timeUploadedMs": "1520880229536"
        },
        "sha256:bf7a245355a65f2ace405a3d85d3044567baba2ad10718048fd586a1eb79e1c4": {
            "imageSizeBytes": "27724197",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.12"
            ],
            "timeCreatedMs": "1492722436621",
            "timeUploadedMs": "1492723972509"
        },
        "sha256:bffd6129137110514bddfcab4d8d48d35fa6e1b0b3f3422f1c9fc6319566cf0e": {
            "imageSizeBytes": "13038770",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.2"
            ],
            "timeCreatedMs": "1460076184170",
            "timeUploadedMs": "1460076412381"
        },
        "sha256:c14cade9ba711fb6791efb55f6ad4132049102124e79db79b22b29381b9cb1d7": {
            "imageSizeBytes": "35033033",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.2-beta.0"
            ],
            "timeCreatedMs": "1531856955741",
            "timeUploadedMs": "1531933909055"
        },
        "sha256:c31ca5ee4edafe8fa3e6d4e9bd6fbd855c4ebc59cdfb30116b5197a32ce2b2bd": {
            "imageSizeBytes": "32934382",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.4-beta.0"
            ],
            "timeCreatedMs": "1526898100895",
            "timeUploadedMs": "1526900992910"
        },
        "sha256:c4600d5c68124efd71d7c084d4cfdb53648662487d137bc37872fd6d295d6d30": {
            "imageSizeBytes": "32514924",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-beta.0"
            ],
            "timeCreatedMs": "1518640058233",
            "timeUploadedMs": "1518643888193"
        },
        "sha256:c584b9f538890675bac1b5d5ffe35e145297c93f737fa6cea3f33724c17b3ad1": {
            "imageSizeBytes": "27559441",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.2"
            ],
            "timeCreatedMs": "1492634560163",
            "timeUploadedMs": "1492641593845"
        },
        "sha256:c5a08cbce957415b0e53473465f6e8a7a3117451319a5eafd0d93b1fd2e1d3c3": {
            "imageSizeBytes": "27613457",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.8"
            ],
            "timeCreatedMs": "1501784503237",
            "timeUploadedMs": "1501790308260"
        },
        "sha256:c6bdf4fccc46cbf0a70f1ff718e1f02b908417d395b70907aa69278de0a03d4b": {
            "imageSizeBytes": "30207407",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.9"
            ],
            "timeCreatedMs": "1520873291196",
            "timeUploadedMs": "1520879477713"
        },
        "sha256:c7135b816bba979ea973ac113b91e91a86134d972b88e50f1cf6751adb63dc38": {
            "imageSizeBytes": "27752346",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.8"
            ],
            "timeCreatedMs": "1474233238053",
            "timeUploadedMs": "1474235018947"
        },
        "sha256:c7bcbeac1b498478fcb211410e9ce712bd4079cd16a1d4f0912d0e9f3b5eb780": {
            "imageSizeBytes": "27735663",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.3-beta.0"
            ],
            "timeCreatedMs": "1476491864922",
            "timeUploadedMs": "1476492949895"
        },
        "sha256:c8c71b57e7aa6521c33771e064039a8141282b0a0f50dc27882229dbd2804a94": {
            "imageSizeBytes": "27711186",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.6"
            ],
            "timeCreatedMs": "1478928514397",
            "timeUploadedMs": "1478930205335"
        },
        "sha256:c98ace532fb024551c242981b7c7316d4053057b639b0cac9c77f732d828c7c4": {
            "imageSizeBytes": "27719493",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.8"
            ],
            "timeCreatedMs": "1484188022549",
            "timeUploadedMs": "1484189657468"
        },
        "sha256:c9ae5cdddf75f23a192ec029c19fc7d1c1aaa8cb7b04643f86d338a901cf9d96": {
            "imageSizeBytes": "13036835",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.2-beta.0"
            ],
            "timeCreatedMs": "1459535409840",
            "timeUploadedMs": "1459535853837"
        },
        "sha256:ca14812da4eb2b45bb4f15eacc77fbe1ba4c37b931d446d8340408c05a415960": {
            "imageSizeBytes": "21484876",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-alpha.2"
            ],
            "timeCreatedMs": "1469821148053",
            "timeUploadedMs": "1469822531936"
        },
        "sha256:ca54685b89b3e1809ea3fa9936e32e3a05083a84483813604178275e02352454": {
            "imageSizeBytes": "33369490",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.5"
            ],
            "timeCreatedMs": "1529584660143",
            "timeUploadedMs": "1529588285499"
        },
        "sha256:cc55644caa02f89afbfbc83eba346dfba327e0acb37571ff79320fe556810a5f": {
            "imageSizeBytes": "30209168",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.11-beta.0"
            ],
            "timeCreatedMs": "1521484320388",
            "timeUploadedMs": "1521489299825"
        },
        "sha256:cd857aa4b07596335d1a541a06157f6b89fb34e5642caf513a206b725eb28e2a": {
            "imageSizeBytes": "27719566",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.9-beta.0"
            ],
            "timeCreatedMs": "1484189097528",
            "timeUploadedMs": "1484190143689"
        },
        "sha256:d0183503d3dda0d63e19e49de2d2bff72b6edb01ea5df111bd2947a626aec687": {
            "imageSizeBytes": "27796435",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.5"
            ],
            "timeCreatedMs": "1477705513779",
            "timeUploadedMs": "1477707083699"
        },
        "sha256:d04ea477d3fcdb85fddcfd8bcb2ad0d60d4e506e01806e108b9676e85ddbf65d": {
            "imageSizeBytes": "31722440",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.5-beta.0"
            ],
            "timeCreatedMs": "1502969203752",
            "timeUploadedMs": "1502990048137"
        },
        "sha256:d1ff73ac23cd1bd846c712f4ce32053af3b933d9c105e21d0e0da37560ee295e": {
            "imageSizeBytes": "29781516",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.3-beta.0"
            ],
            "timeCreatedMs": "1508877260275",
            "timeUploadedMs": "1508879408362"
        },
        "sha256:d29a2e1ac89838c267950bea6d9f869a2e21d7775a17132b018ce9780c33416a": {
            "imageSizeBytes": "29791960",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.7-beta.0"
            ],
            "timeCreatedMs": "1513840699203",
            "timeUploadedMs": "1513879727209"
        },
        "sha256:d2e2f62381390e820a0d84f9650daa8537e1d3d6a1c02f1bbe027abdb0ab65ab": {
            "imageSizeBytes": "31689628",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.9-beta.0"
            ],
            "timeCreatedMs": "1526931863739",
            "timeUploadedMs": "1526939678069"
        },
        "sha256:d3a004568414444541de96aa9f6be9b926eceead5d0f874bda8094056441bad7": {
            "imageSizeBytes": "14192089",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-alpha.2"
            ],
            "timeCreatedMs": "1460404150190",
            "timeUploadedMs": "1460404532894"
        },
        "sha256:d3baf227db6b3a06a53d207c74398e4abac67a95b40e51c9f5b22424ea365422": {
            "imageSizeBytes": "27724818",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.2"
            ],
            "timeCreatedMs": "1473695860046",
            "timeUploadedMs": "1473696391807"
        },
        "sha256:d403e676eae2c85a37f951714333e84dfe35b07886054dfe9b62db7f64127569": {
            "imageSizeBytes": "20608944",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.7"
            ],
            "timeCreatedMs": "1473722477817",
            "timeUploadedMs": "1473724151979"
        },
        "sha256:d4387dff51b1f9c94cd1cfac3a4694347970b90e911159ac6fe2d090c96a6184": {
            "imageSizeBytes": "27553230",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.1"
            ],
            "timeCreatedMs": "1491252863098",
            "timeUploadedMs": "1491258943975"
        },
        "sha256:d4d9db2294ffa3168cf769e32d5c256bfe18408cf0994ee07350b3de4e1a5f1b": {
            "imageSizeBytes": "27734193",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.2-beta.0"
            ],
            "timeCreatedMs": "1476125949274",
            "timeUploadedMs": "1476126971619"
        },
        "sha256:d5b9f663a00615de7b9c1084586a4b83f40343cee86699204e0989b49a9ab459": {
            "imageSizeBytes": "32903620",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.1"
            ],
            "timeCreatedMs": "1523546218259",
            "timeUploadedMs": "1523555671453"
        },
        "sha256:d74d0d82b62ed9c13926fa34abbe1875db889c23078f69ac8c66c96d558066ad": {
            "imageSizeBytes": "27711122",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.7-beta.0"
            ],
            "timeCreatedMs": "1478929630938",
            "timeUploadedMs": "1478930722281"
        },
        "sha256:d77fa71862a3fcd412c20fb4824eed79409be586cb7412f208531ee3891681f7": {
            "imageSizeBytes": "27660018",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-alpha.2"
            ],
            "timeCreatedMs": "1499896046615",
            "timeUploadedMs": "1499911070668"
        },
        "sha256:d8b5693fa16966b6dadb3b61356ae0cc68d2149a9b550b7a9357c7a9f4f3fdfc": {
            "imageSizeBytes": "20611747",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.9"
            ],
            "timeCreatedMs": "1476997423682",
            "timeUploadedMs": "1477431696823"
        },
        "sha256:d9b329fbf64e596b7c7688e034fc01f879b6dd5848c6a4c4d2caa43499b052dc": {
            "imageSizeBytes": "27553528",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.1-beta.0"
            ],
            "timeCreatedMs": "1490723147277",
            "timeUploadedMs": "1490728844317"
        },
        "sha256:dc45941c443569b06d48e511efff760f730e65e8946a439b618751bae65595fb": {
            "imageSizeBytes": "29001205",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.0-alpha.3"
            ],
            "timeCreatedMs": "1494003924328",
            "timeUploadedMs": "1494008035289"
        },
        "sha256:dcfa2e25acf36b5664781a5dcefce0963dbe008c383252417e2a8d77a8576ae8": {
            "imageSizeBytes": "32166317",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.15"
            ],
            "timeCreatedMs": "1521470759036",
            "timeUploadedMs": "1521490403500"
        },
        "sha256:de1f7c8f249dadea592f9ed07c2f76408273572024c5ae46bbf3ef1477a67be8": {
            "imageSizeBytes": "35019709",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.1-beta.0"
            ],
            "timeCreatedMs": "1530133805219",
            "timeUploadedMs": "1530136389402"
        },
        "sha256:e10e89ef0d1872bf74217eb2c4f71f7c4f07425703f9cf1ee0e5607f4a86996d": {
            "imageSizeBytes": "31726868",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.6-beta.0"
            ],
            "timeCreatedMs": "1504179465331",
            "timeUploadedMs": "1504185597232"
        },
        "sha256:e1b34e1212bf9811d9e2672054dac592d0e0738711ced494ba571dbef8a522ab": {
            "imageSizeBytes": "13038644",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.2.3"
            ],
            "timeCreatedMs": "1461360467701",
            "timeUploadedMs": "1461365260063"
        },
        "sha256:e1f4066f2f8e0cf975d8427842d95c86fd872c239e9ed2534f83121ebdce10a1": {
            "imageSizeBytes": "29778492",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.1"
            ],
            "timeCreatedMs": "1507765097592",
            "timeUploadedMs": "1507768326435"
        },
        "sha256:e3b3cf4c28c1c0569d0a1f749cbba154d7ba398b1ccedaa8f6bd849004e4c205": {
            "imageSizeBytes": "27520788",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-beta.3"
            ],
            "timeCreatedMs": "1489206631309",
            "timeUploadedMs": "1489209051872"
        },
        "sha256:e4edd66d8b61cd6268c56a0ecf5f46c0b35884f0998133c57c51750772da31b7": {
            "imageSizeBytes": "22680451",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.2-beta.0"
            ],
            "timeCreatedMs": "1481681360015",
            "timeUploadedMs": "1481684613410"
        },
        "sha256:e5a6365983c6a83a012bf9c50dedb741a8ba6f23f4360a648f0648ca3996c897": {
            "imageSizeBytes": "31140574",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-beta.2"
            ],
            "timeCreatedMs": "1512679231442",
            "timeUploadedMs": "1512683418032"
        },
        "sha256:e6bd902e530a30816b4a8750fdfbfc5f7df9ed894867ed9a9f8fe38c14f01665": {
            "imageSizeBytes": "23569882",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.7-beta.0"
            ],
            "timeCreatedMs": "1490714553620",
            "timeUploadedMs": "1490717341779"
        },
        "sha256:e7377096f0b88b0fcc5dce1c56aed002f999f095a30676c68b8f686a6bb1e943": {
            "imageSizeBytes": "29773128",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0"
            ],
            "timeCreatedMs": "1506640228986",
            "timeUploadedMs": "1506657414657"
        },
        "sha256:e882c6bb6fe69eb2509a5bec3157a46342e78f7d0461e990a9b9735dcabe1d53": {
            "imageSizeBytes": "35021623",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.11.0-rc.1"
            ],
            "timeCreatedMs": "1529540116395",
            "timeUploadedMs": "1529541765559"
        },
        "sha256:e983aa4f0eb3e456242fb8565c57a4cd0903d1d240035ae5060a3bae215597d7": {
            "imageSizeBytes": "29284512",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.0-alpha.0"
            ],
            "timeCreatedMs": "1504309959154",
            "timeUploadedMs": "1504320543104"
        },
        "sha256:ea43a1619fc4473f8454cbce8d56502ba5147650d6b7a004bc895b21bcc621ca": {
            "imageSizeBytes": "29797723",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.6-beta.0"
            ],
            "timeCreatedMs": "1512666066499",
            "timeUploadedMs": "1512669500952"
        },
        "sha256:ea916a285c82563dcd528328255b757b9d02a28f844e253b933a9068040d510a": {
            "imageSizeBytes": "20611748",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.11-beta.0"
            ],
            "timeCreatedMs": "1477951238183",
            "timeUploadedMs": "1477953332338"
        },
        "sha256:eaea6b68ea07e9eeff26a660a41a537f5b0339d6945dc4682e59eae6b59e2f6d": {
            "imageSizeBytes": "27724146",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.12-beta.0"
            ],
            "timeCreatedMs": "1492723525500",
            "timeUploadedMs": "1492724480860"
        },
        "sha256:ec9a14838f9647666f49d2f91d9fe691e0123a8d62aea5fd9fd0b02c7eafa25a": {
            "imageSizeBytes": "31839965",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.11-beta.0"
            ],
            "timeCreatedMs": "1533233013347",
            "timeUploadedMs": "1533259632454"
        },
        "sha256:ecb7588783553dd29de32bcf5b5fae8ad6b531188061eb36603569acf58874e5": {
            "imageSizeBytes": "32910415",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.10.0-rc.1"
            ],
            "timeCreatedMs": "1521505113210",
            "timeUploadedMs": "1521506717883"
        },
        "sha256:ed117d618b49663e48c0579e447c529c9aaec4bef31c86a3c6f033211f89131b": {
            "imageSizeBytes": "31668842",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.4"
            ],
            "timeCreatedMs": "1520873953859",
            "timeUploadedMs": "1520879729011"
        },
        "sha256:ed34fab6644b581adfc690ac7b2f80f0d9c295f51dc668ff76c63d278174a478": {
            "imageSizeBytes": "31743301",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.8"
            ],
            "timeCreatedMs": "1507187900419",
            "timeUploadedMs": "1507202305244"
        },
        "sha256:ed47bbedfae946c788136002cc0a67c4da6aa0c6e41135b14952782feb00671d": {
            "imageSizeBytes": "22680443",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.1-beta.0"
            ],
            "timeCreatedMs": "1481590052520",
            "timeUploadedMs": "1481593308745"
        },
        "sha256:edae874cf6e74bcbfac71abd63b38a222157176a8a01b18ac646467aedf28a1f": {
            "imageSizeBytes": "27740304",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.5-beta.0"
            ],
            "timeCreatedMs": "1477019585321",
            "timeUploadedMs": "1477020667876"
        },
        "sha256:ee3ea5fd68ae18203543722e5a7aaed0e155f57d89f08811f908a5fbd0426d49": {
            "imageSizeBytes": "26795657",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.0-beta.0"
            ],
            "timeCreatedMs": "1487738226595",
            "timeUploadedMs": "1487741821688"
        },
        "sha256:eec4329de0892f4a960b7f1202272f93880d3071a9b40d8407585125b37d527d": {
            "imageSizeBytes": "31547620",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.2"
            ],
            "timeCreatedMs": "1516272351902",
            "timeUploadedMs": "1516304932492"
        },
        "sha256:eff9558dca7077fe9b9b43c2532aa0e2022d3b74f7575d4ffb27b6c21d4a8314": {
            "imageSizeBytes": "27759946",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.1-beta.0"
            ],
            "timeCreatedMs": "1474915353826",
            "timeUploadedMs": "1474916372655"
        },
        "sha256:f03cc10a8f08d2b7e380006d1a530c5b1481cf1f14b402384fceec5b5aa67574": {
            "imageSizeBytes": "20194726",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.3.0-alpha.5"
            ],
            "timeCreatedMs": "1464900615339",
            "timeUploadedMs": "1464903358163"
        },
        "sha256:f0658f0c3344abaae3217d6e08612ea8816a0cdeea1fdd48776dd55d07f17d30": {
            "imageSizeBytes": "29791838",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.6"
            ],
            "timeCreatedMs": "1513838698883",
            "timeUploadedMs": "1513878873546"
        },
        "sha256:f09dcd9825d35b07273f7611c24a5f4fee986b40f0be3126a9d4ef11badab019": {
            "imageSizeBytes": "22680358",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.1"
            ],
            "timeCreatedMs": "1481677590079",
            "timeUploadedMs": "1481684065069"
        },
        "sha256:f1dbcd026e8577226cbb1a83c3e6a5ff35bb0ccbeb1973c00aff60b7a42b4b40": {
            "imageSizeBytes": "25323736",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.5.0-alpha.1"
            ],
            "timeCreatedMs": "1476390534963",
            "timeUploadedMs": "1476476896357"
        },
        "sha256:f2acebc4df545cafea168e9fd3ec39c46ecdd9e4532cba192e1fd389d074f730": {
            "imageSizeBytes": "22812554",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.5.3"
            ],
            "timeCreatedMs": "1487141479702",
            "timeUploadedMs": "1487148748353"
        },
        "sha256:f3a208d30314a89952cf613e5ee671f9d2ed7b197cd6c5d91bebfe02571d7e1b": {
            "imageSizeBytes": "31736733",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.6"
            ],
            "timeCreatedMs": "1505373547240",
            "timeUploadedMs": "1505386791040"
        },
        "sha256:f474819f3ebf18a064260e86fdca04f56a744db5c0d29741bc1bc461b6d5f223": {
            "imageSizeBytes": "29792849",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.4"
            ],
            "timeCreatedMs": "1511156325367",
            "timeUploadedMs": "1511203118031"
        },
        "sha256:f503e96d82727f1126e87e9815d5ae60a7c69f3d788bf8a7f54897c49802b96c": {
            "imageSizeBytes": "22809664",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [],
            "timeCreatedMs": "1484201682378",
            "timeUploadedMs": "1484205665679"
        },
        "sha256:f70af327840f10eb0b47b95a9848593958ead9d875f7a01bd7287fc70478e069": {
            "imageSizeBytes": "31529661",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.1"
            ],
            "timeCreatedMs": "1515067466167",
            "timeUploadedMs": "1515092202868"
        },
        "sha256:f880371b4cee1a810d7caf4c8a2c0b8fa169879545b06a537e4ea6bcdfbbe1f6": {
            "imageSizeBytes": "31722624",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.4"
            ],
            "timeCreatedMs": "1502961072958",
            "timeUploadedMs": "1502988258616"
        },
        "sha256:f9b974a6d0986345b1d7b35d91f47ae7959143e13a0d46da0bbb308ed8c887b5": {
            "imageSizeBytes": "31737208",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.8-beta.0"
            ],
            "timeCreatedMs": "1506650496250",
            "timeUploadedMs": "1506665052827"
        },
        "sha256:fa69448be8ba3a0465da9d2d3bc845186f7b06c80a6a2dd29539ab7e31250069": {
            "imageSizeBytes": "31529704",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.9.2-beta.0"
            ],
            "timeCreatedMs": "1515069851019",
            "timeUploadedMs": "1515092849917"
        },
        "sha256:fb60b2e5191c58358101b9e4aebdab179424361f2513f89deb161e49e76d5815": {
            "imageSizeBytes": "27759948",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v1+prettyjws",
            "tag": [
                "v1.4.0-beta.10"
            ],
            "timeCreatedMs": "1474521427325",
            "timeUploadedMs": "1474526252806"
        },
        "sha256:fc0730e9646cb34458f511763a694e5e6f0e32e5c79c57a06ba12da19eb7a9a1": {
            "imageSizeBytes": "29708440",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.8.0-beta.1"
            ],
            "timeCreatedMs": "1504835258656",
            "timeUploadedMs": "1504838091780"
        },
        "sha256:fc30770ff400149229266f99061635048f2cd70e7f5a02bcdf0c31e29eed9eed": {
            "imageSizeBytes": "32168276",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.15-beta.0"
            ],
            "timeCreatedMs": "1520875777823",
            "timeUploadedMs": "1520882055875"
        },
        "sha256:fcf0d14b84f1b15f1738cfddb8fa692a23e60964a24d1938e869cd540c7b84fc": {
            "imageSizeBytes": "31739816",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.11-beta.0"
            ],
            "timeCreatedMs": "1509738877917",
            "timeUploadedMs": "1509743102431"
        },
        "sha256:fea47371f2fe933cdffa80e6f33b2e52038e20acaa60f1bf00d5d1cc84569356": {
            "imageSizeBytes": "27598359",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.6.6-beta.0"
            ],
            "timeCreatedMs": "1497475506268",
            "timeUploadedMs": "1497478712522"
        },
        "sha256:febc69a8236f09c6776e50a383a70ba89a63b26e645c9e63f7e909b70a138806": {
            "imageSizeBytes": "32165314",
            "layerId": "",
            "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
            "tag": [
                "v1.7.17-beta.0"
            ],
            "timeCreatedMs": "1522836459544",
            "timeUploadedMs": "1522848006462"
        }
    },
    "name": "google-containers/kube-apiserver",
    "tags": [
        "9680e782e08a1a1c94c656190011bd02",
        "v1.10.0",
        "v1.10.0-alpha.0",
        "v1.10.0-alpha.1",
        "v1.10.0-alpha.2",
        "v1.10.0-alpha.3",
        "v1.10.0-beta.0",
        "v1.10.0-beta.1",
        "v1.10.0-beta.2",
        "v1.10.0-beta.3",
        "v1.10.0-beta.4",
        "v1.10.0-rc.1",
        "v1.10.1",
        "v1.10.1-beta.0",
        "v1.10.2",
        "v1.10.2-beta.0",
        "v1.10.3",
        "v1.10.3-beta.0",
        "v1.10.4",
        "v1.10.4-beta.0",
        "v1.10.5",
        "v1.10.5-beta.0",
        "v1.10.6",
        "v1.10.6-beta.0",
        "v1.10.7-beta.0",
        "v1.11.0",
        "v1.11.0-alpha.0",
        "v1.11.0-alpha.1",
        "v1.11.0-alpha.2",
        "v1.11.0-beta.0",
        "v1.11.0-beta.1",
        "v1.11.0-beta.2",
        "v1.11.0-rc.1",
        "v1.11.0-rc.2",
        "v1.11.0-rc.3",
        "v1.11.1",
        "v1.11.1-beta.0",
        "v1.11.2-beta.0",
        "v1.12.0-alpha.0",
        "v1.12.0-alpha.1",
        "v1.2.0",
        "v1.2.0-alpha.8",
        "v1.2.0-beta.0",
        "v1.2.0-beta.1",
        "v1.2.1",
        "v1.2.1-beta.0",
        "v1.2.2",
        "v1.2.2-beta.0",
        "v1.2.3",
        "v1.2.3-beta.0",
        "v1.2.4",
        "v1.2.4-beta.0",
        "v1.2.5",
        "v1.2.5-beta.0",
        "v1.2.6",
        "v1.2.6-beta.0",
        "v1.2.7",
        "v1.2.7-beta.0",
        "v1.2.8-beta.0",
        "v1.3.0",
        "v1.3.0-alpha.0",
        "v1.3.0-alpha.1",
        "v1.3.0-alpha.2",
        "v1.3.0-alpha.3",
        "v1.3.0-alpha.4",
        "v1.3.0-alpha.5",
        "v1.3.0-beta.0",
        "v1.3.0-beta.1",
        "v1.3.0-beta.2",
        "v1.3.0-beta.3",
        "v1.3.1",
        "v1.3.1-beta.0",
        "v1.3.1-beta.1",
        "v1.3.10",
        "v1.3.10-beta.0",
        "v1.3.11-beta.0",
        "v1.3.2",
        "v1.3.2-beta.0",
        "v1.3.3",
        "v1.3.3-beta.0",
        "v1.3.4",
        "v1.3.4-beta.0",
        "v1.3.5",
        "v1.3.5-beta.0",
        "v1.3.6",
        "v1.3.6-beta.0",
        "v1.3.7",
        "v1.3.7-beta.0",
        "v1.3.8",
        "v1.3.8-beta.0",
        "v1.3.9",
        "v1.4.0",
        "v1.4.0-alpha.0",
        "v1.4.0-alpha.1",
        "v1.4.0-alpha.2",
        "v1.4.0-alpha.3",
        "v1.4.0-beta.0",
        "v1.4.0-beta.1",
        "v1.4.0-beta.10",
        "v1.4.0-beta.11",
        "v1.4.0-beta.2",
        "v1.4.0-beta.3",
        "v1.4.0-beta.5",
        "v1.4.0-beta.6",
        "v1.4.0-beta.7",
        "v1.4.0-beta.8",
        "v1.4.1",
        "v1.4.1-beta.0",
        "v1.4.10-beta.0",
        "v1.4.12",
        "v1.4.12-beta.0",
        "v1.4.2",
        "v1.4.2-beta.0",
        "v1.4.2-beta.1",
        "v1.4.3",
        "v1.4.3-beta.0",
        "v1.4.4",
        "v1.4.4-beta.0",
        "v1.4.5",
        "v1.4.5-beta.0",
        "v1.4.6",
        "v1.4.6-beta.0",
        "v1.4.7",
        "v1.4.7-beta.0",
        "v1.4.8",
        "v1.4.8-beta.0",
        "v1.4.9",
        "v1.4.9-beta.0",
        "v1.5.0",
        "v1.5.0-alpha.0",
        "v1.5.0-alpha.1",
        "v1.5.0-alpha.2",
        "v1.5.0-beta.0",
        "v1.5.0-beta.1",
        "v1.5.0-beta.2",
        "v1.5.0-beta.3",
        "v1.5.1",
        "v1.5.1-beta.0",
        "v1.5.2",
        "v1.5.2-beta.0",
        "v1.5.3",
        "v1.5.3-beta.0",
        "v1.5.4",
        "v1.5.4-beta.0",
        "v1.5.5",
        "v1.5.5-beta.0",
        "v1.5.6",
        "v1.5.7",
        "v1.5.7-beta.0",
        "v1.5.8",
        "v1.5.8-beta.0",
        "v1.5.9-beta.0",
        "v1.6.0",
        "v1.6.0-alpha.0",
        "v1.6.0-alpha.1",
        "v1.6.0-alpha.2",
        "v1.6.0-alpha.3",
        "v1.6.0-beta.0",
        "v1.6.0-beta.1",
        "v1.6.0-beta.2",
        "v1.6.0-beta.3",
        "v1.6.0-beta.4",
        "v1.6.0-rc.1",
        "v1.6.1",
        "v1.6.1-beta.0",
        "v1.6.10",
        "v1.6.10-beta.0",
        "v1.6.11",
        "v1.6.11-beta.0",
        "v1.6.12",
        "v1.6.12-beta.0",
        "v1.6.13",
        "v1.6.13-beta.0",
        "v1.6.14-beta.0",
        "v1.6.2",
        "v1.6.2-beta.0",
        "v1.6.3",
        "v1.6.3-beta.0",
        "v1.6.3-beta.1",
        "v1.6.4",
        "v1.6.4-beta.0",
        "v1.6.5",
        "v1.6.5-beta.0",
        "v1.6.6",
        "v1.6.6-beta.0",
        "v1.6.7",
        "v1.6.7-beta.0",
        "v1.6.8",
        "v1.6.8-beta.0",
        "v1.6.9",
        "v1.6.9-beta.0",
        "v1.7.0",
        "v1.7.0-alpha.0",
        "v1.7.0-alpha.1",
        "v1.7.0-alpha.2",
        "v1.7.0-alpha.3",
        "v1.7.0-alpha.4",
        "v1.7.0-beta.0",
        "v1.7.0-beta.1",
        "v1.7.0-beta.2",
        "v1.7.0-rc.1",
        "v1.7.1",
        "v1.7.1-beta.0",
        "v1.7.10",
        "v1.7.10-beta.0",
        "v1.7.11",
        "v1.7.11-beta.0",
        "v1.7.12",
        "v1.7.12-beta.0",
        "v1.7.13",
        "v1.7.13-beta.0",
        "v1.7.14",
        "v1.7.14-beta.0",
        "v1.7.15",
        "v1.7.15-beta.0",
        "v1.7.16",
        "v1.7.16-beta.0",
        "v1.7.17-beta.0",
        "v1.7.2",
        "v1.7.2-beta.0",
        "v1.7.3",
        "v1.7.3-beta.0",
        "v1.7.4",
        "v1.7.4-beta.0",
        "v1.7.5",
        "v1.7.5-beta.0",
        "v1.7.6",
        "v1.7.6-beta.0",
        "v1.7.7",
        "v1.7.7-beta.0",
        "v1.7.8",
        "v1.7.8-beta.0",
        "v1.7.9",
        "v1.7.9-beta.0",
        "v1.8.0",
        "v1.8.0-alpha.0",
        "v1.8.0-alpha.1",
        "v1.8.0-alpha.2",
        "v1.8.0-alpha.3",
        "v1.8.0-beta.0",
        "v1.8.0-beta.1",
        "v1.8.0-rc.1",
        "v1.8.1",
        "v1.8.1-beta.0",
        "v1.8.10",
        "v1.8.10-beta.0",
        "v1.8.11",
        "v1.8.11-beta.0",
        "v1.8.12",
        "v1.8.12-beta.0",
        "v1.8.13",
        "v1.8.13-beta.0",
        "v1.8.14",
        "v1.8.14-beta.0",
        "v1.8.15",
        "v1.8.15-beta.0",
        "v1.8.16-beta.0",
        "v1.8.2",
        "v1.8.2-beta.0",
        "v1.8.3",
        "v1.8.3-beta.0",
        "v1.8.4",
        "v1.8.4-beta.0",
        "v1.8.5",
        "v1.8.5-beta.0",
        "v1.8.6",
        "v1.8.6-beta.0",
        "v1.8.7",
        "v1.8.7-beta.0",
        "v1.8.8",
        "v1.8.8-beta.0",
        "v1.8.9",
        "v1.8.9-beta.0",
        "v1.9.0",
        "v1.9.0-alpha.0",
        "v1.9.0-alpha.1",
        "v1.9.0-alpha.2",
        "v1.9.0-alpha.3",
        "v1.9.0-beta.0",
        "v1.9.0-beta.1",
        "v1.9.0-beta.2",
        "v1.9.1",
        "v1.9.1-beta.0",
        "v1.9.10",
        "v1.9.10-beta.0",
        "v1.9.11-beta.0",
        "v1.9.2",
        "v1.9.2-beta.0",
        "v1.9.3",
        "v1.9.3-beta.0",
        "v1.9.4",
        "v1.9.4-beta.0",
        "v1.9.5",
        "v1.9.5-beta.0",
        "v1.9.6",
        "v1.9.6-beta.0",
        "v1.9.7",
        "v1.9.7-beta.0",
        "v1.9.8",
        "v1.9.8-beta.0",
        "v1.9.9",
        "v1.9.9-beta.0"
    ]
}`

var quayEtcdResp = `{
    "name": "coreos/etcd",
    "tags": [
        "v3.3.7",
        "v3.3.7-arm64",
        "v3.3.7-ppc64le",
        "v3.1.17",
        "v3.2.22",
        "v3.2.22-arm64",
        "v3.2.22-ppc64le",
        "v3.2",
        "v3.2.23",
        "v3.2.23-arm64",
        "v3.2.23-ppc64le",
        "v3.1.18",
        "v3.3.8",
        "v3.3.8-arm64",
        "v3.3.8-ppc64le",
        "latest",
        "v3.1",
        "v3.1.24",
        "v3.1.19",
        "v3.3.9",
        "v3.3.9-arm64",
        "v3.3.9-ppc64le",
        "v3.3",
        "v3.2.24",
        "v3.2.24-arm64",
        "v3.2.24-ppc64le"
    ]
}`
