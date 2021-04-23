# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make total_variation_complete.pdf
	make graph.pdf

total_variation_complete.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > total_variation_complete.tex

total_variation_complete.pdf: ../../*.tex ../../trademark.pdf *.tex total_variation_complete.tex
	# pdflatex --file-line-error -interaction=nonstopmode total_variation_complete.tex
	../../latexrun total_variation_complete.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: total_variation_complete.pdf
	open total_variation_complete.pdf

o:
	make open

terms: total_variation_complete.tex
	bbk_sheet -mode ts total_variation_complete.tex

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
	rm -f total_variation_complete.*
	make reset
	bbk_sheet -mode mk > Makefile