package handler

import (
	"bytes"
	"consume_api/domain"
	"consume_api/utils"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func ConsumeSoap(w http.ResponseWriter, r *http.Request) {
	var request domain.AddRequestJson
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		utils.NewResponse(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, nil).Write(w)
		return
	}

	envelope := domain.Envelope{
		XmlnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
        XmlnsXsd:  "http://www.w3.org/2001/XMLSchema",
		Body: domain.Body{
			Add: domain.AddRequest{
				Xmlns: "http://tempuri.org/",
                IntA:  request.IntA,
                IntB:  request.IntB,
			},
		},
	}

	body, err := xml.MarshalIndent(envelope, "", "	")

	if err != nil {
		utils.NewResponse(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, nil).Write(w)
		return
	}

	soapRequest := []byte(xml.Header + string(body))

	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer(soapRequest))

	if err != nil {
		utils.NewResponse(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, nil).Write(w)
		return
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
    req.Header.Set("SOAPAction", "http://tempuri.org/Add")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.NewResponse(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, nil).Write(w)
		return
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
    if err != nil {
        utils.NewResponse(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, nil).Write(w)
        return
    }

	var soapResponse domain.EnvelopeRes
    err = xml.Unmarshal(res, &soapResponse)
    if err != nil {
        utils.NewResponse(err.Error(), http.StatusInternalServerError, nil).Write(w)
        return
    }

	result := domain.AddResponseJson{
		Result: soapResponse.Body.Add.AddResult,
	}

	jsonRes, errMarshal := json.Marshal(result)

	if errMarshal != nil {
		log.Error().Msg(errMarshal.Error())
		utils.NewResponse(errMarshal.Error(), http.StatusInternalServerError, nil).Write(w)
		return
	}
	w.Write(jsonRes)
}