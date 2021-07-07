# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make entropy.pdf

entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > entropy.tex

entropy.pdf: ../../*.tex ../../trademark.pdf *.tex entropy.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode entropy.tex
	../../latexrun entropy.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: entropy.pdf
	open entropy.pdf

o:
	make open

terms: entropy.tex
	bbk_sheet -mode ts entropy.tex

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
	rm -f entropy.*
	make reset
	bbk_sheet -mode mk > Makefile