# generated automatically by sheet.go

make:
	make sigma_algebras.pdf

sigma_algebras.tex: sheet.tex
	sheet -mode c -in sheet.tex > sigma_algebras.tex

sigma_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebras.tex
	pdflatex sigma_algebras.tex
	make defs


open: sigma_algebras.pdf
	open sigma_algebras.pdf

o:
	make open

defs: sigma_algebras.tex
	defs sigma_algebras.tex

spell:
	aspell -c sigma_algebras.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile