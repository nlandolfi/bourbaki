# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make norm_completeness.pdf
	make graph.pdf

norm_completeness.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > norm_completeness.tex

norm_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex norm_completeness.tex
	pdflatex --file-line-error -interaction=nonstopmode norm_completeness.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norm_completeness.pdf
	open norm_completeness.pdf

o:
	make open

defs: norm_completeness.tex
	defs norm_completeness.tex

spell:
	aspell -c norm_completeness.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm norm_completeness.*
	make reset
	bbk_sheet -mode mk > Makefile