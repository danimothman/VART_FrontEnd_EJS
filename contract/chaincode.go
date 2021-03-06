package main



import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim" 
	//shim library를 사용하겠다~~ >> DB(tx)에 access하고 체인코드를 호출하는 API를 제공
  sc "github.com/hyperledger/fabric/protos/peer" // 여기서 hyperledger에서 peer와 헷갈리기 때문에 이름을 바꿔주기 위해 sc 명령어를 사용한다 
)

// Define the Smart Contract structure
type Chaincode  struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Publicinformation struct {                            //공시정보 구조체
	Companyname       string  `json:"companyname"`         //회사이름
	Establishment     string  `json:"establishment"`       //설립일
	Location          string  `json:"location"`            //위치
	Statejurisdiction string  `json:"statejurisdiction"`   //법인관할자
	
	Tokenprofile struct {                                  //토큰정보 구조체
     Tokenname         string `json:"tokenname"`           //토큰 이름
	 Projecttype       string `json:"projecttype"`         //프로젝트 종류
 	}
    Executives struct {                                    //경영진 구조체
	 Name      	string 		`json:"name"`                  //이름
	 Education 	string      `json:"education"`             //학력
	 Experience string      `json:"experience"`            //경력
	} 
	developerleaders struct {                              //개발자리더 구조체
     Name      	string 		`json:"name"`
	 Education 	string      `json:"education"`
	 Experience string      `json:"experience"`
	}

 }

 
// type Market struct {
// 	Mkname       string `json:"Mkname"`
// 	Mklocation   string `json:"Mklocation"`
// 	MKcpdate     int    `json:"MKcpdate"`
//     MKfounder    string `json:"MKfounder"`

	
// }
// type onchain struct {
//	Milestone    string `json:"milestone"`
//	Move0000000000num      string `json:"movenum"`
//	Wallet       string `json:"wallet"`
   
// }



func (s *Chaincode) Init(APIstub shim.ChaincodeStubInterface) sc.Response { //sc >> Import 된 peer와 CLI peer를 구분하기 위해 사용
	//Init >> chaincode를 인스턴스화, 업그레이드 시킬때 자동으로 실행되는 함수 >> 생성자!!
	//인자로 shim라이브러리의 ChaincodeStubInterface를 매개변수를 사용한다
	/*stub >> 블록체인에 들어있는 Ledger에 접근할때 사용하는 매개체
	rpc개념과 비슷!! (내 시스템 안에 있는 func을 호출하는게 아니라 다른 시스템 안에 있는 fuc을 호출할때는 Remote Procedure Call이라는 형태로 call한다!! 
	내가 call할 때 func을 호출하는 형태와 네트워크를 타고 목표하는 fuc을 호출하는 형태가 일관되게 유지해야한다!!
	but 네트워크를 타고갈때는 data형태로 이동을 한다!! 내 시스템안에서 func call을 하면 func 주소를 가지고 바로 호출하는데 네트워크를 타고 다른 시스템에 가게 되면
	그것을 Serialize를 통해 문자형태로 바꾸어서 네트워크 형태로 보내야 한다!!
    Stub >> 내 시스템에 내에 있는, or 다른 시스템 내에 있는 특정한 형태의 func을 연결해주는 매개체

	*/


	//Shim에 있는 interface를 사용하는데 이름이 기니까 APIstub라는 이름으로 사용하겠다
	// shim.ChaincodeStubInterface >> 원장에 대한 접근 수정을 위한 interface >> get, put과 같은 method가 내장.
	return shim.Success(nil)
} 
func (s *Chaincode) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters() 
	/*stub.getFunc~ 메소드 >>peer가 chaincode를 인스턴스화 시키면서 Init을 수행시키면 
	 
	 >> web브라우저에서 user가 인자값을 넣고 function을 호출시켜 tx를 발생시켰을때 HyperLedger 블록체인에서 어떤게 호출되는 함수이고 어떤게 인자인지 구별하기 위해 사용!!*/
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "readPublicinfo" { // function이 querycar이냐?
		return s.readPublicinfo(APIstub, args) //querycar 함수에 인자를 넣으면서 호출한다
	} else if function == "initLedgerpubilcinfo" {
		return s.initLedgerPubilcinfo(APIstub)
	} else if function == "addPublicinfoinfo" {
		return s.addPublicinfoinfo(APIstub, args)
	} else if function == "readAllPublicinfo" {
		return s.readAllPublicinfo(APIstub)
	} else if function == "updatePublicinfo" {
		return s.updatePublicinfo(APIstub, args)
	} // else if function == "createinformation" {
	//	return s.createOnchain(APIstub, args)
	

	return shim.Error("Invalid Smart Contract function name.")
}
func (s *Chaincode) readPublicinfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	// queryCar가 interface를 사용하게따
	// queryCar을 하려면 어떤 car인지 인자가 필요하다!!
	// sc.Response >> sc : peer 라이브러리 내에서 Response를 쓰는데 Response 내에서는 반환하는 값들이 저장되어있다
	
		if len(args) != 1 { //인자의 길이가 0이면 에러를 발생시킨다
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}
		
	
		PublicinformationAsBytes, _ := APIstub.GetState(args[0]) //GetState >> stateDB에서 인자값에 해당하는 내용을 가져온다 >> 만약 원장에 commit 되지 않은 writeset 데이터는 읽지 않는다
		// _ >> carAsBytes를 반환할수 없으면 에러를 반환한다
		return shim.Success(PublicinformationAsBytes) 
		//shim 라이브러리 내 Success 함수 >> stateDB에 잘 update가 되었다~ >> 성공상태 정보, 바이트 형태의 페이로드 데이터(user가 누군지, car는 어떤 차인지)를 반환 >> 여기서는 carAsBytes를 반환
	}
	
