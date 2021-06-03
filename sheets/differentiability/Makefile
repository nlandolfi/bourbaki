# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differentiability.pdf

differentiability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > differentiability.tex

differentiability.pdf: ../../*.tex ../../trademark.pdf *.tex differentiability.tex
	# pdflatex --file-line-error -interaction=nonstopmode differentiability.tex
	../../latexrun differentiability.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differentiability.pdf
	open differentiability.pdf

o:
	make open

terms: differentiability.tex
	bbk_sheet -mode ts differentiability.tex

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
	rm -f differentiability.*
	make reset
	bbk_sheet -mode mk > Makefile