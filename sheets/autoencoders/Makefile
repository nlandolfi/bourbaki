# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make autoencoders.pdf

autoencoders.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > autoencoders.tex

autoencoders.pdf: ../../*.tex ../../trademark.pdf *.tex autoencoders.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode autoencoders.tex
	../../latexrun autoencoders.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: autoencoders.pdf
	open autoencoders.pdf

o:
	make open

terms: autoencoders.tex
	bbk_sheet -mode ts autoencoders.tex

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
	rm -f autoencoders.*
	make reset
	bbk_sheet -mode mk > Makefile