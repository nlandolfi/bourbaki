# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subsequences.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

subsequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subsequences.tex

subsequences.pdf: ../../*.tex ../../trademark.pdf *.tex subsequences.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subsequences.tex
	../../latexrun subsequences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subsequences.pdf
	open subsequences.pdf

o:
	make open

terms: subsequences.tex
	bbk_sheet -mode ts subsequences.tex

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
	rm -f subsequences.*
	make reset
	bbk_sheet -mode mk > Makefile