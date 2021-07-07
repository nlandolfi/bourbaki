# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make groups.pdf

groups.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > groups.tex

groups.pdf: ../../*.tex ../../trademark.pdf *.tex groups.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode groups.tex
	../../latexrun groups.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: groups.pdf
	open groups.pdf

o:
	make open

terms: groups.tex
	bbk_sheet -mode ts groups.tex

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
	rm -f groups.*
	make reset
	bbk_sheet -mode mk > Makefile