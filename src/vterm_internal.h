#ifndef __VTERM_INTERNAL_H__
#define __VTERM_INTERNAL_H__

#include "vterm.h"

#include <stdarg.h>

typedef struct _VTermEncoding VTermEncoding;

struct _VTermState
{
  VTerm *vt;

  const VTermStateCallbacks *callbacks;
  void *cbdata;

  int rows;
  int cols;

  /* Current cursor position */
  VTermPos pos;

  int at_phantom; /* True if we're on the "81st" phantom column to defer a wraparound */

  int scrollregion_start;
  int scrollregion_end; /* -1 means unbounded */
#define SCROLLREGION_END(state) ((state)->scrollregion_end > -1 ? (state)->scrollregion_end : (state)->rows)

  /* Bitvector of tab stops */
  unsigned char *tabstops;

  /* Mouse state */
  int mouse_col, mouse_row;
  int mouse_buttons;
  int mouse_flags;

  /* Last glyph output, for Unicode recombining purposes */
  uint32_t *combine_chars;
  size_t combine_chars_size;
  int combine_width; // The width of the glyph above
  VTermPos combine_pos;   // Position before movement

  struct {
    int keypad:1;
    int cursor:1;
    int autowrap:1;
    int insert:1;
    int cursor_visible:1;
    int cursor_blink:1;
    unsigned int cursor_shape:2;
    int alt_screen:1;
    int origin:1;
  } mode;

  VTermEncoding *encoding[4];
  int gl_set, gr_set;

  struct {
    unsigned int bold:1;
    unsigned int underline:2;
    unsigned int italic:1;
    unsigned int blink:1;
    unsigned int reverse:1;
    unsigned int strike:1;
    unsigned int font:4; /* To store 0-9 */
  } pen;

  VTermColor default_fg;
  VTermColor default_bg;

  /* Saved state under DEC mode 1048/1049 */
  struct {
    VTermPos pos;

    struct {
      int cursor_visible:1;
      int cursor_blink:1;
      unsigned int cursor_shape:2;
    } mode;
  } saved;
};

struct _VTerm
{
  VTermAllocatorFunctions *allocator;
  void *allocdata;

  int rows;
  int cols;

  int is_utf8;

  enum VTermParserState {
    NORMAL,
    ESC,
    CSI,
    OSC,
    DCS,
  } parser_state;
  const VTermParserCallbacks *parser_callbacks;
  void *cbdata;

  /* len == malloc()ed size; cur == number of valid bytes */
  char  *strbuffer;
  size_t strbuffer_len;
  size_t strbuffer_cur;

  char  *outbuffer;
  size_t outbuffer_len;
  size_t outbuffer_cur;

  VTermState *state;
  VTermScreen *screen;
};

struct _VTermEncoding {
  int (*decode)(VTermEncoding *enc, uint32_t cp[], int *cpi, int cplen,
                  const char bytes[], size_t *pos, size_t len);
};

typedef enum {
  ENC_UTF8,
  ENC_SINGLE_94
} VTermEncodingType;

void *vterm_allocator_malloc(VTerm *vt, size_t size);
void  vterm_allocator_free(VTerm *vt, void *ptr);

void vterm_push_output_bytes(VTerm *vt, const char *bytes, size_t len);
void vterm_push_output_vsprintf(VTerm *vt, const char *format, va_list args);
void vterm_push_output_sprintf(VTerm *vt, const char *format, ...);

void vterm_state_free(VTermState *state);

void vterm_state_resetpen(VTermState *state);
void vterm_state_setpen(VTermState *state, const long args[], int argcount);

void vterm_screen_free(VTermScreen *screen);

VTermEncoding *vterm_lookup_encoding(VTermEncodingType type, char designation);

int vterm_unicode_width(int codepoint);
int vterm_unicode_is_combining(int codepoint);

#endif
