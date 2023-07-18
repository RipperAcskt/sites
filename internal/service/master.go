package service

import (
	"context"
	"fmt"
	"os"
	"sites/internal/service/sites"
	"sync"
	"time"
)

var (
	ErrLowBalance  = fmt.Errorf("low balance")
	ErrNoTask      = fmt.Errorf("now such task")
	StatusProgress = "in_progress"
	StatusPause    = "pause"
	StatusComplete = "complete"
	StatusDeleted  = "deleted"
)

type MasterRepo interface {
	GetBalance(ctx context.Context, userId string) (float64, error)
	SetBalance(ctx context.Context, userId string, balance float64) error
}

type Master struct {
	MasterRepo
	store map[string][]WorkInfo
}

type WorkInfo struct {
	ctx      context.Context
	cancel   context.CancelFunc
	timeLeft time.Duration
	Status   string
	Email    string
	Time     int
}

func NewMaster(repo MasterRepo) *Master {
	return &Master{repo, make(map[string][]WorkInfo)}
}

func (m *Master) CreateMaster(ctx context.Context, userId, email string, price float64, time int) error {
	balance, err := m.GetBalance(ctx, userId)
	if err != nil {
		return fmt.Errorf("get balance failed: %w", err)
	}

	if balance < price {
		return ErrLowBalance
	}

	err = m.SetBalance(ctx, userId, balance-price)
	if err != nil {
		return fmt.Errorf("set balance failed: %w", err)
	}

	cancelCtx, cancel := context.WithCancel(context.Background())
	info := WorkInfo{
		ctx:    cancelCtx,
		cancel: cancel,
		Status: StatusProgress,
		Email:  email,
		Time:   time,
	}

	_, ok := m.store[userId]
	if !ok {
		m.store[userId] = make([]WorkInfo, 0, 1)
		m.store[userId] = append(m.store[userId], info)
	} else {
		m.store[userId] = append(m.store[userId], info)
	}
	err = m.MasterProgram(&info, email, time)
	if err != nil {
		return fmt.Errorf("master program failed: %w", err)
	}
	return nil
}

func (m *Master) GetTasks(userId string, admin bool) map[string][]WorkInfo {
	info := make(map[string][]WorkInfo)
	if admin {
		return m.store
	}
	info[userId] = m.store[userId]
	return info
}

func (m *Master) StopTask(userId string, orderID int) error {
	task, ok := m.store[userId]
	if !ok {
		return ErrNoTask
	}

	if len(task) <= orderID {
		return ErrNoTask
	}

	task[orderID].Status = StatusPause
	task[orderID].cancel()

	return nil
}

func (m *Master) DeleteTask(userId string, orderID int) error {
	task, ok := m.store[userId]
	if !ok {
		return ErrNoTask
	}

	if len(task) <= orderID {
		return ErrNoTask
	}

	task[orderID].Status = StatusDeleted
	task[orderID].cancel()
	task[orderID] = task[len(task)-1]
	task[len(task)-1] = WorkInfo{}
	task = task[:len(task)-1]
	m.store[userId] = task
	return nil
}

func (m *Master) MasterProgram(info *WorkInfo, email string, t int) error {
	if info.Status == StatusDeleted || info.Status == StatusComplete {
		return fmt.Errorf("task alredy complete or delete")
	}
	var minute *time.Ticker
	var start time.Time
	if info.Status == StatusPause {
		minute = time.NewTicker(info.timeLeft)
		start = time.Now().Add(info.timeLeft)
	}
	if info.Status == StatusProgress {
		minute = time.NewTicker(time.Minute * time.Duration(t))
		start = time.Now().Add(time.Minute * time.Duration(t))
	}
	go func() {
		select {
		case <-minute.C:
			info.Status = StatusComplete
			return
		case <-info.ctx.Done():
			info.timeLeft = time.Until(start)
			return
		}
	}()
	go func(start time.Time, info *WorkInfo) {
		for {

			var wg sync.WaitGroup
			file, _ := os.Create("result")

			wg.Add(1)
			go func() {
				select {
				case <-info.ctx.Done():
					return
				default:
					sites.POST0(email)
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST1(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST10(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST100(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST101(email)
				if status >= 400 {
					file.WriteString("POST101 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST102(email)
				if status >= 400 {
					file.WriteString("POST102 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST103(email)
				if status >= 400 {
					file.WriteString("POST103 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST104(email)
				if status >= 400 {
					file.WriteString("POST104 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST105(email)
				if status >= 400 {
					file.WriteString("POST105 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST106(email)
				if status >= 400 {
					file.WriteString("POST106 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST107(email)
				if status >= 400 {
					file.WriteString("POST107 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST108(email)
				if status >= 400 {
					file.WriteString("POST108 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST109(email)
				if status >= 400 {
					file.WriteString("POST109 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST11(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST110(email)
				if status >= 400 {
					file.WriteString("POST110 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST111(email)
				if status >= 400 {
					file.WriteString("POST111 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST112(email)
				if status >= 400 {
					file.WriteString("POST112 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST113(email)
				if status >= 400 {
					file.WriteString("POST113 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST114(email)
				if status >= 400 {
					file.WriteString("POST114 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST115(email)
				if status >= 400 {
					file.WriteString("POST115 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST116(email)
				if status >= 400 {
					file.WriteString("POST116 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST117(email)
				if status >= 400 {
					file.WriteString("POST117 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST118(email)
				if status >= 400 {
					file.WriteString("POST118 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST119(email)
				if status >= 400 {
					file.WriteString("POST119 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST12(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST120(email)
				if status >= 400 {
					file.WriteString("POST120 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST121(email)
				if status >= 400 {
					file.WriteString("POST121 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST122(email)
				if status >= 400 {
					file.WriteString("POST122 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST123(email)
				if status >= 400 {
					file.WriteString("POST123 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST124(email)
				if status >= 400 {
					file.WriteString("POST124 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST125(email)
				if status >= 400 {
					file.WriteString("POST125 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST126(email)
				if status >= 400 {
					file.WriteString("POST126 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST127(email)
				if status >= 400 {
					file.WriteString("POST127 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST128(email)
				if status >= 400 {
					file.WriteString("POST128 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST129(email)
				if status >= 400 {
					file.WriteString("POST129 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST13(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST130(email)
				if status >= 400 {
					file.WriteString("POST130 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST131(email)
				if status >= 400 {
					file.WriteString("POST131 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST132(email)
				if status >= 400 {
					file.WriteString("POST132 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST133(email)
				if status >= 400 {
					file.WriteString("POST133 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST134(email)
				if status >= 400 {
					file.WriteString("POST134 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST135(email)
				if status >= 400 {
					file.WriteString("POST135 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST136(email)
				if status >= 400 {
					file.WriteString("POST136 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST137(email)
				if status >= 400 {
					file.WriteString("POST137 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST138(email)
				if status >= 400 {
					file.WriteString("POST138 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST139(email)
				if status >= 400 {
					file.WriteString("POST139 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST14(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST140(email)
				if status >= 400 {
					file.WriteString("POST140 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST141(email)
				if status >= 400 {
					file.WriteString("POST141 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST142(email)
				if status >= 400 {
					file.WriteString("POST142 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST143(email)
				if status >= 400 {
					file.WriteString("POST143 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST144(email)
				if status >= 400 {
					file.WriteString("POST144 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST145(email)
				if status >= 400 {
					file.WriteString("POST145 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST146(email)
				if status >= 400 {
					file.WriteString("POST146 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST147(email)
				if status >= 400 {
					file.WriteString("POST147 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST148(email)
				if status >= 400 {
					file.WriteString("POST148 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST149(email)
				if status >= 400 {
					file.WriteString("POST149 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST15(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST150(email)
				if status >= 400 {
					file.WriteString("POST150 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST151(email)
				if status >= 400 {
					file.WriteString("POST151 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST152(email)
				if status >= 400 {
					file.WriteString("POST152 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST153(email)
				if status >= 400 {
					file.WriteString("POST153 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST154(email)
				if status >= 400 {
					file.WriteString("POST154 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST155(email)
				if status >= 400 {
					file.WriteString("POST155 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST156(email)
				if status >= 400 {
					file.WriteString("POST156 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST157(email)
				if status >= 400 {
					file.WriteString("POST157 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST158(email)
				if status >= 400 {
					file.WriteString("POST158 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST159(email)
				if status >= 400 {
					file.WriteString("POST159 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST16(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST160(email)
				if status >= 400 {
					file.WriteString("POST160 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST161(email)
				if status >= 400 {
					file.WriteString("POST161 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST162(email)
				if status >= 400 {
					file.WriteString("POST162 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST163(email)
				if status >= 400 {
					file.WriteString("POST163 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST164(email)
				if status >= 400 {
					file.WriteString("POST164 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST165(email)
				if status >= 400 {
					file.WriteString("POST165 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST166(email)
				if status >= 400 {
					file.WriteString("POST166 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST167(email)
				if status >= 400 {
					file.WriteString("POST167 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST168(email)
				if status >= 400 {
					file.WriteString("POST168 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST169(email)
				if status >= 400 {
					file.WriteString("POST169 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST17(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST170(email)
				if status >= 400 {
					file.WriteString("POST170 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST171(email)
				if status >= 400 {
					file.WriteString("POST171 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST172(email)
				if status >= 400 {
					file.WriteString("POST172 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST173(email)
				if status >= 400 {
					file.WriteString("POST173 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST174(email)
				if status >= 400 {
					file.WriteString("POST174 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST175(email)
				if status >= 400 {
					file.WriteString("POST175 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST176(email)
				if status >= 400 {
					file.WriteString("POST176 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST177(email)
				if status >= 400 {
					file.WriteString("POST177 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST178(email)
				if status >= 400 {
					file.WriteString("POST178 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST179(email)
				if status >= 400 {
					file.WriteString("POST179 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST18(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST180(email)
				if status >= 400 {
					file.WriteString("POST180 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST181(email)
				if status >= 400 {
					file.WriteString("POST181 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST182(email)
				if status >= 400 {
					file.WriteString("POST182 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST183(email)
				if status >= 400 {
					file.WriteString("POST183 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST184(email)
				if status >= 400 {
					file.WriteString("POST184 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST185(email)
				if status >= 400 {
					file.WriteString("POST185 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST186(email)
				if status >= 400 {
					file.WriteString("POST186 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST187(email)
				if status >= 400 {
					file.WriteString("POST187 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST188(email)
				if status >= 400 {
					file.WriteString("POST188 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST189(email)
				if status >= 400 {
					file.WriteString("POST189 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST19(email)
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST190(email)
				if status >= 400 {
					file.WriteString("POST190 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST191(email)
				if status >= 400 {
					file.WriteString("POST191 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST192(email)
				if status >= 400 {
					file.WriteString("POST192 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST193(email)
				if status >= 400 {
					file.WriteString("POST193 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST194(email)
				if status >= 400 {
					file.WriteString("POST194 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST195(email)
				if status >= 400 {
					file.WriteString("POST195 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST196(email)
				if status >= 400 {
					file.WriteString("POST196 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST197(email)
				if status >= 400 {
					file.WriteString("POST197 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST2(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST20(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST21(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST22(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST23(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST24(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST25(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST26(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST27(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST28(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST29(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST3(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST30(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST31(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST32(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST33(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST34(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST35(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST36(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST37(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST38(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST39(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST4(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST40(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST41(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST42(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST43(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST44(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST45(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST46(email)
			}()
			wg.Add(1)
			go func() {

				sites.POST47(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST48(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST49(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST5(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST50(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST51(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST52(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST53(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST54(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST55(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST56(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST57(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST58(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST59(email)

			}()
			wg.Add(1)
			go func() {

				sites.POST6(email)

			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST60(email)
				if status >= 400 {
					file.WriteString("POST60 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST63(email)
				if status >= 400 {
					file.WriteString("POST63 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST64(email)
				if status >= 400 {
					file.WriteString("POST64 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST65(email)
				if status >= 400 {
					file.WriteString("POST65 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST67(email)
				if status >= 400 {
					file.WriteString("POST67 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST68(email)
				if status >= 400 {
					file.WriteString("POST68 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST69(email)
				if status >= 400 {
					file.WriteString("POST69 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST7(email)

			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST71(email)
				if status >= 400 {
					file.WriteString("POST71 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST72(email)
				if status >= 400 {
					file.WriteString("POST72 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST73(email)
				if status >= 400 {
					file.WriteString("POST73 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST77(email)
				if status >= 400 {
					file.WriteString("POST77 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST78(email)
				if status >= 400 {
					file.WriteString("POST78 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST793(email)
				if status >= 400 {
					file.WriteString("POST793 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST796(email)
				if status >= 400 {
					file.WriteString("POST796 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST797(email)
				if status >= 400 {
					file.WriteString("POST797 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST798(email)
				if status >= 400 {
					file.WriteString("POST798 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST799(email)
				if status >= 400 {
					file.WriteString("POST799 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST8(email)

			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST80(email)
				if status >= 400 {
					file.WriteString("POST80 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST800(email)
				if status >= 400 {
					file.WriteString("POST800 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST801(email)
				if status >= 400 {
					file.WriteString("POST801 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST802(email)
				if status >= 400 {
					file.WriteString("POST802 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST804(email)
				if status >= 400 {
					file.WriteString("POST804 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST805(email)
				if status >= 400 {
					file.WriteString("POST805 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST807(email)
				if status >= 400 {
					file.WriteString("POST807 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST808(email)
				if status >= 400 {
					file.WriteString("POST808 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST81(email)
				if status >= 400 {
					file.WriteString("POST81 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST811(email)
				if status >= 400 {
					file.WriteString("POST811 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST812(email)
				if status >= 400 {
					file.WriteString("POST812 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST815(email)
				if status >= 400 {
					file.WriteString("POST815 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST818(email)
				if status >= 400 {
					file.WriteString("POST818 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST82(email)
				if status >= 400 {
					file.WriteString("POST82 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST820(email)
				if status >= 400 {
					file.WriteString("POST820 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST821(email)
				if status >= 400 {
					file.WriteString("POST821 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST824(email)
				if status >= 400 {
					file.WriteString("POST824 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST825(email)
				if status >= 400 {
					file.WriteString("POST825 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST826(email)
				if status >= 400 {
					file.WriteString("POST826 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST827(email)
				if status >= 400 {
					file.WriteString("POST827 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST828(email)
				if status >= 400 {
					file.WriteString("POST828 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST83(email)
				if status >= 400 {
					file.WriteString("POST83 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST830(email)
				if status >= 400 {
					file.WriteString("POST830 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST831(email)
				if status >= 400 {
					file.WriteString("POST831 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST832(email)
				if status >= 400 {
					file.WriteString("POST832 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST833(email)
				if status >= 400 {
					file.WriteString("POST833 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST835(email)
				if status >= 400 {
					file.WriteString("POST835 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST836(email)
				if status >= 400 {
					file.WriteString("POST836 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST837(email)
				if status >= 400 {
					file.WriteString("POST837 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST838(email)
				if status >= 400 {
					file.WriteString("POST838 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST839(email)
				if status >= 400 {
					file.WriteString("POST839 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST84(email)
				if status >= 400 {
					file.WriteString("POST84 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST840(email)
				if status >= 400 {
					file.WriteString("POST840 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST842(email)
				if status >= 400 {
					file.WriteString("POST842 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST843(email)
				if status >= 400 {
					file.WriteString("POST843 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST844(email)
				if status >= 400 {
					file.WriteString("POST844 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST845(email)
				if status >= 400 {
					file.WriteString("POST845 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST846(email)
				if status >= 400 {
					file.WriteString("POST846 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST847(email)
				if status >= 400 {
					file.WriteString("POST847 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST848(email)
				if status >= 400 {
					file.WriteString("POST848 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST849(email)
				if status >= 400 {
					file.WriteString("POST849 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST85(email)
				if status >= 400 {
					file.WriteString("POST85 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST850(email)
				if status >= 400 {
					file.WriteString("POST850 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST851(email)
				if status >= 400 {
					file.WriteString("POST851 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST854(email)
				if status >= 400 {
					file.WriteString("POST854 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST855(email)
				if status >= 400 {
					file.WriteString("POST855 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST856(email)
				if status >= 400 {
					file.WriteString("POST856 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST857(email)
				if status >= 400 {
					file.WriteString("POST857 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST86(email)
				if status >= 400 {
					file.WriteString("POST86 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST862(email)
				if status >= 400 {
					file.WriteString("POST862 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST864(email)
				if status >= 400 {
					file.WriteString("POST864 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST867(email)
				if status >= 400 {
					file.WriteString("POST867 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST868(email)
				if status >= 400 {
					file.WriteString("POST868 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST869(email)
				if status >= 400 {
					file.WriteString("POST869 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST87(email)
				if status >= 400 {
					file.WriteString("POST87 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST870(email)
				if status >= 400 {
					file.WriteString("POST870 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST871(email)
				if status >= 400 {
					file.WriteString("POST871 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST872(email)
				if status >= 400 {
					file.WriteString("POST872 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST873(email)
				if status >= 400 {
					file.WriteString("POST873 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST874(email)
				if status >= 400 {
					file.WriteString("POST874 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST875(email)
				if status >= 400 {
					file.WriteString("POST875 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST876(email)
				if status >= 400 {
					file.WriteString("POST876 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST877(email)
				if status >= 400 {
					file.WriteString("POST877 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST878(email)
				if status >= 400 {
					file.WriteString("POST878 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST879(email)
				if status >= 400 {
					file.WriteString("POST879 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST88(email)
				if status >= 400 {
					file.WriteString("POST88 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST882(email)
				if status >= 400 {
					file.WriteString("POST882 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST883(email)
				if status >= 400 {
					file.WriteString("POST883 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST884(email)
				if status >= 400 {
					file.WriteString("POST884 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST885(email)
				if status >= 400 {
					file.WriteString("POST885 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST888(email)
				if status >= 400 {
					file.WriteString("POST888 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST89(email)
				if status >= 400 {
					file.WriteString("POST89 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST890(email)
				if status >= 400 {
					file.WriteString("POST890 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST891(email)
				if status >= 400 {
					file.WriteString("POST891 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST895(email)
				if status >= 400 {
					file.WriteString("POST895 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST897(email)
				if status >= 400 {
					file.WriteString("POST897 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST899(email)
				if status >= 400 {
					file.WriteString("POST899 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				sites.POST9(email)

			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST90(email)
				if status >= 400 {
					file.WriteString("POST90 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST901(email)
				if status >= 400 {
					file.WriteString("POST901 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST902(email)
				if status >= 400 {
					file.WriteString("POST902 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST905(email)
				if status >= 400 {
					file.WriteString("POST905 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST906(email)
				if status >= 400 {
					file.WriteString("POST906 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST907(email)
				if status >= 400 {
					file.WriteString("POST907 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST909(email)
				if status >= 400 {
					file.WriteString("POST909 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST91(email)
				if status >= 400 {
					file.WriteString("POST91 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST910(email)
				if status >= 400 {
					file.WriteString("POST910 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST911(email)
				if status >= 400 {
					file.WriteString("POST911 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST912(email)
				if status >= 400 {
					file.WriteString("POST912 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST913(email)
				if status >= 400 {
					file.WriteString("POST913 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST915(email)
				if status >= 400 {
					file.WriteString("POST915 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST916(email)
				if status >= 400 {
					file.WriteString("POST916 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST917(email)
				if status >= 400 {
					file.WriteString("POST917 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST918(email)
				if status >= 400 {
					file.WriteString("POST918 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST919(email)
				if status >= 400 {
					file.WriteString("POST919 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST92(email)
				if status >= 400 {
					file.WriteString("POST92 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST920(email)
				if status >= 400 {
					file.WriteString("POST920 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST921(email)
				if status >= 400 {
					file.WriteString("POST921 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST922(email)
				if status >= 400 {
					file.WriteString("POST922 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST923(email)
				if status >= 400 {
					file.WriteString("POST923 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST925(email)
				if status >= 400 {
					file.WriteString("POST925 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST926(email)
				if status >= 400 {
					file.WriteString("POST926 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST927(email)
				if status >= 400 {
					file.WriteString("POST927 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST928(email)
				if status >= 400 {
					file.WriteString("POST928 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST929(email)
				if status >= 400 {
					file.WriteString("POST929 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST93(email)
				if status >= 400 {
					file.WriteString("POST93 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST930(email)
				if status >= 400 {
					file.WriteString("POST930 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST931(email)
				if status >= 400 {
					file.WriteString("POST931 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST932(email)
				if status >= 400 {
					file.WriteString("POST932 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST933(email)
				if status >= 400 {
					file.WriteString("POST933 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST935(email)
				if status >= 400 {
					file.WriteString("POST935 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST937(email)
				if status >= 400 {
					file.WriteString("POST937 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST938(email)
				if status >= 400 {
					file.WriteString("POST938 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST939(email)
				if status >= 400 {
					file.WriteString("POST939 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST94(email)
				if status >= 400 {
					file.WriteString("POST94 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST940(email)
				if status >= 400 {
					file.WriteString("POST940 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST941(email)
				if status >= 400 {
					file.WriteString("POST941 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST942(email)
				if status >= 400 {
					file.WriteString("POST942 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST944(email)
				if status >= 400 {
					file.WriteString("POST944 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST945(email)
				if status >= 400 {
					file.WriteString("POST945 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST946(email)
				if status >= 400 {
					file.WriteString("POST946 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST948(email)
				if status >= 400 {
					file.WriteString("POST948 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST95(email)
				if status >= 400 {
					file.WriteString("POST95 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST950(email)
				if status >= 400 {
					file.WriteString("POST950 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST951(email)
				if status >= 400 {
					file.WriteString("POST951 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST952(email)
				if status >= 400 {
					file.WriteString("POST952 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST953(email)
				if status >= 400 {
					file.WriteString("POST953 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST955(email)
				if status >= 400 {
					file.WriteString("POST955 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST956(email)
				if status >= 400 {
					file.WriteString("POST956 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST957(email)
				if status >= 400 {
					file.WriteString("POST957 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST958(email)
				if status >= 400 {
					file.WriteString("POST958 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST959(email)
				if status >= 400 {
					file.WriteString("POST959 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST96(email)
				if status >= 400 {
					file.WriteString("POST96 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST960(email)
				if status >= 400 {
					file.WriteString("POST960 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST962(email)
				if status >= 400 {
					file.WriteString("POST962 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST963(email)
				if status >= 400 {
					file.WriteString("POST963 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST966(email)
				if status >= 400 {
					file.WriteString("POST966 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST967(email)
				if status >= 400 {
					file.WriteString("POST967 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST968(email)
				if status >= 400 {
					file.WriteString("POST968 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST969(email)
				if status >= 400 {
					file.WriteString("POST969 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST97(email)
				if status >= 400 {
					file.WriteString("POST97 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST970(email)
				if status >= 400 {
					file.WriteString("POST970 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST972(email)
				if status >= 400 {
					file.WriteString("POST972 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST973(email)
				if status >= 400 {
					file.WriteString("POST973 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST98(email)
				if status >= 400 {
					file.WriteString("POST98 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST99(email)
				if status >= 400 {
					file.WriteString("POST99 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST427(email)
				if status >= 400 {
					file.WriteString("POST427 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST428(email)
				if status >= 400 {
					file.WriteString("POST428 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST429(email)
				if status >= 400 {
					file.WriteString("POST429 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST430(email)
				if status >= 400 {
					file.WriteString("POST430 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST431(email)
				if status >= 400 {
					file.WriteString("POST431 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST432(email)
				if status >= 400 {
					file.WriteString("POST432 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST433(email)
				if status >= 400 {
					file.WriteString("POST433 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST434(email)
				if status >= 400 {
					file.WriteString("POST434 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST435(email)
				if status >= 400 {
					file.WriteString("POST435 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST436(email)
				if status >= 400 {
					file.WriteString("POST436 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST437(email)
				if status >= 400 {
					file.WriteString("POST437 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST438(email)
				if status >= 400 {
					file.WriteString("POST438 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST439(email)
				if status >= 400 {
					file.WriteString("POST439 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST440(email)
				if status >= 400 {
					file.WriteString("POST440 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST441(email)
				if status >= 400 {
					file.WriteString("POST441 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST442(email)
				if status >= 400 {
					file.WriteString("POST442 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST443(email)
				if status >= 400 {
					file.WriteString("POST443 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST444(email)
				if status >= 400 {
					file.WriteString("POST444 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST445(email)
				if status >= 400 {
					file.WriteString("POST445 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST446(email)
				if status >= 400 {
					file.WriteString("POST446 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST447(email)
				if status >= 400 {
					file.WriteString("POST447 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST448(email)
				if status >= 400 {
					file.WriteString("POST448 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST449(email)
				if status >= 400 {
					file.WriteString("POST449 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST450(email)
				if status >= 400 {
					file.WriteString("POST450 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST451(email)
				if status >= 400 {
					file.WriteString("POST451 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST452(email)
				if status >= 400 {
					file.WriteString("POST452 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST453(email)
				if status >= 400 {
					file.WriteString("POST453 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST454(email)
				if status >= 400 {
					file.WriteString("POST454 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST455(email)
				if status >= 400 {
					file.WriteString("POST455 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST456(email)
				if status >= 400 {
					file.WriteString("POST456 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST457(email)
				if status >= 400 {
					file.WriteString("POST457 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST458(email)
				if status >= 400 {
					file.WriteString("POST458 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST459(email)
				if status >= 400 {
					file.WriteString("POST459 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST460(email)
				if status >= 400 {
					file.WriteString("POST460 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST461(email)
				if status >= 400 {
					file.WriteString("POST461 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST462(email)
				if status >= 400 {
					file.WriteString("POST462 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST463(email)
				if status >= 400 {
					file.WriteString("POST463 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST464(email)
				if status >= 400 {
					file.WriteString("POST464 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST465(email)
				if status >= 400 {
					file.WriteString("POST465 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST466(email)
				if status >= 400 {
					file.WriteString("POST466 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST467(email)
				if status >= 400 {
					file.WriteString("POST467 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST468(email)
				if status >= 400 {
					file.WriteString("POST468 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST469(email)
				if status >= 400 {
					file.WriteString("POST469 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST470(email)
				if status >= 400 {
					file.WriteString("POST470 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST471(email)
				if status >= 400 {
					file.WriteString("POST471 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST472(email)
				if status >= 400 {
					file.WriteString("POST472 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST473(email)
				if status >= 400 {
					file.WriteString("POST473 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST474(email)
				if status >= 400 {
					file.WriteString("POST474 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST475(email)
				if status >= 400 {
					file.WriteString("POST475 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST476(email)
				if status >= 400 {
					file.WriteString("POST476 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST477(email)
				if status >= 400 {
					file.WriteString("POST477 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST478(email)
				if status >= 400 {
					file.WriteString("POST478 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST479(email)
				if status >= 400 {
					file.WriteString("POST479 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST480(email)
				if status >= 400 {
					file.WriteString("POST480 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST481(email)
				if status >= 400 {
					file.WriteString("POST481 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST482(email)
				if status >= 400 {
					file.WriteString("POST482 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST483(email)
				if status >= 400 {
					file.WriteString("POST483 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST484(email)
				if status >= 400 {
					file.WriteString("POST484 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST485(email)
				if status >= 400 {
					file.WriteString("POST485 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST486(email)
				if status >= 400 {
					file.WriteString("POST486 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST487(email)
				if status >= 400 {
					file.WriteString("POST487 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST488(email)
				if status >= 400 {
					file.WriteString("POST488 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST489(email)
				if status >= 400 {
					file.WriteString("POST489 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST490(email)
				if status >= 400 {
					file.WriteString("POST490 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST491(email)
				if status >= 400 {
					file.WriteString("POST491 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST492(email)
				if status >= 400 {
					file.WriteString("POST492 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST493(email)
				if status >= 400 {
					file.WriteString("POST493 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST494(email)
				if status >= 400 {
					file.WriteString("POST494 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST495(email)
				if status >= 400 {
					file.WriteString("POST495 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST496(email)
				if status >= 400 {
					file.WriteString("POST496 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST497(email)
				if status >= 400 {
					file.WriteString("POST497 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST498(email)
				if status >= 400 {
					file.WriteString("POST498 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST499(email)
				if status >= 400 {
					file.WriteString("POST499 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST500(email)
				if status >= 400 {
					file.WriteString("POST500 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST501(email)
				if status >= 400 {
					file.WriteString("POST501 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST502(email)
				if status >= 400 {
					file.WriteString("POST502 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST503(email)
				if status >= 400 {
					file.WriteString("POST503 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST504(email)
				if status >= 400 {
					file.WriteString("POST504 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST505(email)
				if status >= 400 {
					file.WriteString("POST505 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST506(email)
				if status >= 400 {
					file.WriteString("POST506 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST507(email)
				if status >= 400 {
					file.WriteString("POST507 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST508(email)
				if status >= 400 {
					file.WriteString("POST508 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST509(email)
				if status >= 400 {
					file.WriteString("POST509 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST510(email)
				if status >= 400 {
					file.WriteString("POST510 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST511(email)
				if status >= 400 {
					file.WriteString("POST511 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST512(email)
				if status >= 400 {
					file.WriteString("POST512 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST513(email)
				if status >= 400 {
					file.WriteString("POST513 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST514(email)
				if status >= 400 {
					file.WriteString("POST514 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST515(email)
				if status >= 400 {
					file.WriteString("POST515 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST516(email)
				if status >= 400 {
					file.WriteString("POST516 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST517(email)
				if status >= 400 {
					file.WriteString("POST517 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST518(email)
				if status >= 400 {
					file.WriteString("POST518 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST519(email)
				if status >= 400 {
					file.WriteString("POST519 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST520(email)
				if status >= 400 {
					file.WriteString("POST520 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST521(email)
				if status >= 400 {
					file.WriteString("POST521 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST522(email)
				if status >= 400 {
					file.WriteString("POST522 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST523(email)
				if status >= 400 {
					file.WriteString("POST523 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST524(email)
				if status >= 400 {
					file.WriteString("POST524 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST525(email)
				if status >= 400 {
					file.WriteString("POST525 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Add(1)
			go func() {

				status, _ := sites.POST526(email)
				if status >= 400 {
					file.WriteString("POST526 incorrect ")
					file.WriteString(fmt.Sprint(status))
					file.WriteString("\n")
				}
			}()
			wg.Wait()
		}
	}(start, info)

	return nil
}
