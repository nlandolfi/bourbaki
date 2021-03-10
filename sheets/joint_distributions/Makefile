# generated automatically by sheet.go

make:
	make joint_distributions.pdf

joint_distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > joint_distributions.tex

joint_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex joint_distributions.tex
	pdflatex joint_distributions.tex
	make defs


open: joint_distributions.pdf
	open joint_distributions.pdf

o:
	make open

defs: joint_distributions.tex
	defs joint_distributions.tex

spell:
	aspell -c joint_distributions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile