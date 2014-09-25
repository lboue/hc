package pair

import (
    "github.com/brutella/hap"
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestPairVerifyIntegration(t *testing.T) {
    accessory, err := hap.NewAccessory("HAP Test", "123-45-678")
    assert.Nil(t, err)
    
    storage, err := hap.NewFileStorage(os.TempDir())
    assert.Nil(t, err)
    context := hap.NewContext(storage)
    controller, err := NewVerifyServerController(context, accessory)
    assert.Nil(t, err)
    
    name := "UnitTest"
    client_controller := NewVerifyClientController(context, accessory, name)
    context.SaveClient(hap.NewClient(name,client_controller.LTPK)) // make LTPK available for server
    
    tlvVerifyStartRequest := client_controller.InitialKeyVerifyRequest()
    // 1) C -> S
    tlvVerifyStartRespond, err := controller.Handle(tlvVerifyStartRequest)
    assert.Nil(t, err)
    // 2) S -> C
    tlvFinishRequest, err := client_controller.Handle(tlvVerifyStartRespond)
    assert.Nil(t, err)
    
    tlvFinishRespond, err := controller.Handle(tlvFinishRequest)
    assert.Nil(t, err)
    
    response, err := client_controller.Handle(tlvFinishRespond)
    assert.Nil(t, err)
    assert.Nil(t, response)
} 