# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make pair_unions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

pair_unions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > pair_unions.tex

pair_unions.pdf: ../../*.tex ../../trademark.pdf *.tex pair_unions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode pair_unions.tex
	../../latexrun pair_unions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: pair_unions.pdf
	open pair_unions.pdf

o:
	make open

terms: pair_unions.tex
	bbk_sheet -mode ts pair_unions.tex

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
	rm -f pair_unions.*
	make reset
	bbk_sheet -mode mk > Makefile