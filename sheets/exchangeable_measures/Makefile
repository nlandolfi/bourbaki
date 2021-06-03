# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make exchangeable_measures.pdf

exchangeable_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > exchangeable_measures.tex

exchangeable_measures.pdf: ../../*.tex ../../trademark.pdf *.tex exchangeable_measures.tex
	# pdflatex --file-line-error -interaction=nonstopmode exchangeable_measures.tex
	../../latexrun exchangeable_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: exchangeable_measures.pdf
	open exchangeable_measures.pdf

o:
	make open

terms: exchangeable_measures.tex
	bbk_sheet -mode ts exchangeable_measures.tex

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
	rm -f exchangeable_measures.*
	make reset
	bbk_sheet -mode mk > Makefile