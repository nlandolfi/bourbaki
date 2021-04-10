# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make measure_derivatives.pdf
	make graph.pdf

measure_derivatives.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > measure_derivatives.tex

measure_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex measure_derivatives.tex
	# pdflatex --file-line-error -interaction=nonstopmode measure_derivatives.tex
	../../latexrun measure_derivatives.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measure_derivatives.pdf
	open measure_derivatives.pdf

o:
	make open

terms: measure_derivatives.tex
	bbk_sheet -mode ts measure_derivatives.tex

spell:
	aspell -c measure_derivatives.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f measure_derivatives.*
	make reset
	bbk_sheet -mode mk > Makefile