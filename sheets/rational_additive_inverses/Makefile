# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rational_additive_inverses.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

rational_additive_inverses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rational_additive_inverses.tex

rational_additive_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex rational_additive_inverses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rational_additive_inverses.tex
	../../latexrun rational_additive_inverses.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_additive_inverses.pdf
	open rational_additive_inverses.pdf

o:
	make open

terms: rational_additive_inverses.tex
	bbk_sheet -mode ts rational_additive_inverses.tex

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
	rm -f rational_additive_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile