# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make powers_and_unions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

powers_and_unions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > powers_and_unions.tex

powers_and_unions.pdf: ../../*.tex ../../trademark.pdf *.tex powers_and_unions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode powers_and_unions.tex
	../../latexrun powers_and_unions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: powers_and_unions.pdf
	open powers_and_unions.pdf

o:
	make open

terms: powers_and_unions.tex
	bbk_sheet -mode ts powers_and_unions.tex

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
	rm -f powers_and_unions.*
	make reset
	bbk_sheet -mode mk > Makefile