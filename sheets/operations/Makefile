# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make operations.pdf
	make graph.pdf

operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > operations.tex

operations.pdf: ../../*.tex ../../trademark.pdf *.tex operations.tex
	# pdflatex --file-line-error -interaction=nonstopmode operations.tex
	../../latexrun operations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: operations.pdf
	open operations.pdf

o:
	make open

terms: operations.tex
	bbk_sheet -mode ts operations.tex

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
	rm -f operations.*
	make reset
	bbk_sheet -mode mk > Makefile