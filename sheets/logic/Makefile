# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make logic.pdf

logic.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > logic.tex

logic.pdf: ../../*.tex ../../trademark.pdf *.tex logic.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode logic.tex
	../../latexrun logic.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: logic.pdf
	open logic.pdf

o:
	make open

terms: logic.tex
	bbk_sheet -mode ts logic.tex

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
	rm -f logic.*
	make reset
	bbk_sheet -mode mk > Makefile