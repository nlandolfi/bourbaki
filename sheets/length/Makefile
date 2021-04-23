# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make length.pdf
	make graph.pdf

length.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > length.tex

length.pdf: ../../*.tex ../../trademark.pdf *.tex length.tex
	# pdflatex --file-line-error -interaction=nonstopmode length.tex
	../../latexrun length.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: length.pdf
	open length.pdf

o:
	make open

terms: length.tex
	bbk_sheet -mode ts length.tex

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
	rm -f length.*
	make reset
	bbk_sheet -mode mk > Makefile