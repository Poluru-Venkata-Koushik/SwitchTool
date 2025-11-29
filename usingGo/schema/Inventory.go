package schema

type Device struct{
	Hostname string `json:"Hostname" bson:"Hostname"`
	IP 		 string `json:"IP" bson:"IP"`
	Model 	 string `json:"Model" bson:"Model"`
	ID 		 string `json:"ID" bson:"_id"`

}
