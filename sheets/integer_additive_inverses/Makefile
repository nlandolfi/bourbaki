# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_additive_inverses.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

integer_additive_inverses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integer_additive_inverses.tex

integer_additive_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex integer_additive_inverses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_additive_inverses.tex
	../../latexrun integer_additive_inverses.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_additive_inverses.pdf
	open integer_additive_inverses.pdf

o:
	make open

terms: integer_additive_inverses.tex
	bbk_sheet -mode ts integer_additive_inverses.tex

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
	rm -f integer_additive_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile