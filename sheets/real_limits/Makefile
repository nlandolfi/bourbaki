# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_limits.pdf
	make graph.pdf

real_limits.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_limits.tex

real_limits.pdf: ../../*.tex ../../trademark.pdf *.tex real_limits.tex
	# pdflatex --file-line-error -interaction=nonstopmode real_limits.tex
	../../latexrun real_limits.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_limits.pdf
	open real_limits.pdf

o:
	make open

terms: real_limits.tex
	bbk_sheet -mode ts real_limits.tex

spell:
	aspell -c real_limits.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_limits.*
	make reset
	bbk_sheet -mode mk > Makefile