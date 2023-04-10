# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rational_multiplicative_inverses.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

rational_multiplicative_inverses.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > rational_multiplicative_inverses.tex

rational_multiplicative_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex rational_multiplicative_inverses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rational_multiplicative_inverses.tex
	../../latexrun rational_multiplicative_inverses.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_multiplicative_inverses.pdf
	open rational_multiplicative_inverses.pdf

o:
	make open

terms: rational_multiplicative_inverses.tex
	bbk_sheet -mode ts rational_multiplicative_inverses.tex

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
	rm -f rational_multiplicative_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile