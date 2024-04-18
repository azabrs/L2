package main

type Command{
    Execute() string
}

type ToggleOnCommand struct{
    receiver *Receiver
}

func (c ToggleOnCommand) Execute() string{
    return c.receiver.TogleOn()
}

type ToggleOffCommand struct{
    receiver *Receiver
}

func (c ToggleOffCommand) Execute() string{
    return c.receiver.TogleOff()
}

type Receiver struct{

}

func (r *Receiver)TogleOn() string{
    return "TogleOn"
}

func (r *Receiver)TogleOff() string{
    return "TogleOff"
}

type Invocker struct{
    commands []Command
}

func (i *Invocker)StoreCommand(command Command){
   i.commands = append(i.commands, command) 
}


func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}