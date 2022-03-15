package vterm

/*
#include <vterm.h>

inline static int _attr_bold(VTermScreenCell *cell) { return cell->attrs.bold; }
inline static int _attr_underline(VTermScreenCell *cell) { return cell->attrs.underline; }
inline static int _attr_italic(VTermScreenCell *cell) { return cell->attrs.italic; }
inline static int _attr_blink(VTermScreenCell *cell) { return cell->attrs.blink; }
inline static int _attr_reverse(VTermScreenCell *cell) { return cell->attrs.reverse; }
inline static int _attr_strike(VTermScreenCell *cell) { return cell->attrs.strike; }
inline static int _attr_font(VTermScreenCell *cell) { return cell->attrs.font; }
inline static int _attr_dwl(VTermScreenCell *cell) { return cell->attrs.dwl; }
inline static int _attr_dhl(VTermScreenCell *cell) { return cell->attrs.dhl; }

int _go_handle_damage(VTermRect, void*);
int _go_handle_bell(void*);
int _go_handle_set_term_prop(VTermProp, VTermValue*, void*);
int _go_handle_resize(int, int, void*);
int _go_handle_moverect(VTermRect, VTermRect, void*);
int _go_handle_movecursor(VTermPos, VTermPos, int, void*);

static VTermScreenCallbacks _screen_callbacks = {
	_go_handle_damage,
	_go_handle_moverect,
	_go_handle_movecursor,
	_go_handle_set_term_prop,
	_go_handle_bell,
	_go_handle_resize,
	NULL,
	NULL
};

static void
_vterm_screen_set_callbacks(VTermScreen *screen, uintptr_t user) {
	vterm_screen_set_callbacks(screen, &_screen_callbacks, (void *)user);
}

static bool _vterm_value_get_boolean(VTermValue *val) {
	return val->boolean;
}

static int _vterm_value_get_number(VTermValue *val) {
	return val->number;
}

static const char *_vterm_value_get_string(VTermValue *val) {
	return val->string.str;
}

typedef struct {
	uint8_t color_type;
	uint8_t red, green, blue;
} rgb_color;

*/
import "C"

import (
	"errors"
	"runtime/cgo"
	"unsafe"
)

type Attr int

const (
	AttrNone       Attr = 0
	AttrBold            = Attr(C.VTERM_ATTR_BOLD)
	AttrUnderline       = Attr(C.VTERM_ATTR_UNDERLINE)
	AttrItalic          = Attr(C.VTERM_ATTR_ITALIC)
	AttrBlink           = Attr(C.VTERM_ATTR_BLINK)
	AttrReverse         = Attr(C.VTERM_ATTR_REVERSE)
	AttrStrike          = Attr(C.VTERM_ATTR_STRIKE)
	AttrFont            = Attr(C.VTERM_ATTR_FONT)
	AttrForeground      = Attr(C.VTERM_ATTR_FOREGROUND)
	AttrBackground      = Attr(C.VTERM_ATTR_BACKGROUND)
	AttrNAttrrs
)

type Modifier uint

const (
	ModNone  Modifier = 0
	ModShift          = Modifier(C.VTERM_MOD_SHIFT)
	ModAlt            = Modifier(C.VTERM_MOD_ALT)
	ModCtrl           = Modifier(C.VTERM_MOD_CTRL)
)

type Key uint

