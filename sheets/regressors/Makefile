# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make regressors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

regressors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > regressors.tex

regressors.pdf: ../../*.tex ../../trademark.pdf *.tex regressors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode regressors.tex
	../../latexrun regressors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: regressors.pdf
	open regressors.pdf

o:
	make open

terms: regressors.tex
	bbk_sheet -mode ts regressors.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f regressors.*
	make reset
	bbk_sheet -mode mk > Makefile