# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make sequences.pdf
	make graph.pdf

sequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sequences.tex

sequences.pdf: ../../*.tex ../../trademark.pdf *.tex sequences.tex
	# pdflatex --file-line-error -interaction=nonstopmode sequences.tex
	../../latexrun sequences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sequences.pdf
	open sequences.pdf

o:
	make open

terms: sequences.tex
	bbk_sheet -mode ts sequences.tex

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
	rm -f sequences.*
	make reset
	bbk_sheet -mode mk > Makefile