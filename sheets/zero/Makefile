# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make zero.pdf

zero.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > zero.tex

zero.pdf: ../../*.tex ../../trademark.pdf *.tex zero.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode zero.tex
	../../latexrun zero.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: zero.pdf
	open zero.pdf

o:
	make open

terms: zero.tex
	bbk_sheet -mode ts zero.tex

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
	rm -f zero.*
	make reset
	bbk_sheet -mode mk > Makefile