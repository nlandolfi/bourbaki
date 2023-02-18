# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make halfspaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

halfspaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > halfspaces.tex

halfspaces.pdf: ../../*.tex ../../trademark.pdf *.tex halfspaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode halfspaces.tex
	../../latexrun halfspaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: halfspaces.pdf
	open halfspaces.pdf

o:
	make open

terms: halfspaces.tex
	bbk_sheet -mode ts halfspaces.tex

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
	rm -f halfspaces.*
	make reset
	bbk_sheet -mode mk > Makefile