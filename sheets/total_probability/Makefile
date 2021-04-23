# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make total_probability.pdf
	make graph.pdf

total_probability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > total_probability.tex

total_probability.pdf: ../../*.tex ../../trademark.pdf *.tex total_probability.tex
	# pdflatex --file-line-error -interaction=nonstopmode total_probability.tex
	../../latexrun total_probability.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: total_probability.pdf
	open total_probability.pdf

o:
	make open

terms: total_probability.tex
	bbk_sheet -mode ts total_probability.tex

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
	rm -f total_probability.*
	make reset
	bbk_sheet -mode mk > Makefile