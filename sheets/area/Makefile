# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make area.pdf
	make graph.pdf

area.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > area.tex

area.pdf: ../../*.tex ../../trademark.pdf *.tex area.tex
	# pdflatex --file-line-error -interaction=nonstopmode area.tex
	../../latexrun area.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: area.pdf
	open area.pdf

o:
	make open

terms: area.tex
	bbk_sheet -mode ts area.tex

spell:
	aspell -c area.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f area.*
	make reset
	bbk_sheet -mode mk > Makefile