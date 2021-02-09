# generated automatically by sheet.go

make:
	make marginal_distributions.pdf

marginal_distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > marginal_distributions.tex

marginal_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex marginal_distributions.tex
	pdflatex marginal_distributions.tex
	make defs


open: marginal_distributions.pdf
	open marginal_distributions.pdf

o:
	make open

defs: marginal_distributions.tex
	defs marginal_distributions.tex

spell:
	aspell -c marginal_distributions.tex

s:
	make spell