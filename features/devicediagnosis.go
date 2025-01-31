package features

import (
	"github.com/enbility/eebus-go/spine"
	"github.com/enbility/eebus-go/spine/model"
)

type DeviceDiagnosis struct {
	*FeatureImpl
}

func NewDeviceDiagnosis(localRole, remoteRole model.RoleType, spineLocalDevice *spine.DeviceLocalImpl, entity *spine.EntityRemoteImpl) (*DeviceDiagnosis, error) {
	feature, err := NewFeatureImpl(model.FeatureTypeTypeDeviceDiagnosis, localRole, remoteRole, spineLocalDevice, entity)
	if err != nil {
		return nil, err
	}

	dd := &DeviceDiagnosis{
		FeatureImpl: feature,
	}

	return dd, nil
}

// request DeviceDiagnosisStateData from a remote entity
func (d *DeviceDiagnosis) RequestState() (*model.MsgCounterType, error) {
	return d.requestData(model.FunctionTypeDeviceDiagnosisStateData, nil, nil)
}

// get the current diagnosis state for an device entity
func (d *DeviceDiagnosis) GetState() (*model.DeviceDiagnosisStateDataType, error) {
	rData := d.featureRemote.Data(model.FunctionTypeDeviceDiagnosisStateData)
	if rData == nil {
		return nil, ErrDataNotAvailable
	}

	data := rData.(*model.DeviceDiagnosisStateDataType)
	if data == nil {
		return nil, ErrDataNotAvailable
	}

	return data, nil
}

func (d *DeviceDiagnosis) SendState(operatingState *model.DeviceDiagnosisStateDataType) {
	d.featureLocal.SetData(model.FunctionTypeDeviceDiagnosisStateData, operatingState)

	_, _ = d.featureLocal.NotifyData(model.FunctionTypeDeviceDiagnosisStateData, nil, nil, false, nil, d.featureRemote)
}
