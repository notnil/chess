package uci

import (
	"errors"
	"strings"
)

// Option corresponds to the "option" engine output:
// This command tells the GUI which parameters can be changed in the engine.
// 	This should be sent once at engine startup after the "uci" and the "id" commands
// 	if any parameter can be changed in the engine.
// 	The GUI should parse this and build a dialog for the user to change the settings.
// 	Note that not every option needs to appear in this dialog as some options like
// 	"Ponder", "UCI_AnalyseMode", etc. are better handled elsewhere or are set automatically.
// 	If the user wants to change some settings, the GUI will send a "setoption" command to the engine.
// 	Note that the GUI need not send the setoption command when starting the engine for every option if
// 	it doesn't want to change the default value.
// 	For all allowed combinations see the example below,
// 	as some combinations of this tokens don't make sense.
// 	One string will be sent for each parameter.
// 	* name
// 		The option has the name id.
// 		Certain options have a fixed value for , which means that the semantics of this option is fixed.
// 		Usually those options should not be displayed in the normal engine options window of the GUI but
// 		get a special treatment. "Pondering" for example should be set automatically when pondering is
// 		enabled or disabled in the GUI options. The same for "UCI_AnalyseMode" which should also be set
// 		automatically by the GUI. All those certain options have the prefix "UCI_" except for the
// 		first 6 options below. If the GUI get an unknown Option with the prefix "UCI_", it should just
// 		ignore it and not display it in the engine's options dialog.
// 		*  = Hash, type is spin
// 			the value in MB for memory for hash tables can be changed,
// 			this should be answered with the first "setoptions" command at program boot
// 			if the engine has sent the appropriate "option name Hash" command,
// 			which should be supported by all engines!
// 			So the engine should use a very small hash first as default.
// 		*  = NalimovPath, type string
// 			this is the path on the hard disk to the Nalimov compressed format.
// 			Multiple directories can be concatenated with ";"
// 		*  = NalimovCache, type spin
// 			this is the size in MB for the cache for the nalimov table bases
// 			These last two options should also be present in the initial options exchange dialog
// 			when the engine is booted if the engine supports it
// 		*  = Ponder, type check
// 			this means that the engine is able to ponder.
// 			The GUI will send this whenever pondering is possible or not.
// 			Note: The engine should not start pondering on its own if this is enabled, this option is only
// 			needed because the engine might change its time management algorithm when pondering is allowed.
// 		*  = OwnBook, type check
// 			this means that the engine has its own book which is accessed by the engine itself.
// 			if this is set, the engine takes care of the opening book and the GUI will never
// 			execute a move out of its book for the engine. If this is set to false by the GUI,
// 			the engine should not access its own book.
// 		*  = MultiPV, type spin
// 			the engine supports multi best line or k-best mode. the default value is 1
// 		*  = UCI_ShowCurrLine, type check, should be false by default,
// 			the engine can show the current line it is calculating. see "info currline" above.
// 		*  = UCI_ShowRefutations, type check, should be false by default,
// 			the engine can show a move and its refutation in a line. see "info refutations" above.
// 		*  = UCI_LimitStrength, type check, should be false by default,
// 			The engine is able to limit its strength to a specific Elo number,
// 		   This should always be implemented together with "UCI_Elo".
// 		*  = UCI_Elo, type spin
// 			The engine can limit its strength in Elo within this interval.
// 			If UCI_LimitStrength is set to false, this value should be ignored.
// 			If UCI_LimitStrength is set to true, the engine should play with this specific strength.
// 		   This should always be implemented together with "UCI_LimitStrength".
// 		*  = UCI_AnalyseMode, type check
// 		   The engine wants to behave differently when analysing or playing a game.
// 		   For example when playing it can use some kind of learning.
// 		   This is set to false if the engine is playing a game, otherwise it is true.
// 		 *  = UCI_Opponent, type string
// 		   With this command the GUI can send the name, title, elo and if the engine is playing a human
// 		   or computer to the engine.
// 		   The format of the string has to be [GM|IM|FM|WGM|WIM|none] [|none] [computer|human]
// 		   Example:
// 		   "setoption name UCI_Opponent value GM 2800 human Gary Kasparow"
// 		   "setoption name UCI_Opponent value none none computer Shredder"
// 	* type
// 		The option has type t.
// 		There are 5 different types of options the engine can send
// 		* check
// 			a checkbox that can either be true or false
// 		* spin
// 			a spin wheel that can be an integer in a certain range
// 		* combo
// 			a combo box that can have different predefined strings as a value
// 		* button
// 			a button that can be pressed to send a command to the engine
// 		* string
// 			a text field that has a string as a value,
// 			an empty string has the value ""
// 	* default
// 		the default value of this parameter is x
// 	* min
// 		the minimum value of this parameter is x
// 	* max
// 		the maximum value of this parameter is x
// 	* var
// 		a predefined value of this parameter is x
// 	Example:
//     Here are 5 strings for each of the 5 possible types of options
// 	   "option name Nullmove type check default true\n"
//       "option name Selectivity type spin default 2 min 0 max 4\n"
// 	   "option name Style type combo default Normal var Solid var Normal var Risky\n"
// 	   "option name NalimovPath type string default c:\\n"
// 	   "option name Clear Hash type button\n"
type Option struct {
	Name    string
	Type    OptionType
	Default string
	Min     string
	Max     string
	Vars    []string
}

