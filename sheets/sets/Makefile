# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make sets.pdf
	make graph.pdf

sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sets.tex

sets.pdf: ../../*.tex ../../trademark.pdf *.tex sets.tex
	# pdflatex --file-line-error -interaction=nonstopmode sets.tex
	../../latexrun sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sets.pdf
	open sets.pdf

o:
	make open

terms: sets.tex
	bbk_sheet -mode ts sets.tex

spell:
	aspell -c sets.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f sets.*
	make reset
	bbk_sheet -mode mk > Makefile