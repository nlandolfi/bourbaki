# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make hyperplanes.pdf

hyperplanes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > hyperplanes.tex

hyperplanes.pdf: ../../*.tex ../../trademark.pdf *.tex hyperplanes.tex
	# pdflatex --file-line-error -interaction=nonstopmode hyperplanes.tex
	../../latexrun hyperplanes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: hyperplanes.pdf
	open hyperplanes.pdf

o:
	make open

terms: hyperplanes.tex
	bbk_sheet -mode ts hyperplanes.tex

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
	rm -f hyperplanes.*
	make reset
	bbk_sheet -mode mk > Makefile