const (
	KeyNone      = Key(0)
	KeyEnter     = Key(C.VTERM_KEY_ENTER)
	KeyTab       = Key(C.VTERM_KEY_TAB)
	KeyBackspace = Key(C.VTERM_KEY_BACKSPACE)
	KeyEscape    = Key(C.VTERM_KEY_ESCAPE)
	KeyUp        = Key(C.VTERM_KEY_UP)
	KeyDown      = Key(C.VTERM_KEY_DOWN)
	KeyLeft      = Key(C.VTERM_KEY_LEFT)
	KeyRight     = Key(C.VTERM_KEY_RIGHT)
	KeyIns       = Key(C.VTERM_KEY_INS)
	KeyDel       = Key(C.VTERM_KEY_DEL)
	KeyHome      = Key(C.VTERM_KEY_HOME)
	KeyEnd       = Key(C.VTERM_KEY_END)
	KeyPageUp    = Key(C.VTERM_KEY_PAGEUP)
	KeyPageDown  = Key(C.VTERM_KEY_PAGEDOWN)
	KeyFunction0 = Key(C.VTERM_KEY_FUNCTION_0)
	KeyKp0       = Key(C.VTERM_KEY_KP_0)
	KeyKp1       = Key(C.VTERM_KEY_KP_1)
	KeyKp2       = Key(C.VTERM_KEY_KP_2)
	KeyKp3       = Key(C.VTERM_KEY_KP_3)
	KeyKp4       = Key(C.VTERM_KEY_KP_4)
	KeyKp5       = Key(C.VTERM_KEY_KP_5)
	KeyKp6       = Key(C.VTERM_KEY_KP_6)
	KeyKp7       = Key(C.VTERM_KEY_KP_7)
	KeyKp8       = Key(C.VTERM_KEY_KP_8)
	KeyKp9       = Key(C.VTERM_KEY_KP_9)
	KeyKpMult    = Key(C.VTERM_KEY_KP_MULT)
	KeyKpPlus    = Key(C.VTERM_KEY_KP_PLUS)
	KeyKpComma   = Key(C.VTERM_KEY_KP_COMMA)
	KeyKpMinus   = Key(C.VTERM_KEY_KP_MINUS)
	KeyKpPeriod  = Key(C.VTERM_KEY_KP_PERIOD)
	KeyKpDivide  = Key(C.VTERM_KEY_KP_DIVIDE)
	KeyKpEnter   = Key(C.VTERM_KEY_KP_ENTER)
	KeyKpEqual   = Key(C.VTERM_KEY_KP_EQUAL)
)

type VTerm struct {
	term   *C.VTerm
	screen *Screen
}

type Pos struct {
	pos C.VTermPos
}

func NewPos(row, col int) *Pos {
	var pos Pos
	pos.pos.col = C.int(col)
	pos.pos.row = C.int(row)
	return &pos
}

func (pos *Pos) Col() int {
	return int(pos.pos.col)
}

func (pos *Pos) Row() int {
	return int(pos.pos.row)
}

type Rect struct {
	rect C.VTermRect
}

func (rect *Rect) StartRow() int {
	return int(rect.rect.start_row)
}

func (rect *Rect) EndRow() int {
	return int(rect.rect.end_row)
}

func (rect *Rect) StartCol() int {
	return int(rect.rect.start_col)
}

func (rect *Rect) EndCol() int {
	return int(rect.rect.end_col)
}

func NewRect(start_row, end_row, start_col, end_col int) *Rect {
	var rect Rect
	rect.rect.start_row = C.int(start_row)
	rect.rect.end_row = C.int(end_row)
	rect.rect.start_col = C.int(start_col)
	rect.rect.end_col = C.int(end_col)
	return &rect
}

type ScreenCell struct {
	cell C.VTermScreenCell
}

type ParserCallbacks struct {
	Text func([]byte, interface{}) int
	/*
		int (*control)(unsigned char control, void *user);
		int (*control)(unsigned char control, void *user);
		int (*escape)(const char *bytes, size_t len, void *user);
		int (*csi)(const char *leader, const long args[], int argcount, const char *intermed, char command, void *user);
		int (*osc)(const char *command, size_t cmdlen, void *user);
		int (*dcs)(const char *command, size_t cmdlen, void *user);
		int (*resize)(int rows, int cols, void *user);
	*/
}

type VTermColor struct {
	color C.VTermColor
}

func _VTermColorToRGB(col *C.VTermColor) *C.rgb_color {
	return *(**C.rgb_color)(unsafe.Pointer(&col))
}

func (c *VTermColor) _to_rgb_color() *C.rgb_color {
	return _VTermColorToRGB(&c.color)
}

