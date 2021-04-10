# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make dice.pdf
	make graph.pdf

dice.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dice.tex

dice.pdf: ../../*.tex ../../trademark.pdf *.tex dice.tex
	# pdflatex --file-line-error -interaction=nonstopmode dice.tex
	../../latexrun dice.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dice.pdf
	open dice.pdf

o:
	make open

terms: dice.tex
	bbk_sheet -mode ts dice.tex

spell:
	aspell -c dice.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f dice.*
	make reset
	bbk_sheet -mode mk > Makefile