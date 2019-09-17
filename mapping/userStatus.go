package mapping

type userStatus struct {
	Disabled mapping
	Enabled  mapping
}

var (
	UserStatus userStatus
)

func init() {
	UserStatus = userStatus{
		Disabled: mapping{
			Value: "-1",
			Label: "disabled",
		},
		Enabled: mapping{
			Value: "1",
			Label: "enabled",
		},
	}
}