const (
	VTERM_COLOR_RGB        uint8 = 0x00
	VTERM_COLOR_INDEXED    uint8 = 0x01
	VTERM_COLOR_TYPE_MASK  uint8 = 0x01
	VTERM_COLOR_DEFAULT_FG uint8 = 0x02
	VTERM_COLOR_DEFAULT_BG uint8 = 0x04
)

func (c *VTermColor) IsIndexed() bool {
	return (c.GetType() & VTERM_COLOR_TYPE_MASK) == VTERM_COLOR_INDEXED
}

func (c *VTermColor) IsRgb() bool {
	return (c.GetType() & VTERM_COLOR_TYPE_MASK) == VTERM_COLOR_RGB
}

func (c *VTermColor) IsDefaultFg() bool {
	return c.GetType()&VTERM_COLOR_DEFAULT_FG > 0
}

func (c *VTermColor) IsDefaultBg() bool {
	return c.GetType()&VTERM_COLOR_DEFAULT_BG > 0
}

func (c *VTermColor) GetType() uint8 {
	colors := c._to_rgb_color()
	return uint8(colors.color_type)
}

func (c *VTermColor) GetIndex() uint8 {
	colors := c._to_rgb_color()
	return uint8(colors.red)
}

func (c *VTermColor) GetRGB() (r, g, b uint8) {
	colors := c._to_rgb_color()
	return uint8(colors.red), uint8(colors.green), uint8(colors.blue)
}

func (sc *ScreenCell) Chars() []rune {
	chars := make([]rune, int(sc.cell.width))
	for i := 0; i < len(chars); i++ {
		chars[i] = rune(sc.cell.chars[i])
	}
	return chars
}

func (sc *ScreenCell) Width() int {
	return int(sc.cell.width)
}

func (sc *ScreenCell) Fg() VTermColor {
	return VTermColor{sc.cell.fg}
}

func (sc *ScreenCell) Bg() VTermColor {
	return VTermColor{sc.cell.bg}
}

type Attrs struct {
	Bold      int
	Underline int
	Italic    int
	Blink     int
	Reverse   int
	Strike    int
	Font      int
	Dwl       int
	Dhl       int
}

func (sc *ScreenCell) Attrs() *Attrs {
	return &Attrs{
		Bold:      int(C._attr_bold(&sc.cell)),
		Underline: int(C._attr_underline(&sc.cell)),
		Italic:    int(C._attr_italic(&sc.cell)),
		Blink:     int(C._attr_blink(&sc.cell)),
		Reverse:   int(C._attr_reverse(&sc.cell)),
		Strike:    int(C._attr_strike(&sc.cell)),
		Font:      int(C._attr_font(&sc.cell)),
		Dwl:       int(C._attr_dwl(&sc.cell)),
		Dhl:       int(C._attr_dhl(&sc.cell)),
	}
}

func New(rows, cols int) *VTerm {
	term := C.vterm_new(C.int(rows), C.int(cols))
	vt := &VTerm{
		term: term,
		screen: &Screen{
			screen: C.vterm_obtain_screen(term),
		},
	}
	C._vterm_screen_set_callbacks(C.vterm_obtain_screen(term), C.uintptr_t(cgo.NewHandle(vt)))
	return vt
}

func (vt *VTerm) Close() error {
	C.vterm_free(vt.term)
	return nil
}

func (vt *VTerm) Size() (int, int) {
	var rows, cols C.int
	C.vterm_get_size(vt.term, &rows, &cols)
	return int(rows), int(cols)
}

func (vt *VTerm) SetSize(rows, cols int) {
	C.vterm_set_size(vt.term, C.int(rows), C.int(cols))
}

func (vt *VTerm) KeyboardStartPaste() {
	C.vterm_keyboard_start_paste(vt.term)
}

func (vt *VTerm) KeyboardStopPaste() {
	C.vterm_keyboard_end_paste(vt.term)
}

func (vt *VTerm) KeyboardUnichar(c rune, mods Modifier) {
	C.vterm_keyboard_unichar(vt.term, C.uint32_t(c), C.VTermModifier(mods))
}

