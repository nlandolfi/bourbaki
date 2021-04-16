# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make optimizers.pdf
	make graph.pdf

optimizers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimizers.tex

optimizers.pdf: ../../*.tex ../../trademark.pdf *.tex optimizers.tex
	# pdflatex --file-line-error -interaction=nonstopmode optimizers.tex
	../../latexrun optimizers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimizers.pdf
	open optimizers.pdf

o:
	make open

terms: optimizers.tex
	bbk_sheet -mode ts optimizers.tex

spell:
	aspell -c optimizers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f optimizers.*
	make reset
	bbk_sheet -mode mk > Makefile