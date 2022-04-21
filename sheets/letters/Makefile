# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make letters.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

letters.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > letters.tex

letters.pdf: ../../*.tex ../../trademark.pdf *.tex letters.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode letters.tex
	../../latexrun letters.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: letters.pdf
	open letters.pdf

o:
	make open

terms: letters.tex
	bbk_sheet -mode ts letters.tex

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
	rm -f letters.*
	make reset
	bbk_sheet -mode mk > Makefile