func (vt *VTerm) KeyboardKey(key Key, mods Modifier) {
	C.vterm_keyboard_key(vt.term, C.VTermKey(key), C.VTermModifier(mods))
}

func (vt *VTerm) ObtainState() *State {
	return &State{
		state: C.vterm_obtain_state(vt.term),
	}
}

func (vt *VTerm) Read(b []byte) (int, error) {
	curlen := C.vterm_output_read(vt.term, (*C.char)(unsafe.Pointer(&b[0])), C.size_t(len(b)))
	return int(curlen), nil
}

func (vt *VTerm) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	return int(C.vterm_input_write(vt.term, (*C.char)(unsafe.Pointer(&b[0])), C.size_t(len(b)))), nil
}

func (vt *VTerm) ObtainScreen() *Screen {
	return vt.screen
}

func (vt *VTerm) UTF8() bool {
	return C.vterm_get_utf8(vt.term) != C.int(0)
}

func (vt *VTerm) SetUTF8(b bool) {
	var v C.int
	if b {
		v = 1
	}
	C.vterm_set_utf8(vt.term, v)
}

type VTermValue struct {
	Boolean bool
	Number  int
	String  string
	Color   VTermColor
}

const (
	_ = iota
	VTERM_PROP_CURSORVISIBLE
	VTERM_PROP_CURSORBLINK
	VTERM_PROP_ALTSCREEN
	VTERM_PROP_TITLE
	VTERM_PROP_ICONNAME
	VTERM_PROP_REVERSE
	VTERM_PROP_CURSORSHAPE
	VTERM_PROP_MOUSE
)

type Screen struct {
	screen *C.VTermScreen

	UserData      interface{}
	OnDamage      func(*Rect) int
	OnResize      func(int, int) int
	OnMoveRect    func(*Rect, *Rect) int
	OnMoveCursor  func(*Pos, *Pos, bool) int
	OnBell        func() int
	OnSetTermProp func(int, *VTermValue) int
	/*
		int (*sb_pushline)(int cols, const VTermScreenCell *cells, void *user);
		int (*sb_popline)(int cols, VTermScreenCell *cells, void *user);
	*/
}

func (scr *Screen) Flush() error {
	C.vterm_screen_flush_damage(scr.screen)
	return nil // TODO
}

func (sc *Screen) GetCellAt(row, col int) (*ScreenCell, error) {
	return sc.GetCell(NewPos(row, col))
}

func (sc *Screen) GetCell(pos *Pos) (*ScreenCell, error) {
	var cell ScreenCell
	if C.vterm_screen_get_cell(sc.screen, pos.pos, &cell.cell) == 0 {
		return nil, errors.New("GetCell")
	}
	return &cell, nil
}

func (scr *Screen) GetChars(r *[]rune, rect *Rect) int {
	l := len(*r)
	buf := make([]C.uint32_t, l)
	ret := int(C.vterm_screen_get_chars(scr.screen, &buf[0], C.size_t(l), rect.rect))
	*r = make([]rune, ret)
	for i := 0; i < ret; i++ {
		(*r)[i] = rune(buf[i])
	}
	return ret
}

func (scr *Screen) Reset(hard bool) {
	var v C.int
	if hard {
		v = 1
	}
	C.vterm_screen_reset(scr.screen, v)
}

func (scr *Screen) EnableAltScreen(e bool) {
	var v C.int
	if e {
		v = 1
	}
	C.vterm_screen_enable_altscreen(scr.screen, v)
}

func (scr *Screen) IsEOL(pos *Pos) bool {
	return C.vterm_screen_is_eol(scr.screen, pos.pos) != C.int(0)
}

type State struct {
	state *C.VTermState
}

func (s *State) SetDefaultColors(fg, bg VTermColor) {
	C.vterm_state_set_default_colors(s.state, &fg.color, &bg.color)
}

