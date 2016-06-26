function! SetDark()
	set notermguicolors
	colorscheme il_tuo_colorscheme_
endfunction

function! SetLight()
	set termguicolors
	colorscheme solarised8_light
endfunction

command! Dark call SetDark()
command! Light call SetLight()
"
"function! Colorschemeselector(scheme)
"	if a:scheme == "solarised8_light"
"		echom "Setting termguicolors"
"		set termguicolors
"		execute 'colorscheme ' a:scheme
"	else
"		echom "Setting notermguicolors"
"		set notermguicolors
"		execute 'colorscheme ' a:scheme
"	endif
"endfunction
"command! -nargs=1 Colorscheme call Colorschemeselector(<f-args>)