// UnmarshalText implements the encoding.TextUnmarshaler interface and parses
// data like the following:
// option name EvalFile type string default nn-82215d0fd0df.nnue
func (o *Option) UnmarshalText(text []byte) error {
	o.Type = OptionNoType
	parts := strings.Split(string(text), " ")
	if len(parts) == 0 {
		return errors.New("uci: invalid option line")
	}
	if parts[0] != "option" {
		return errors.New("uci: invalid option line")
	}
	ref := ""
	for i := 1; i < len(parts); i++ {
		s := parts[i]
		if ref == "" {
			ref = s
			continue
		}
		switch ref {
		case "name":
			o.Name = s
		case "type":
			ot, err := optionTypeFromString(s)
			if err != nil {
				return err
			}
			o.Type = ot
		case "default":
			o.Default = s
		case "min":
			o.Min = s
		case "max":
			o.Max = s
		case "var":
			o.Vars = append(o.Vars, s)
		}
		ref = ""
	}
	if o.Name == "" || o.Type == OptionNoType {
		return errors.New("uci: invalid option line")
	}
	return nil
}

// OptionType corresponds to the "option"'s type engine output:
// * type
// The option has type t.
// There are 5 different types of options the engine can send
// * check
// 	a checkbox that can either be true or false
// * spin
// 	a spin wheel that can be an integer in a certain range
// * combo
// 	a combo box that can have different predefined strings as a value
// * button
// 	a button that can be pressed to send a command to the engine
// * string
// 	a text field that has a string as a value,
// 	an empty string has the value ""
type OptionType string

const (
	// OptionNoType indicates no option type
	OptionNoType OptionType = "notype"
	// OptionCheck indicates check option type
	OptionCheck OptionType = "check"
	// OptionSpin indicates spin option type
	OptionSpin OptionType = "spin"
	// OptionCombo indicates combo option type
	OptionCombo OptionType = "combo"
	// OptionButton indicates button option type
	OptionButton OptionType = "button"
	// OptionString indicates string option type
	OptionString OptionType = "string"
)

func optionTypeFromString(s string) (OptionType, error) {
	for _, ot := range []OptionType{OptionCheck, OptionSpin, OptionCombo, OptionButton, OptionString} {
		if string(ot) == s {
			return ot, nil
		}
	}
	return OptionNoType, errors.New("uci: invalid option type " + s)
}
