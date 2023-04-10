# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make second_derivatives.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

second_derivatives.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > second_derivatives.tex

second_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex second_derivatives.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode second_derivatives.tex
	../../latexrun second_derivatives.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: second_derivatives.pdf
	open second_derivatives.pdf

o:
	make open

terms: second_derivatives.tex
	bbk_sheet -mode ts second_derivatives.tex

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
	rm -f second_derivatives.*
	make reset
	bbk_sheet -mode mk > Makefile