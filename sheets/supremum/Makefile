# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make supremum.pdf

supremum.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > supremum.tex

supremum.pdf: ../../*.tex ../../trademark.pdf *.tex supremum.tex
	# pdflatex --file-line-error -interaction=nonstopmode supremum.tex
	../../latexrun supremum.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: supremum.pdf
	open supremum.pdf

o:
	make open

terms: supremum.tex
	bbk_sheet -mode ts supremum.tex

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
	rm -f supremum.*
	make reset
	bbk_sheet -mode mk > Makefile