func (s *Chaincode) initLedgerpubilcinfo(APIstub shim.ChaincodeStubInterface) sc.Response {
	Publicinformations := []Publicinformation{
			
		Publicinformation{
			Companyname: "이더리움",
			Establishment:"2014-02-15",
			Location:"대전광역시 서구 탄방동", 
			Statejurisdiction:"korean",
			Tokenprofile:  struct{
				Tokenname         string `json:"tokenname"`
				Projecttype       string `json:"projecttype"`
			} {	
				Tokenname:"이더리움",
				Projecttype:"플랫폼",
			  },
		    Executives: struct{
			 Name      	string       `json:"name"`
			 Education 	string       `json:"education"`
			 Experience string       `json:"experience"`
			} {
				Name:"leeSangho",      	
		 		Education:"Kumoh National Institute of Technology",
				Experience:"LG display",
			 },

			developerleaders: struct {
				Name      	string 		`json:"name"`
	 			Education 	string      `json:"education"`
				Experience string      `json:"experience"`
			} {
				Name:"kim",      	
				Education:"Seoul National Institute of Technology",
				Experience:"korea display",
			  },
		},

		Publicinformation{
			Companyname: "파트너스",
			Establishment:"2019-03-14",
			Location:"대전광역시동구 가양동", 
			Statejurisdiction:"korean",

				Tokenprofile:  struct{
				Tokenname         string `json:"tokenname"`
				Projecttype       string `json:"projecttype"`
			} {
				Tokenname:"파트",
				Projecttype:"유틸리티토큰",
			  },

		    Executives: struct{
			 Name      	string       `json:"name"`
			 Education 	string       `json:"education"`
			 Experience string       `json:"experience"`
			} {
				Name:"홍길동",      	
		 		Education:" Kyungpook National University",
				Experience:"IBM",
			 },

			developerleaders: struct {
				Name      	string 		`json:"name"`
	 			Education 	string      `json:"education"`
				Experience  string      `json:"experience"`
			} {
				Name:"phone",      	
				Education:" Pusan National University",
				Experience:"LA display",
			  },
			},
	}
	
		i := 0
		for i < len(Publicinformations) {
			fmt.Println("i is ", i)
			PublicinformationAsBytes, _ := json.Marshal(Publicinformations[i]) //network 형태로 전송시키려면 car의 구조체를 bytecode(JSON)로 변환시켜줘야 한다!!
												   // '_'를 써서 에러는 처리하지 않음
			APIstub.PutState("Publicinfo"+strconv.Itoa(i), PublicinformationAsBytes)
			fmt.Println("Added", Publicinformations[i])
			i = i + 1
		}
	
		return shim.Success(nil)
	}
	
	func (s *Chaincode) addcompanyinfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
		if len(args) != 13 {
			return shim.Error("Incorrect number of arguments. Expecting 5")
		}
	
		var Publicinformation = Publicinformation{
			 Companyname: args[1],
			 Establishment:args[2],
			 Location:args[3],
			 Statejurisdiction:args[4],

		        Tokenprofile:  struct{
				Tokenname         string `json:"tokenname"`
				Projecttype       string `json:"projecttype"`
			} {
				Tokenname:args[5], 
				Projecttype:args[6],            
			},
			Executives: struct{
				Name      	string       `json:"name"`
				Education 	string       `json:"education"`
				Experience  string       `json:"experience"`
			   } {
				   Name:args[7],     	
				   Education:args[8],
				   Experience:args[9],
				},
			developerleaders: struct {
			     Name         	string 		`json:"name"`
		         Education 	    string      `json:"education"`
				 Experience     string      `json:"experience"`
				} {
				 Name:args[10],      	
				 Education:args[11],
				 Experience:args[12],
				  },
		}
		PublicinformationAsBytes, _ := json.Marshal(Publicinformation)
		APIstub.PutState(args[0], PublicinformationAsBytes) 
		// PutState가 발생되면 transaction이 일어나게 된다 >> 지정된 key와 value를 트랜잭션의 writeset에 data-write proposal 수행(일단 tx를 endorsor peer한테까지만 전달된 초기상태인 상황이다!!) >> 데이터 추가에 대한 요청(proposal)들만 수행 >> proposal들이 모여서 일정시간 검증되면 새로운 block이 생성된다 >> 또 다른 과정!!
	
		return shim.Success(nil)
	}
	
	func (s *Chaincode) readAllPublicinfo(APIstub shim.ChaincodeStubInterface) sc.Response {
	
		startKey := "Publicinfo0"
		endKey   := "Publicinfo999"
	
		resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
		//GetStateByRange >> 원장에서 data를 읽어올때 더미로 읽어들이고 싶을때
		if err != nil {
			return shim.Error(err.Error())
		}
		defer resultsIterator.Close() //finally 블럭처럼 마지막에 Clean-up 작업을 하기 위해 사용
	
		// buffer is a JSON array containing QueryResults
		var buffer bytes.Buffer
		buffer.WriteString("[")
	
		bArrayMemberAlreadyWritten := false
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			// Add a comma before array members, suppress it for the first array member
			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString("\"")
	
			buffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			buffer.WriteString(string(queryResponse.Value))
			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true
		}
		buffer.WriteString("]")
	
		fmt.Printf("- queryAllinfo:\n%s\n", buffer.String())
	
		return shim.Success(buffer.Bytes())
	}
	

	func (s *Chaincode) updatePublicinfo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
		if len(args) != 7 {
			return shim.Error("Incorrect number of arguments. Expecting 2")
		}
	
		PublicinformationAsBytes, _ := APIstub.GetState(args[0]) // 인자인 key값(args[0])에 대한 stateDB 내의 date를 보여준다
	    Publicinformations := Publicinformation{}
		json.Unmarshal(PublicinformationAsBytes, &Publicinformations) // 전달받은 carAsBytes라는 JSON 형식의 데이터를 car 구조체 안의 값으로 집어넣음
		Publicinformations.Companyname		 = args[1]
		Publicinformations.Establishment	 = args[2]
		Publicinformations.Location     	 = args[3]
		Publicinformations.Statejurisdiction = args[4]
		Publicinformations.Tokenprofile.Tokenname	 = args[5]
		Publicinformations.Tokenprofile.Projecttype   = args[6]
		
	
		PublicinformationAsBytes, _ = json.Marshal(Publicinformations)
		APIstub.PutState(args[0], PublicinformationAsBytes)
	
		return shim.Success(nil)
	}
    

	// The main tion is only relevant in unit test mode. Only included here for completeness.
	func main() { 
	
		// Create a new Smart Contract
		err := shim.Start(new(Chaincode)) //shim.start >> 스마트 컨트랙트 생성
		if err != nil { //스마트 컨트랙트를 발생시켰을때 에러가 발생되었으면 출력하겠다
			fmt.Printf("Error creating new Smart Contract: %s", err)
		}
	}
