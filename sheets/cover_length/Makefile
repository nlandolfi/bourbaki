# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cover_length.pdf

cover_length.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cover_length.tex

cover_length.pdf: ../../*.tex ../../trademark.pdf *.tex cover_length.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cover_length.tex
	../../latexrun cover_length.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cover_length.pdf
	open cover_length.pdf

o:
	make open

terms: cover_length.tex
	bbk_sheet -mode ts cover_length.tex

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
	rm -f cover_length.*
	make reset
	bbk_sheet -mode mk > Makefile