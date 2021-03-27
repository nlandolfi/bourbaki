# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make standard_deviation.pdf
	make graph.pdf

standard_deviation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > standard_deviation.tex

standard_deviation.pdf: ../../*.tex ../../trademark.pdf *.tex standard_deviation.tex
	pdflatex --file-line-error -interaction=nonstopmode standard_deviation.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: standard_deviation.pdf
	open standard_deviation.pdf

o:
	make open

defs: standard_deviation.tex
	defs standard_deviation.tex

spell:
	aspell -c standard_deviation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f standard_deviation.*
	make reset
	bbk_sheet -mode mk > Makefile