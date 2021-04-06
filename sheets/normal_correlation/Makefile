# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make normal_correlation.pdf
	make graph.pdf

normal_correlation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > normal_correlation.tex

normal_correlation.pdf: ../../*.tex ../../trademark.pdf *.tex normal_correlation.tex
	pdflatex --file-line-error -interaction=nonstopmode normal_correlation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_correlation.pdf
	open normal_correlation.pdf

o:
	make open

terms: normal_correlation.tex
	bbk_sheet -mode ts normal_correlation.tex

spell:
	aspell -c normal_correlation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f normal_correlation.*
	make reset
	bbk_sheet -mode mk > Makefile