package logic

import "testing"

func TestQcLogic(t *testing.T) {
	qcLogic := NewQcLogic()
	var qcs []Qc
	qcs, err := qcLogic.GetQCList()
	if err != nil {
		panic(err)
	}
	qcs[2].Content = "test\nhello"
	err = qcLogic.UpdateQC(qcs[2])
	if err != nil {
		panic(err)
	}

	err = qcLogic.CreateQC(Qc{
		Name:     "test1",
		Filepath: "~/.qc/",
		Status:   true,
		Content:  "tstajta\nhfakhfa\n",
	})
	if err != nil {
		panic(err)
	}

	err = qcLogic.DeleteQC(Qc{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
}
