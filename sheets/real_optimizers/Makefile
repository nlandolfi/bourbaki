# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_optimizers.pdf
	make graph.pdf

real_optimizers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_optimizers.tex

real_optimizers.pdf: ../../*.tex ../../trademark.pdf *.tex real_optimizers.tex
	# pdflatex --file-line-error -interaction=nonstopmode real_optimizers.tex
	../../latexrun real_optimizers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_optimizers.pdf
	open real_optimizers.pdf

o:
	make open

terms: real_optimizers.tex
	bbk_sheet -mode ts real_optimizers.tex

spell:
	aspell -c real_optimizers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_optimizers.*
	make reset
	bbk_sheet -mode mk > Makefile