// index between 0 and 15, 0-7 are normal colors and 8-15 are bright colors.
func (s *State) SetPaletteColor(index int, col VTermColor) {
	if index < 0 || index >= 256 {
		panic("Index out of range")
	}
	C.vterm_state_set_palette_color(s.state, C.int(index), &col.color)
}

func (s *State) GetDefaultColors() (fg, bg VTermColor) {
	c_fg := C.VTermColor{}
	c_bg := C.VTermColor{}
	C.vterm_state_get_default_colors(s.state, &c_fg, &c_bg)
	fg = VTermColor{c_fg}
	bg = VTermColor{c_bg}
	return
}

func (s *State) GetCursorPos() (int, int) {
	var pos C.VTermPos
	C.vterm_state_get_cursorpos(s.state, &pos)
	return int(pos.row), int(pos.col)
}

// index between 0 and 15, 0-7 are normal colors and 8-15 are bright colors.
func (s *State) GetPaletteColor(index int) VTermColor {
	if index < 0 || index >= 256 {
		panic("Index out of range")
	}
	c_color := C.VTermColor{}
	C.vterm_state_get_palette_color(s.state, C.int(index), &c_color)
	return VTermColor{c_color}
}

//export _go_handle_damage
func _go_handle_damage(rect C.VTermRect, user unsafe.Pointer) C.int {
	onDamage := cgo.Handle(user).Value().(*VTerm).ObtainScreen().OnDamage
	if onDamage != nil {
		return C.int(onDamage(&Rect{rect: rect}))
	}
	return 0
}

//export _go_handle_bell
func _go_handle_bell(user unsafe.Pointer) C.int {
	onBell := cgo.Handle(user).Value().(*VTerm).ObtainScreen().OnBell
	if onBell != nil {
		return C.int(onBell())
	}
	return 0
}

const (
	vterm_valuetype_bool   = 1
	vterm_valuetype_int    = 2
	vterm_valuetype_string = 3
	vterm_valuetype_color  = 4
)

//export _go_handle_set_term_prop
func _go_handle_set_term_prop(prop C.VTermProp, val *C.VTermValue, user unsafe.Pointer) C.int {
	onSetTermProp := cgo.Handle(user).Value().(*VTerm).ObtainScreen().OnSetTermProp

	if onSetTermProp != nil {
		value := VTermValue{}

		switch int(C.vterm_get_prop_type(prop)) {
		case vterm_valuetype_bool:
			value.Boolean = bool(C._vterm_value_get_boolean(val))
		case vterm_valuetype_int:
			value.Number = int(C._vterm_value_get_number(val))
		case vterm_valuetype_string:
			value.String = C.GoString(C._vterm_value_get_string(val))
		case vterm_valuetype_color:
			return 0 // TODO
		default:
			return 0
		}

		return C.int(onSetTermProp(int(prop), &value))
	}
	return 0
}

//export _go_handle_resize
func _go_handle_resize(row, col C.int, user unsafe.Pointer) C.int {
	onResize := cgo.Handle(user).Value().(*VTerm).ObtainScreen().OnResize
	if onResize != nil {
		return C.int(onResize(int(row), int(col)))
	}
	return 0
}

//export _go_handle_moverect
func _go_handle_moverect(dest, src C.VTermRect, user unsafe.Pointer) C.int {
	onMoveRect := cgo.Handle(user).Value().(*VTerm).ObtainScreen().OnMoveRect
	if onMoveRect != nil {
		return C.int(onMoveRect(&Rect{rect: dest}, &Rect{rect: src}))
	}
	return 0
}

//export _go_handle_movecursor
func _go_handle_movecursor(pos, oldpos C.VTermPos, visible C.int, user unsafe.Pointer) C.int {
	onMoveCursor := cgo.Handle(user).Value().(*VTerm).ObtainScreen().OnMoveCursor
	if onMoveCursor != nil {
		var b bool
		if visible != C.int(0) {
			b = true
		}
		return C.int(onMoveCursor(&Pos{pos: pos}, &Pos{pos: oldpos}, b))
	}
	return 0
}
